# Redis

## GO Redis

- [Ping-Pong](10_go-redis/10_ping-pong.go)
```shell
go run 10_ping-pong.go
```

- [String](10_go-redis/21_string.go)
```shell
go run 21_string.go
```

- [List](10_go-redis/23_list.go)
```shell
go run 23_list.go
```

## ristretto

### 简介

ristretto 是 dgraph 团队开源的一款高性能内存缓存库，旨在解决高并发场景下的缓存性能和吞吐瓶颈。dgraph 专攻的方向是高性能图数据库，ristretto 就是其图数据库和 KV 数据库产品的核心依赖。

与 golang 社区常见的其他单进程内存缓存类库（groupcache，bigcache，fastcache 等）相比，ristretto 在缓存命中率和读写吞吐率上的综合表现更优。ristretto 主要有以下优点：

- 高命中率 - 特殊设计的录入/驱逐政策
  - 驱逐（SampledLFU）：与精确 LRU 相当，但在搜索和数据跟踪上有更好的性能
  - 录入（TinyLFU）：以极小的内存开销获取额外的性能提升
- 高吞吐率
- 权重感知的驱逐策略 - 价值权重大的条目可以驱逐多个价值权重小的条目
  - 依托权重可以扩展出缓存最大内存占用、缓存最多条目数等场景
- 完全并发支持
- 性能指标 - 吞吐量、命中率及其他统计数据的性能指标
- 用户友好的 API 设计
- 支持指定缓存失效时间

ristretto 在 v0.1.0(2021-06-03) 版本发布时已正式标注为生产可用！

### Config

当创建 Ristretto 实例时，`Config`结构被传递到`NewCache`中。

#### NumCounters`int64`

NumCounters 是为允许和逐出而保留的 4 位访问计数器的数目。我们已经看到了很好的性能设置为 10 倍的项目数，您希望保留在缓存中的项目满。

例如，如果希望每个项目的成本为1，而MaxCost为100，则将NumCounters设置为1000。或者，如果使用可变成本值，但预期缓存在满时可容纳大约10000个项目，请将NumCounters设置为100000。重要的是完整缓存中唯一项的数量，而不一定是MaxCost值。

#### MaxCost `int64`

MaxCost是如何作出驱逐决定的。例如，如果MaxCost为100，并且成本为1的新项目将总缓存成本增加到101，则将逐出1个项目。MaxCost也可以用来表示以字节为单位的最大大小。例如，如果MaxCost为1000000（1MB），并且缓存中已满1000个1KB项，则新项（已接受）将导致5个1KB项被逐出。MaxCost可以是任何值，只要它与调用Set时使用cost值的方式相匹配。

#### BufferItems `int64`

BufferItems是获取缓冲区的大小。我们找到的最好值是64。如果出于某种原因，您看到Get性能随着大量争用而降低（您不应该），请尝试以64为增量增加此值。这是一个fine-tuning机制，你可能不必碰这个。

#### Metrics `bool`

当想要real-time记录各种统计数据时，度量是正确的。这是一个配置标志的原因是有10%的吞吐量性能开销。

一个受害者`func(hashes [2]uint64, value interface{}, cost int64)`

每次驱逐都要叫一个维克特。

- KeyToHash`func(key interface{}) [2]uint64`：KeyToHash是用于每个密钥的哈希算法。如果该值为零，则根据底层接口类型，Ristretto有多种默认值。注意，如果您想要128位散列，您应该使用完整的`[2]uint64`，否则只需在`0`位置填充`uint64`，它的行为将类似于任何64位散列。
- 成本`func(value interface{}) int64`：Cost是一个可选函数，您可以传递给Config，以便在运行时评估项目成本，并且仅适用于未删除的Set调用（如果计算项目成本特别昂贵，并且您不希望在将要删除的项目上浪费时间，则此函数非常有用）。向Ristretto发出信号，表示您希望使用此成本函数：

1. 将Cost字段设置为non-nil函数。
2. 当调用Set以获取新项或项更新时，请使用0的`cost`。

## Ref

1.[Golang之redis中间件框架](https://blog.csdn.net/QianLiStudent/article/details/103990921)
