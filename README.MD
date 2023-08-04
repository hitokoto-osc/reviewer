# 审核员服务

一言的审核员服务组件，目前主要重构后实现基于 ThinkJS 版本的审核员服务。未来将会跟进实现 v2 版本的审核员服务。

## 开发环境

* Node.js v20.x - 用于自动生成 Git Hooks
* GO v1.20.x - 主要开发环境

核心组件 以及 外部数据源：
* GoFrame
* Redis
* MySQL
* RabbitMQ

## 初始化
```shell
make init-env # 将会自动下载 Go 依赖包，以及初始化 Git Hooks
```

## 编译

```shell
make build
```

## 测试

```shell
make test
```
