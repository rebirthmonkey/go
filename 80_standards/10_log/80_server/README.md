# apiserver 日志输出

在 Gin 框架中，最标准的日志打印方式是通过 log.L(c).Info()  ，其中 c 里面包含了 requestid user 这些信息，传到 log 包中可以解析出来，作为日志的公共字段打印出来，从而丰富日志的信息。