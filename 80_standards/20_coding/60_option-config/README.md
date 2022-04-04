# Options & Config

## 简介

### Options vs. Config
- Option：外部命令行传过来的参数
- Config：程序内部使用的配置信息


一般来说，命令行传过来的 options 需要经过一些处理，才能供程序直接使用。比如：初始化数据库需要用的 host (ip:port)，可以通过 options 的 bind-address 和 bind-port 组合而来。

### 步骤
- 先读取外部 Options
- 然后转换为内部 Config

### Lab
- [Options and Config](example.go)


```shell
go run example.go config.go options.go
```