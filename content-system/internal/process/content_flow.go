package process

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/zerokkcoder/content-system/internal/dao"
// 	"gorm.io/gorm"

// 	flow "github.com/s8sg/goflow/flow/v1"
// 	goflow "github.com/s8sg/goflow/v1"
// )

// func ExecContentFlow(db *gorm.DB) {
// 	contentFlow := &ContentFlow{
// 		contentDao: dao.NewContentDao(db),
// 	}
// 	fs := goflow.FlowService{
// 		Port:              8081,
// 		RedisURL:          "localhost:6379",
// 		WorkerConcurrency: 5,
// 	}
// 	_ = fs.Register("content-flow", contentFlow.flowHandle)
// 	_ = fs.Start()
// }

// type ContentFlow struct {
// 	contentDao *dao.ContentDao
// }

// func (c *ContentFlow) flowHandle(workflow *flow.Workflow, context *flow.Context) error {
// 	// 创建节点
// 	dag := workflow.Dag()
// 	dag.Node("input", c.input)
// 	dag.Node("verify", c.verify)
// 	dag.Node("finish", c.finish)
// 	// 创建分支
// 	branches := dag.ConditionalBranch("branches",
// 		[]string{"category", "thumbnail", "format", "pass", "fail"}, func(b []byte) []string {
// 			var data map[string]interface{}
// 			if err := json.Unmarshal(b, &data); err != nil {
// 				return nil
// 			}
// 			if data["approval_status"].(float64) == 2 {
// 				return []string{"category", "thumbnail", "format", "pass"}
// 			}
// 			return []string{"fail"}
// 		}, flow.Aggregator(func(m map[string][]byte) ([]byte, error) {
// 			return []byte("ok"), nil
// 		}))
// 	branches["category"].Node("category", c.category)
// 	branches["thumbnail"].Node("thumbnail", c.thumbnail)
// 	branches["format"].Node("format", c.format)
// 	branches["pass"].Node("pass", c.pass)
// 	branches["fail"].Node("fail", c.fail)

// 	// 构建依赖关系
// 	dag.Edge("input", "verify")
// 	dag.Edge("verify", "branches")
// 	dag.Edge("branches", "finish")

// 	return nil
// }

// // input 输入节点
// func (c *ContentFlow) input(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec input")
// 	var input map[string]int64
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	id := input["content_id"]
// 	detail, err := c.contentDao.First(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result, err := json.Marshal(map[string]interface{}{
// 		"title":      detail.Title,
// 		"video_url":  detail.VideoURL,
// 		"content_id": detail.ID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// // verify 验证节点
// func (c *ContentFlow) verify(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec verify")
// 	var detail map[string]interface{}
// 	if err := json.Unmarshal(data, &detail); err != nil {
// 		return nil, err
// 	}
// 	var (
// 		title    = detail["title"]
// 		videoURL = detail["video_url"]
// 		id       = detail["content_id"]
// 	)
// 	// 机器审核/人工审核
// 	if int(id.(float64))%2 == 0 {
// 		detail["approval_status"] = 3
// 	} else {
// 		detail["approval_status"] = 2
// 	}
// 	fmt.Println(id, title, videoURL)
// 	return json.Marshal(detail)
// }

// // category 分类节点
// func (c *ContentFlow) category(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec category")
// 	var input map[string]interface{}
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	contentID := int64(input["content_id"].(float64))
// 	err := c.contentDao.UpdateByID(contentID, "category", "category")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte("category"), nil
// }

// // thumbnail 缩略图节点
// func (c *ContentFlow) thumbnail(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec thumbnail")
// 	var input map[string]interface{}
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	contentID := int64(input["content_id"].(float64))
// 	err := c.contentDao.UpdateByID(contentID, "thumbnail", "thumbnail")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte("thumbnail"), nil
// }

// // format 缩略图节点
// func (c *ContentFlow) format(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec format")
// 	var input map[string]interface{}
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	contentID := int64(input["content_id"].(float64))
// 	err := c.contentDao.UpdateByID(contentID, "format", "format")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []byte("format"), nil
// }

// // pass 通过节点
// func (c *ContentFlow) pass(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec pass")
// 	var input map[string]interface{}
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	contentID := int64(input["content_id"].(float64))
// 	err := c.contentDao.UpdateByID(contentID, "approval_status", 2)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// // fail 不通过节点
// func (c *ContentFlow) fail(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec fail")
// 	var input map[string]interface{}
// 	if err := json.Unmarshal(data, &input); err != nil {
// 		return nil, err
// 	}
// 	contentID := int64(input["content_id"].(float64))
// 	err := c.contentDao.UpdateByID(contentID, "approval_status", 3)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// func (c *ContentFlow) finish(data []byte, option map[string][]string) ([]byte, error) {
// 	fmt.Println("exec finish")
// 	fmt.Println(string(data))
// 	return data, nil
// }
