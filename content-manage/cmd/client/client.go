package main

import (
	"content_manage/api/operate"
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewAppClient(conn)
	// 创建内容
	// reply, err := client.CreateContent(context.Background(), &operate.CreateContentReq{
	// 	Content: &operate.Content{
	// 		Title:       "test content manage create",
	// 		VideoUrl:    "https://example.com/video.mp4",
	// 		Author:      "zerokk",
	// 		Description: "test",
	// 	},
	// })

	//更新内容
	// reply, err := client.UpdateContent(context.Background(), &operate.UpdateContentReq{
	// 	Content: &operate.Content{
	// 		Id:          4,
	// 		Title:       "test content manage create1",
	// 		VideoUrl:    "https://example.com/video1.mp4",
	// 		Author:      "zerokk1",
	// 		Description: "test1",
	// 	},
	// })

	// 删除内容
	// reply, err := client.DeleteContent(context.Background(), &operate.DeleteContentReq{
	// 	Id: 4,
	// })

	// 查找内容
	reply, err := client.FindContent(context.Background(), &operate.FindContentReq{
		// Title:    "test content manage create1",
		// Author:   "zerokk1",
		Id:       3,
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] DeleteContent reply: %v", reply)
}
