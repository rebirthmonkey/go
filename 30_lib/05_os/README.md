# OS

## Args

提取命令行输入的参数

- [命令行参数解析](02_os-args.go)

```shell
go run 02_os-args.go a b c d
```


## Flag

提取命令行输入的flag：cmd -a xxx

- [命令行flag解析](06_flag-parse.go)

```shell
go run 06_flag-parse.go -p 888
go run 06_flag-parse.go # 未输入参数，则采用默认
```

多flag

- [命令行多flag解析](07_flag-parse-args.go)

```shell
go run 07_flag-parse-args.go -d dd -l ll -w true
go run 07_flag-parse-args.go x y z # 未输入flag，则采用默认
```

## Config

从配置文件中获取参数

- [配置文件解析](11_goconfig-ini.go)

```shell
go run 11_goconfig-ini.go
```

## Ctl

为程序创建 CLI

- [CTL-Cobra框架](20_ctl-cobra/README.md)

 



