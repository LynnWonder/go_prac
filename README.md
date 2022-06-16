# gin_practices
> 一个简单的增删改查 gin 项目
## 项目目录结构

> 在 golang 中，每一个目录即是一个包，每个包都是一等公民，因此不会出现像其他语言一样的 src 目录。

```plain
.
├── README.md
├── biz
│   ├── common
│   │   └── global.go
│   ├── config
│   │   ├── config.go
│   │   └── logger.go
│   ├── dal
│   │   ├── gorm.go
│   │   └── person
│   │       └── person.go
│   ├── handler
│   │   ├── error.go
│   │   ├── handler.go
│   │   └── person
│   │       └── person.go
│   ├── model
│   │   └── person.go
│   └── service
│       └── person
│           ├── errors.go
│           └── person.go
├── conf
│   └── default.yaml
├── go.mod   标记每一个依赖包的版本，整理时 go mod tidy 
├── go.sum   记录每个依赖包的哈希值，这里不需要太过关注
├── img.png
├── logs
│   └── gin-prac.log
├── main.go
├── router   初次练习时用的路由，无需注意
│   ├── api
│   │   ├── index.go
│   │   └── userRouter
│   │       ├── auth.go
│   │       └── index.go
│   └── router.go
└── router.go    
    
```
1. conf 文件夹存基础配置文件
2. biz 文件夹，业务逻辑的组装层，类似于 DDD 的 domain 层
   - biz/common 存放常量和全局变量
   - biz/config 解析 yaml 文件、初始化日志收集等
3. biz/dal 文件夹，编写数据访问层
4. 注册路由 router.go
5. biz/handler handler 函数
6. biz/service 文件夹，类似于 DDD 的 application 层，实现从 handler 向 dal 的过渡

## 项目启动
配置 goland 启动

![img.png](img.png)

或
```golang
go run main.go
// 最简单的一个路由
// 127.0.0.1:8080/api/user-auth/login
```

## 一些问题

- 为什么还要使用 github.com/natefinch/lumberjack 写滚动日志
- gin 是在 net/http 基础上进行了一些封装？'

- handler chain 的概念

## TODO 
梳理上述的问题，添加中间件
