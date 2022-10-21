# API

## 简介

用来衡量 API 性能的指标主要有 3  个：

- 并发数（Concurrent）：并发数是指某个时间范围内，同时在使用系统的用户个数。广义上的并发数是指同时使用系统的用户个数，这些用户可能调用不同的 API；严格意义上的并发数是指同时请求同一个 API 的用户个数。
- 请求响应时间（TTLB）：请求响应时间指的是从客户端发出请求到得到响应的整个时间。这个过程从客户端发起的一个请求开始，到客户端收到服务器端的响应结束。在一些工具中，请求响应时间通常会被称为 TTLB（Time to last  byte，意思是从发送一个请求开始，到客户端收到最后一个字节的响应为止所消费的时间）。请求响应时间的单位一般为“秒”或“毫秒”。
- 每秒查询数（QPS）：每秒查询数 QPS 是对一个特定的查询服务器在规定时间内所处理流量多少的衡量标准。QPS = 并发数 /  平均请求响应时间。

这三个指标中，衡量 API 性能的最主要指标是 QPS，但是在说明 QPS 时，需要指明是多少并发数下的 QPS，否则毫无意义，因为不同并发数下的 QPS 是不同的。举个例子，单用户 100  QPS 和 100 用户 100 QPS 是两个不同的概念。前者说明 API 可以在一秒内串行执行 100 个请求，而后者说明在并发数为 100 的情况下，API 可以在一秒内处理 100 个请求。当 QPS 相同时，并发数越大，说明 API 性能越好，并发处理能力越强。在并发数设置过大时，API 同时要处理很多请求，会频繁切换上下文，而真正用于处理请求的时间变少，请求响应时间会变长，反而使得 QPS 会降低。API 会有一个合适的并发数，在该并发数下，API 的 QPS 可以达到最大，但该并发数不一定是最佳并发数，还要参考该并发数下的平均请求响应时间。

此外，在有些 API 接口中，也会测试 API 接口的 TPS（Transactions Per  Second，每秒事务数）。一个事务是指客户端向服务器发送请求，然后服务器做出反应的过程。客户端在发送请求时开始计时，收到服务器响应后结束计时，以此来计算使用的时间和完成的事务个数。那么，TPS 和 QPS 有什么区别呢？如果是对一个查询接口（单场景）压测，且这个接口内部不会再去请求其他接口，那么  TPS=QPS，否则，TPS≠QPS。如果是对多个接口（混合场景）压测，假设 N 个接口都是查询接口，且这个接口内部不会再去请求其他接口，QPS=N*TPS。

在项目上线前，我们需要对 API 接口进行性能测试。通常 API 接口的性能延时要小于 500ms ，如果大于这个值，需要考虑优化性能。

## wrk

### Lab

#### Install

```shell
git clone https://github.com/wg/wrk
cd wrk
make -j4
sudo install ./wrk /usr/bin
```

> 编译需要unzip

#### 参数

```bash
wrh --help
```

- -t：线程数（线程数不要太多，是核数的 2 到 4 倍就行，多了反而会因为线程切换过多造成效率降低）。
- -c：并发数。
- -d：测试的持续时间，默认为 10s。
- -T：请求超时时间。
- -H：指定请求的  HTTP Header，有些 API 需要传入一些 Header，可通过 wrk 的 -H  参数来传入。
- -latency：打印响应时间分布。
- -s：指定 Lua 脚本，Lua 脚本可以实现更复杂的请求。

#### VM Test

```shell
wrk -t50 -c1000 -d5s -T5s --latency http://192.168.1.10:8080/healthz

wrk -t50 -c1000 -d5s -T5s --latency http://192.168.1.10:8080/v1/users
```

> 假设已经安装`wrk`到path。否则，使用相对路径`path/to/wrk`代替`wrk`

> `http://192.168.1.10:8080/v1/users` 需要被替换

可以用`50_web/20_gin/96_insecure/`中的apiserver搭建测试服务器

```bash
cd ../../../50_web/20_gin/96_insecure/
go run cmd/apiserver.go -c configs/config.yaml &
sleep 10
wrk -t50 -c1000 -d5s -T5s --latency http://127.0.0.1:8080/healthz
wrk -t50 -c1000 -d5s -T5s --latency http://127.0.0.1:8080/healthz
```

下面是对测试结果的解析。

- 50 threads  and 1000 connections：用 50 个线程模拟 1000 个连接，分别对应 -t 和 -c 参数。

Thread Stats 是线程统计，包括 Latency 和  Req/Sec。

- Latency：响应时间，有平均值、标准偏差、最大值、正负一个标准差占比。
- Req/Sec：每个线程每秒完成的请求数,  同样有平均值、标准偏差、最大值、正负一个标准差占比。
- Latency Distribution 是响应时间分布。
- 2276265 requests in 30.10s, 412.45MB read：30.10s  完成的总请求数（2276265）和数据读取量（412.45MB）
- Socket errors: connect 1754, read 40,  write 0, timeout 0：错误统计，会统计 connect  连接失败请求个数（1754）、读失败请求个数、写失败请求个数、超时请求个数。
- Requests/sec：QPS。
- Transfer/sec：平均每秒读取 13.70MB 数据（吞吐量）。

#### K8s Test

```shell
wrk -t10 -c100 -d5s -T5s --latency http://129.204.97.254:8080/v1/users


wrk -t10 -c100 -d5s -T5s --latency http://192.168.34.15:8080/v1/users
```

> `http://129.204.97.254:8080/v1/users` 需要被替换



