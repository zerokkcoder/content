package services

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/redis/go-redis/v9"
	"github.com/zerokkcoder/content-system/internal/api/operate"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
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
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)
	endpoint := "discovery:///content_system"
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("localhost:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		panic(err)
	}
	appClient := operate.NewAppClient(conn)
	app.operationAppClient = appClient
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
	// mysqlDB = mysqlDB.Debug()

	// 创建 Reporter
	report := reporter.NewReporter("http://localhost:9411/api/v2/spans")
	// 创建本地节点
	endpoint, err := zipkin.NewEndpoint("mysql", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// 创建Zipkin Tracer
	tracer, err := zipkin.NewTracer(report,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithTraceID128Bit(true))
	if err != nil {
		panic(err)
	}
	zipTracer := zipkinot.Wrap(tracer)
	opentracing.SetGlobalTracer(zipTracer)
	// 使用 zipkin 插件
	err = mysqlDB.Use(gormopentracing.New(gormopentracing.WithTracer(zipTracer)))
	if err != nil {
		panic(err)
	}
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
