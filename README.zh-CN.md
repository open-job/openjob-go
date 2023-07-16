# openjob-go

> [中文说明](README.zh-CN.md) | [English](README.md)

提供 Go 对接 openjob-server 的基础封装。

## HTTP API

### 对接 server API

- 注册客户端 API: /server/register
- 客户端下线 API: /server/unregister
- 上报心跳 API: /server/heartbeat
- 上报执行状态/结果
- 上报执行日志
- 拉取主题信息

### 提供 client API

提供 client API 让 openjob-server 调用

- 响应检查
- 接收任务
