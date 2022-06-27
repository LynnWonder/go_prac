**背景知识**：

项目使用了 kitex(RPC 框架)，其底层实现了 Serialization 序列化和 Transport 传输。
- kitex 框架和命令行工具默认支持 thrift 和 proto3 两种 IDL，对应的 kitex 支持 thrift 和 protobuf 两种序列化协议。

- 在传输上 kitex 使用扩展的 thrift 作为底层传输协议（thrift 既是 IDL 格式也是序列化协议和传输协议）

注：
IDL 即 interface definition Language 为接口定义语言，是用来约定进行 rpc 的双方的语言。

本文件夹是编写 echo.thrift 之后使用 kitex cli 自动生成的一个 echo 服务：

```shell
# 生成 output 目录，里面含有编译产物
sh build.sh

# 运行
sh output/bootstrap.sh
```