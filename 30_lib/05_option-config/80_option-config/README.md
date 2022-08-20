# Option-Config

## 简介

### options

options 结构体一般用于存储启动时需要的信息，它的主要来源是 Arg、Flag 以及配置文件 Config-file（虽然被称为 config-file，但是确切的应该是选项文件 option-file）。Option 不会存储在 runtime 的数据结构，只是在启动时使用。其具体结构体包括：

#### Options

完整的 Options 结构体，包含 server、HTTPS、DB、Log 等的 option。在本例中位于 apiserver/options 包中，它用到的其他组件的 Options 结构体位于 pkg/options/ 包中。

- NewOptions()：创建 Options 结构体，因为创建时会自动填充默认值，所以不需要 Complete() 函数。
- Options.Validate()：给 Options 的值做校验。
- Options.ApplyTo(Config)：将 Options 的值转化、传递到 Config 结构体中。
- Options.Flags()/AddFlags()：为命令行添加针对 server、db、logs 等的 flag，并集成到 Options 中。

### config

针对某一个 app，在 app runtime 启动时，Options 就会被转换成 Config，用于在运行中给 app/server 使用。在本例中位于 apiserver/config 包中。

#### Config

Options 结构体面向配置文件，而 Config 结构体则更面向 app 运行时的结构体。

- NewConfig()：创建默认 config，如 `config := apiserver.NewConfig()`
- Options.ApplyTo(Config)：将 Options 转换、映射成 Config，如 `Options.ApplyTo(config)`。
  - Config.CreateConfigFromOptions(Options)：也可以用本函数合并以上 2 步，将 Options 直接转成新的 Config。
- Config.Complete()：为 Config 补充默认值
- Config.Validate()：验证 Config，从而将 Config 转换成 CompletedConfig 结构体

#### CompletedConfig

已经补充、验证完所有所需值的完整 Config

- CompletedConfig.New()：基于完整的 CompleteConfig 创建 runtime app，如`apiserver := completeConfig.New()`
  - 在本例中通过 createAPIServer(Config) 合并以上几步

### app

app 是针对 App 的结构体，它默认包含 Command、APIServer 两部分，位于 pkg/app 包中。

#### Command

cobra.Command 是 App 自带 Cobra 的 Command 结构体，用于处理 flag、option、config-file 等。

##### Option 模式

- WithOption()：将 Options 赋值给 App
- WithRunFunc()：将 回调函数赋值给 App 的启动流程
- WithDescription()：将 desc 赋值给 App 的 Description

##### buildCommand() 流程

- 创建 cmd
- 将 commands[] 添加到 cmd
- 设置 cmd.RunE 为 a.runCommand
- 为 cmd 的 FlagSet 添加 flags
- 将 cmd 赋值给 App.cmd
- App 通过 Run() 调用 App.cmd.Execute()

#### apiServer

##### apiServer

apiServer 结构体用于包含各种类型的 server，**是整个App核心的扩展处**，常包含 genericServer、grpcServer 等，本案例包含 genericServer（Gin），它位于 internal/apiserver/server.go。

- CompleteConfig.New()：把完整的 completeConfig 变成一个 runtime App（apiServer）
- server.PrepareRun()：对 apiserver 进行如 OpenAPI 以及其他 API 的安装等初始化操作，转换为 PreapreAPIServer，如 `preparedapiserver := apiserver.PrepareRun()`

##### preparedAPIServer

完成初始化后 apiServer，它位于 internal/apiserver/server.go。

- preparedapiserver.Run()：运行 preparedapiserver，如 `preparedapiserver.Run()`

## Lab

```shell
go run cmd/apiserver.go -c configs/config.yaml
```

#### Question

介绍 Cobra Command 加载的 flat、config-file 的逻辑及顺序？