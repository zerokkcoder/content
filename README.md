# Content

## 目录说明
- content-system: api 网关
- content-manage: 内容管理服务
- content-flow: 内容加工服务

项目中使用到 Prometheus，zipkin，Grafana，所以需要在系统中对应安装并启动。

## 运行
- 运行etcd(需安装etcd)
```
$ etcd
```
- 运行 content-manage

content-manage 目录下运行：
```
$ kratos run
```

- 运行 content-flow

content-flow 目录下运行
```
$ go run cmd/server/main.go
```
上面运行的 goflow 的 start 方法。也可以运行下面的命令，使用 goflow 的 startwork：
```
$ go run cmd/work/main.go
```

- 运行 content-system

content-system 目录下运行：
```
$ go run cmd/main.go
```



