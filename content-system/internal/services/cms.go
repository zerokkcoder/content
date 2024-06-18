package services

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/zerokkcoder/content-system/internal/api/operate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsApp struct {
	db  *gorm.DB
	rdb *redis.Client
	// flowService *goflow.FlowService
	operationAppClient operate.AppClient
}

func NewCmsApp() *CmsApp {
	app := &CmsApp{}
	// 连接数据库
	app.connDB()
	// 连接redis
	app.connRdb()
	// 连接内容操作服务
	app.connOperateAppClient()
	// app.flowService = flowService()
	// go func() {
	// 	process.ExecContentFlow(app.db)
	// }()
	return app
}

func (app *CmsApp) connOperateAppClient() {
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
	client := operate.NewAppClient(conn)
	app.operationAppClient = client
}

func (app *CmsApp) connDB() {
	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)

	mysqlDB = mysqlDB.Debug()

	app.db = mysqlDB
}

// func flowService() *goflow.FlowService {
// 	fs := &goflow.FlowService{
// 		RedisURL: "localhost:6379",
// 	}
// 	return fs
// }

func (app *CmsApp) connRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	app.rdb = rdb
}
