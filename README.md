# gin project

## 项目目录结构

> 在 golang 中，每一个目录即是一个包，每个包都是一等公民，因此不会出现像其他语言一样的 src 目录。

```plain
.
├── README.md
├── internal 不能被外部引用的内容
├── pkg 公用的 pkg
├── conf
│   └── default.yaml
├── go.mod   标记每一个依赖包的版本，整理时 go mod tidy 
├── go.sum   记录每个依赖包的哈希值，这里不需要太过关注
├── logs     日志文件
└── main.go 
    
```

## 项目启动
配置 goland 启动

项目使用了 [kitex](https://www.cloudwego.io/zh/docs/kitex/overview/) (RPC 框架)，其底层实现了 Serialization 序列化和 Transport 传输。
具体看 rpc 路由，需提前启动服务端应用，参考 rpc/README.md 
背景知识：
- kitex 框架和命令行工具默认支持 thrift 和 proto3 两种 IDL，对应的 kitex 支持 thrift 和 protobuf 两种序列化协议。 
  
- 在传输上 kitex 使用扩展的 thrift 作为底层传输协议（thrift 既是 IDL 格式也是序列化协议和传输协议）

启动方式如下：
```golang
go run main.go server
```

## TODO

- 增加数据库操作

## reference
- gin 处理 http 请求和 tcp 连接的模块实际上是 golang 的 net/http 模块，gin 实则是丰富了 net/http 做路由匹配部分的能力以及其他一些能力。 
- golang 项目结构参考 [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md),
    这个结构虽然有些比较老了比如 vendor 文件夹，但它也给了我们一些提示。其他可参考的项目结构[Go工程目录](https://blog.csdn.net/hezhanran/article/details/122056826)
- IDL 即 interface definition Language 为接口定义语言，是用来约定进行 rpc 的双方的语言。
- vendor 文件夹还是早期的管理依赖的方式，目前默认使用 go module，如果遇到使用 vendor 文件夹的也不必太在意
[Gorm docs](https://gorm.io/zh_CN/docs)
[gin 源码阅读(1) - gin 与 net/http 的关系](https://www.cnblogs.com/457220157-FTD/p/15331188.html)

