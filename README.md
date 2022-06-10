# gin_practices

## 项目目录结构

> 在 golang 中，每一个目录即是一个包，每个包都是一等公民，因此不会出现像其他语言一样的 src 目录。

```plain
.
├── go.mod # 标记每一个依赖包的版本
├── go.sum # 记录每个依赖包的哈希值，这里不需要太过关注
├── main.go
└── router
    ├── api # 总的 api 路由 /api
    │   ├── index.go
    │   └── userRouter # 子路由 /user-auth
    │       ├── auth.go
    │       └── index.go
    └── router.go # 路由初始化
```

## 项目启动

```golang
go run main.go
```

## 一些问题

- gin 是在 net/http 基础上进行了一些封装？'

- handler chain 的概念
