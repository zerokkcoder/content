// 构建单一工作流
package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"math/rand"
// 	"strconv"

// 	flow "github.com/s8sg/goflow/flow/v1"
// 	goflow "github.com/s8sg/goflow/v1"
// )

// func Input(data []byte, option map[string][]string) ([]byte, error) {
// 	var input map[string]int
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	outputInt := input["input"]
// 	return []byte(strconv.Itoa(outputInt)), nil
// }

// func AddOne(data []byte, option map[string][]string) ([]byte, error) {
// 	num, _ := strconv.Atoi(string(data))
// 	outputInt := num + rand.Intn(10) + 1
// 	fmt.Println("AddOne = ", outputInt)
// 	return []byte(strconv.Itoa(outputInt)), nil
// }

// func AddTwo(data []byte, option map[string][]string) ([]byte, error) {
// 	num, _ := strconv.Atoi(string(data))
// 	outputInt := num + rand.Intn(101) + 100
// 	fmt.Println("AddTwo = ", outputInt)
// 	return []byte(strconv.Itoa(outputInt)), nil
// }

// func Output(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("data = ", string(data))
// 	return []byte("ok"), nil
// }

// func MyFlow(flow *flow.Workflow, context *flow.Context) error {
// 	// 创建节点
// 	dag := flow.Dag()
// 	dag.Node("input", Input)
// 	dag.Node("add-one", AddOne)
// 	dag.Node("add-two", AddTwo)
// 	dag.Node("output", Output)

// 	// 构建依赖关系
// 	dag.Edge("input", "add-one")
// 	dag.Edge("add-one", "add-two")
// 	dag.Edge("add-two", "output")

// 	return nil
// }

// func main() {
// 	// 创建工作流服务
// 	fs := goflow.FlowService{
// 		Port:              8081,
// 		RedisURL:          "localhost:6379",
// 		WorkerConcurrency: 5,
// 	}
// 	// 注册工作流
// 	fs.Register("add-flow", MyFlow)
// 	if err := fs.Start(); err != nil {
// 		panic(err)
// 	}
// }
