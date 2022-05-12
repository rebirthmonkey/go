# Go 并发

## Goroutine协程

Go 的高并发其实是由 goroutine 实现的，Goroutine 可以被认为是轻量的线程，创建一个 goroutine 只需在调用函数的前面加上 `go` 这个关键字就行。goroutine 就是一段代码，一个函数入口，以及在堆上为其分配的一个堆栈。所以它非常廉价，可以很轻松的创建上万个 goroutine。但它们并不是被 OS 所调度执行，而是通过系统的线程来多路派遣这些函数的执行，使得每个用`go`关键字执行的函数可以运行成为一个单位协程。当一个协程阻塞的时候，调度器就会自动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。而且调度的开销非常小，一颗 CPU 调度的规模不下于每秒百万次，这使得在程序中能够创建大量的 goroutine，实现高并发的同时，依旧能保持高性能。在单个程序中，所有 goroutines 都是共享相同的地址空间。

## Channel

一个 channel 相当于一个先进先出（FIFO）的队列。也就是说，通道中的各个元素值都是严格地按照发送的顺序排列的，先被发送通道的元素值一定会先被接收。元素值的发送和接收都需要用到操作符 <-，一个左尖括号紧接着一个减号形象地代表了元素值的传输方向。元素值从外界进入通道时会被复制。更具体地说，进入通道的并不是在接收操作符右边的那个元素值，而是它的副本。

channel 是一种 Go 语言结构，它通过传递特定元素类型的值来为两个 goroutines 提供同步执行和交流数据的机制 。可以将 `channel`视为管道， `goroutine`可以从管道发送和接收来自其他 `goroutine`的信息。 <-  标识符表示了 channel 的传输方向，接收或者发送。

### 状态

channel 存在 3 种状态：

- nil：未初始化的状态，只进行了声明，或者手动赋值为`nil`
- active：正常的 channel，可读或可写
- closed：已关闭，千万不要误认为关闭 channel 后，channel的值是nil

### 操作

channel 可进行 3 种操作：

1. 读
2. 写
3. 关闭：close() 的作用是保证不能再向 channel 中发送值。 channel 被关闭后，仍然是可以从中接收值的。

### 发送/接收

通过 channel 发送值，可使用 <- 作为二元运算符。通过 channel 接收值，可使用它作为一元运算符。

```go
ic <- 3       // 向channel中发送3
work := <-wc  // 从channel中接收指针到work
```

#### Range


伴有 range 分句的 for 语句会连续读取通过 channel 发送的值，直到 channel 被关闭。

### 类型

Channels 是一种被 make 分配的引用类型

```go
ic := make(chan int)        // 不带缓存的  int channel
wc := make(chan *Work, 10)  // 带缓冲工作的 channel
```

#### 无缓冲

如果 channel 是无缓冲的，发送者会一直阻塞直到有接收者从中接收值。也就是说，一个 goroutine 向 channel 发生完后会一直阻塞，直到 channel 中的消息被另一个 goroutine 消费完了之后才能接触阻塞。

**无缓冲通道**的特点是，发送的数据需要被读取后，发送才会完成（解除阻塞）。它阻塞场景：

- 读空 channel：通道中无数据，但执行读通道。
- 写了 channel 但无人读：通道中无数据，向通道写数据，但无协程读取。

#### 有缓冲

**有缓存通道**的特点是，有缓存时可以向通道中写入数据后直接返回，缓存中有数据时可以从通道中读到数据直接返回，这时有缓存通道是不会阻塞的。如果是带缓冲的，只有当值被拷贝到缓冲区且缓冲区已满时，发送者才会阻塞直到有接收者从中接收。接收者会一直阻塞直到 channel 中有值可被接收。它阻塞的场景是：

- 读空 channel：通道的缓存无数据，但执行读通道。
- 写满 channle：通道的缓存已经占满，向通道写数据，但无协程读。

### select

select 是执行选择操作的一个结构，它里面有一组 case 语句。它会执行其中无阻塞的那一个，如果都阻塞了，那就等待其中一个不阻塞，进而继续执行，它有一个 default 语句，该语句是永远不会阻塞的，可以借助它实现无阻塞的操作。



## 锁

当多个 goroutine 同时进行处理的时候，就会遇到同时抢占一个资源的情况，所以希望某个 goroutine 等待另一个 goroutine 处理完某一个步骤之后才能继续。sync 包就是为了让 goroutine 同步而出现的，当然还可以使用channel实现。

### 锁

锁有两种：互斥锁（mutex）和读写锁（RWMutex）

- 互斥锁：当数据被加锁了之后，除次外的其他协程不能对数据进行读操作和写操作。这当然能解决并发程序对资源的操作。但是，效率上是个问题，因为当加锁后，其他协程只有等到解锁后才能对数据进行读写操作。
- 读写锁：读数据的时候上读锁，写数据的时候上写锁。有写锁的时候，数据不可读不可写。有读锁的时候，数据可读，不可写。

### waitGroup

sync.WaitGroup 是 go 标准库的一部分，它等待一系列 goroutines 执行结束。

```go
var wg sync.WaitGroup
wg.Add(5)
for i := 0; i < 5; **i++** {
    go func() {
        **fmt.Print(i)** // 局部变量i被6个goroutine共享
        wg.Done()
    }()
}
wg.Wait() // 等待5个goroutine执行结束
fmt.Println()
```






## Lab
- [goroutine](10_goroutine.go)
```shell
go run 10_goroutine.go
```

- [goroutine](12_goroutine-anonym.go)
```shell
go run 12_goroutine-anonym.go 
```

- [channel](20_channel.go)
```shell
go run 20_channel.go 
```

- [确保 goroutine 完成](21_channnel.go)：main 函数中起了一个 goroutine，通过非缓冲队列的使用，能够保证在 goroutine 执行结束之前 main 函数不会提前退出

```shell
go run 21_channnel.go
```

- [channel range](22_channel-range.go)：通过 for range 处理 channel 中的所有消息
```shell
go run 22_channel-range.go
```

- [channel selelct](26_channel-select.go)：通过 select 自动选择先结束的 goroutine 来处理
```shell
go run 26_channel-select.go
```

- [多核并行](30_multi-process.go)：不展示任何结果，goroutine还没开始，main就会退出

```shell
go run 30_multi-process.go
```

- [多核并行](31_multi-process.go)：每次执行的结果都不相同，在多个routine还没执行完成的情况下，因为routine 9完成了，所以main就会退出

```shell
go run 31_multi-process.go
```

- [多核并行](32_multi-process-channel.go)：通过channel来确认10个routine都已完成

```shell
go run 32_multi-process-channel.go
```

- [多核并行](33_multi-process-wg.go)：通过 waitGroup 来确认 10 个 routine 都已完成

```shell
go run 33_multi-process-wg.go
```

- [Mutex锁](41_sync-mutex.go)
```shell
go run 41_sync-mutex.go 
```

- [Once只执行一次](43_sync-once.go)
```shell
go run 43_sync-once.go 
```



## Ref

1. [总结了才知道，原来channel有这么多用法！](https://segmentfault.com/a/1190000017958702)
2. [Golang之context](https://www.toutiao.com/article/7053778893844529695)
3. [golang中Context的使用场景](https://mp.weixin.qq.com/s/xbDFN-JhTIQ4xWanEC1Bxw)
4. 
