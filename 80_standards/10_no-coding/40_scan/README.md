# 代码扫描



### 格式化工具

[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) 是 Go 语言官方提供的工具，它能够为我们自动格式化 Go 语言代码并对所有引入的包进行管理，包括自动增删依赖的包引用、将依赖包按字母序排序并分类。很多人都用过 gofmt 来对代码进行格式化，goimports 等于gofmt加上依赖包管理（goimports=gofmt + import）。

很多 IDE 都支持在保存代码时执行 goimports，使用 goimports 可能会偶尔引入一些错误的包，但是它带来的好处也是明显的，这里建议大家每次写完代码后都 goimports 一下，如果不想用 goimports，那么至少要用 gofmt 格式化下代码。

### 静态代码检查工具

[golangci-lint ](https://github.com/golangci/golangci-lint) 是一个 linter 聚合器，它的速度很快，平均速度是 gometalinter 的 5 倍。它易于集成和使用，具有良好的输出并且具有最小数量的误报，支持 YAML 配置。目前是 Go 项目静态代码检查的最佳工具。

golangci-lint 虽然可以灵活的通过YAML文件配置检查和不检查哪些规则，但这里建议每个项目，起初都开启最严的检查配置，不要随意忽略某个检查项。

> golangci-lint官方文档，请参考：[golangci-lint官方文档](https://golangci-lint.run/)

