package main

import "sync"

// RequestVoteArgs 定义了 RequestVote RPC 请求的参数
type RequestVoteArgs struct {
	Term         int // 候选人的任期号
	CandidateID  int // 请求选票的候选人的 ID
	LastLogIndex int // 候选人最后日志条目的索引值
	LastLogTerm  int // 候选人最后日志条目的任期号
}

// RequestVoteReply 定义了 RequestVote RPC 响应的结构
type RequestVoteReply struct {
	Term        int  // 当前任期号，以便于候选人更新自己的任期号
	VoteGranted bool // 候选人赢得了此张选票时为真
}

// Raft 定义 Raft 服务器的结构
type Raft struct {
	mu          sync.Mutex
	currentTerm int
	votedFor    *int
	log         []LogEntry
	commitIndex int
	// 更多状态...
}

// LogEntry 定义了日志条目的结构
type LogEntry struct {
	Term    int
	Command interface{}
}

// RequestVote 处理 RequestVote RPC 请求
func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) error {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	// 假设 lastLogIndex 和 lastLogTerm 是这个服务器最后日志条目的所有和任期
	lastLogIndex := len(rf.log) - 1
	lastLogTerm := rf.log[lastLogIndex].Term

	// 回复的任期号始终是服务器当前的人气好
	reply.Term = rf.currentTerm

	// 检查候选人的任期是否至少和本服务器的任期一样新
	if args.Term < rf.currentTerm {
		reply.VoteGranted = false
		return nil
	}

	// 如果候选人的日志不如本服务器的日志新，则不投票
	if args.LastLogTerm < lastLogTerm || (args.LastLogTerm == lastLogTerm && args.LastLogIndex < lastLogIndex) {
		reply.VoteGranted = false
		return nil
	}

	// 检查是否已经给别的候选人投票
	if rf.votedFor == nil || *rf.votedFor == args.CandidateID {
		rf.votedFor = &args.CandidateID // 更新 votedFor
		reply.VoteGranted = true
	} else {
		reply.VoteGranted = false
	}

	// 如果收到的任期号比本身服务器的任期号更大，更新本服务器的任期号
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = nil
	}

	return nil
}

// AppendEntriesArgs 定义了 AppendEntries RPC 请求的参数
type AppendEntriesArgs struct {
	Term         int        // 领导者的任期号
	LeaderID     int        // 领导者的 ID，跟随者可以据此重定向请求
	PrevLogIndex int        // 上一次成功复制的最后一条日志条目所对应的index
	PrevLogTerm  int        // prevLogIndex 条目的任期号
	Entries      []LogEntry // 准备好的日志条目（心跳时为空；为了简化，假设足够一次发送完）
	LeaderCommit int        // 领导者的已知提交的最高的日志条目的索引值
}

// AppendEntriesReply 定义了 AppendEntries RPC 响应的结构
type AppendEntriesReply struct {
	Term    int  // 当前任期号，对于领导者而言，它会更新自己的任期号
	Success bool // 如果跟随者已经包含了匹配上 PrevLogIndex 以及 PrevLogTerm 的那条日志条目，则为真
}

// AppendEntries 处理 AppendEntries RPC 请求
func (rf *Raft) AppendEntries(args AppendEntriesArgs, reply *AppendEntriesReply) error {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply.Term = rf.currentTerm

	// 如果领导者的任期小于当前任期，拒绝
	if args.Term < rf.currentTerm {
		reply.Success = false
		return nil
	}

	// 如果日志在prevLogIndex日志的条目的任期号与prevLogTerm不匹配，拒绝
	if args.PrevLogIndex >= len(rf.log) || rf.log[args.PrevLogIndex].Term != args.PrevLogTerm {
		reply.Success = false
		return nil
	}

	// 如果存在冲突的条目（同一个索引但任期号不同），删除这一条目及其之后的所有的条目
	conflictIndex := -1
	for i := args.PrevLogIndex + 1; i < len(rf.log) && i-args.PrevLogIndex-1 < len(args.Entries); i++ {
		if rf.log[i].Term != args.Entries[i-args.PrevLogIndex-1].Term {
			conflictIndex = i
			break
		}
	}
	if conflictIndex != -1 {
		rf.log = rf.log[:conflictIndex]
	}

	// 追加任何在已有的日志中不存在的条目
	if conflictIndex != -1 || len(args.Entries) > len(rf.log)-args.PrevLogIndex-1 {
		rf.log = append(rf.log, args.Entries[len(args.Entries)-1:]...)
	}

	// 如果领导者的已知已提交的日志条目的索引值大于本地的，更新本地的 commitIndex
	if args.LeaderCommit > rf.commitIndex {
		lastNewIndex := args.PrevLogIndex + len(args.Entries)
		if args.LeaderCommit < lastNewIndex {
			rf.commitIndex = args.LeaderCommit
		} else {
			rf.commitIndex = lastNewIndex
		}
		// 应用到状态机（这一步通常时异步的，这里只是示例）
		// applyToStateMachine(rf.commitIndex)
	}

	reply.Success = true
	// 如果收到的任期号比本服务器的任期号更大，更新本服务器的任期号
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.votedFor = nil
	}

	return nil
}
