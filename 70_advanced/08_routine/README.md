# Goroutine协程

## 简介

Go 直接内置了对并发的支持， Go 里的并发指的是能让某个函数独立于其他函数运行的能力。Go 的并发其实是由协程 goroutine 实现，Goroutine 可以被认为是轻量的线程，创建一个 goroutine 只需在调用函数的前面加上 `go` 这个关键字就行。

### 并发/并行

并发(concurrency)与并行(parallelism)不同：

- 并行：是让不同的代码片段同时在不同的物理处理器上执行，并行的关键是同时做很多事情。
- 并发：是指同时管理很多事情，这些事情可能只做了一半就被暂停去做别的事情了。Go 采用并发通过切换多个线程达到减少物理处理器空闲等待的目的。

在很多情况下，并发的效果比并行好，因为 OS 和硬件的总资源一般很少，但能支持系统同时做很多事情。这种“使用较少的资源做更多的事情”的哲学，也是指导 Go 设计的哲学。 如果希望让 goroutine 并行，必须使用多于一个逻辑处理器。当有多个逻辑处理器时，调度器会将 goroutine  平等分配到每个逻辑处理器上。这会让 goroutine 在不同的线程上运行。不过要想真的实现并行的效果，用户需要让自己的程序运行在有多个物理处理器的机器上。否则，哪怕 Go 运行时使用多个线程，goroutine 依然会在同一个物理处理器上并发运行，达不到并行的效果。下图展示了在一个逻辑处理器上并发运行  goroutine 和在两个逻辑处理器上并行运行两个并发的 goroutine 之间的区别。

<img src="figures/image-20221017192212848.png" alt="image-20221017192212848" style="zoom:50%;" />

### 进程/线程/协程

- 进程 Process：当运行一个应用程序的时候，OS 会为这个应用程序启动一个进程。可以将这个进程看作一个包含了应用程序在运行中需要用到和维护的各种资源的容器。这些资源包括但不限于内存地址空间、文件和设备的句柄以及线程。
- 线程 Thread：一个线程是一个执行空间，这个空间会被 OS 调度来运行函数中所写的代码。每个进程至少包含一个线程，每个进程的初始线程被称作主线程。因为执行这个线程的空间是应用程序的本身的空间，所以当主线程终止时，应用程序也会终止。OS 将线程调度到某个处理器上运行，这个处理器并不一定是进程所在的处理器。

<img src="figures/image-20221017192346270.png" alt="image-20221017192346270" style="zoom:50%;" />

### goroutine 原理

goroutine 就是一段代码，一个函数入口，以及在堆上为其分配的一个堆栈。所以它非常廉价，可以很轻松的创建上万个 goroutine。但它们并不是被 OS 所调度执行，而是通过 Go 自己的调度器来多路派遣这些函数的执行，使得每个用`go`关键字执行的函数可以运行成为一个单位协程。

当一个函数创建为 goroutine 时，Go 会将其视为一个独立的工作单元，这个单元会被调度到可用的逻辑处理器上执行。Go 运行时的调度器是一个复杂的软件，能管理被创建的所有  goroutine 并为其分配执行时间。这个调度器在 OS 之上，将 OS 的线程与语言运行时的逻辑处理器绑定，并在逻辑处理器上运行  goroutine。调度器在任何给定的时间，都会全面控制哪个 goroutine 要在哪个逻辑处理器上运行。 

当一个协程阻塞的时候，Go 的调度器就会自动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。而且调度的开销非常小，一颗 CPU 调度的规模不下于每秒百万次，这使得在程序中能够创建大量的 goroutine，在实现高并发的同时依旧能保持高性能。在单个程序中，所有 goroutine 都是共享相同的地址空间。

调度器对可以创建的逻辑处理器的数量没有限制，但语言运行时默认限制每个程序最多创建 1W 个线程。这个限制值可以通过调用 runtime/debug 包的 SetMaxThreads 方法来更改。如果程序试图使用更多的线程，就会崩溃。

## 协程通讯

Go 的并发同步模型来自一个叫作 CSP（Communicating Sequential  Processes，CSP）的范型。CSP 是一种消息传递模型，通过在 goroutine  之间传递数据来传递消息，而不是对数据进行加锁来实现同步访问。

Go 中用于协程间通信和管理的有 channel 和 sync 包。比如 channel 可以通知协程做特定操作（退出、阻塞等），sync 可以加锁和同步。

### Channel

用于在 goroutine 之间同步和传递数据的关键数据类型叫作 channel。

#### 原理

`goroutine`可以从管道发送和接收来自其他 `goroutine`的信息，channel 通过传递特定元素类型的值来为两个 goroutines 提供同步执行和交流数据的机制 。

channel 相当于一个 FIFO 队列，通道中的各个元素值都是严格地按照发送的顺序排列，先被发送通道的元素值一定会先被接收。元素值的发送和接收都需要用到操作符 <-，一个左尖括号紧接着一个减号形象地代表了元素值的传输方向。元素值从外界进入通道时会被复制，也就是进入通道的并不是在接收操作符右边的那个元素值，而是它的副本。

一般情况下：

- 当一个 goroutine 向 channel 发送信息后，它会处于阻塞状态，直到 channel 中的信息被读取。
- 当一个 goroutine 向 channel 读取信息后，如果 channel 中没有信息，则它也会处于阻塞状态，直到有另一个 goroutine 向其中发送信息。

#### 类型

Channels 是一种被 make 分配的引用类型

```go
ic := make(chan int)        // 不带缓存的  int channel
wc := make(chan *Work, 10)  // 带缓冲工作的 channel
```

##### 无缓冲

如果 channel 是无缓冲的，发送者会一直阻塞直到有接收者从中接收值。也就是说，一个 goroutine 向 channel 发生完后会一直阻塞，直到 channel 中的消息被另一个 goroutine 消费完了之后才能接触阻塞。

**无缓冲通道**的特点是，发送的数据需要被读取后，发送才会完成（解除阻塞）。它阻塞场景：

- 写了 channel 但无人读：通道中无数据，向通道写数据，但无 goroutine 读取。
- 读空 channel：通道中无数据，但执行读通道。

##### 有缓冲

有缓存时，当向通道中写入数据后，如果通道未满，则可以直接返回。当向通道读取数据时，当缓存中有数据时，可以从通道中读到数据直接返回，这时有缓存通道是不会阻塞的。

它阻塞的场景是：

- 读空 channel：channel 的缓存无数据，但执行读通道，接收者会一直阻塞直到 channel 中有值可被接收。
- 写满 channel：channel 的缓存已经占满，但向通道写数据，发送者才会阻塞直到有接收者从中接收。

#### 状态

channel 存在 3 种状态：

- nil：未初始化的状态，只进行了声明，或手动赋值为`nil`
- active：正常的 channel，可读或可写
- closed：已关闭，千万不要误认为关闭 channel 后，channel的值是 nil

#### 操作

channel 可进行 3 种操作：

1. 读：`work := <-c`
2. 写：`ic <- 3`
3. 关闭：close() 的作用是保证不能再向 channel 中发送值。 channel 被关闭后，仍然是可以从中接收值的。

通过 channel 发送值，可使用 <- 作为二元运算符。通过 channel 接收值，可使用它作为一元运算符。

```go
ic <- 3       // 向channel中发送3
work := <- ic  // 从channel中接收指针到work
```

##### 循环读/ Range

伴有 range 分句的 for 语句会连续读取通过 channel 发送的值，直到 channel 被关闭。

##### select

select 是执行选择操作的一个结构，它里面有一组 case 语句。它会执行其中无阻塞的那一个，如果都阻塞了，那就等待其中一个不阻塞，进而继续执行。它有一个 default 语句，该语句是永远不会阻塞的，可以借助它实现无阻塞的操作。

### Sync

当多个 goroutine 同时进行处理的时候，就会遇到同时抢占一个资源的情况，所以希望某个 goroutine 等待另一个 goroutine 处理完某一个步骤之后才能继续。sync 就是为了让 goroutine 同步而出现的，是 channel 的一种代替方案。

#### 锁

锁有两种：互斥锁（mutex）和读写锁（RWMutex）

- 互斥锁：当数据被加锁了之后，除次外的其他协程不能对数据进行读操作和写操作。这当然能解决并发程序对资源的操作。但是，效率上是个问题，因为当加锁后，其他协程只有等到解锁后才能对数据进行读写操作。
- 读写锁：读数据的时候上读锁，写数据的时候上写锁。有写锁的时候，数据不可读不可写。有读锁的时候，数据可读，不可写。

#### waitGroup

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
wg.Wait() // 等待上述 5 个 goroutine 执行结束
fmt.Println()
```

### Context

使用 Channel 每次都要在协程内部增加对 channel 的判断，也要在外部设置关闭条件。Context 是协程的上下文，主要用于跟踪协程的状态，可以做一些简单的协程控制，也能记录一些协程信息。通过 Context 可以进一步简化控制代码，且更为友好的是，大多数 go 库，如 http、各种 db driver、grpc 等都内置了对 ctx.Done() 的判断，只需要将 ctx 传入即可。

```go
// 空的父context
pctx := context.TODO()

// 子context（携带有超时信息），cancel函数（可以主动触发取消）
ctx, cancel := context.WithTimeout(pctx, 5*time.Second)

for i := 0; i < 2; i++ {
   go func(i int) {
      // do something

  // 大部分工具库内置了对ctx的判断，下面的部分几乎可以省略
      select {
      case <-ctx.Done():
         fmt.Printf("%d Done\n", i)
      }
   }(i)
}

// 调用cancel会直接关闭ctx.Done()返回的管道，不用等到超时
//cancel()

time.Sleep(6 * time.Second)
```

#### 原理

在启动一个程序后，会创建一个 context 作为根 context，当需要创建更多的协程时，则基于这个根 context 来创建子 context，并且将根 context 的信息携带给子 context。同理子 context 也可以创建更多的自己的子 context 给后续创建的协程。当其中有一个父协程处理超时，可以通知自己下面的所有子协程，这样这些子协程也不用再继续自己的任务，避免浪费资源和性能。

因此， context 整体是一个树形结构，不同的 context 间可能是兄弟节点或父子节点关系。由于 Context 接口有多种不同的实现，所以树的节点可能也是多种不同的 context 实现。总的来说 Context 的特点是：

- 树形结构：每次调用 WithCancel、WithValue、WithTimeout、WithDeadline 实际是为当前节点在追加子节点。
- 继承性：某个节点被取消，其对应的子树也会全部被取消。
- 多样性：节点存在不同的实现，故每个节点会附带不同的功能。

如下图中，根 context 的第一个子 context 执行超时了，它后续的所有子 context 都能接收到 context 被取消的通知，进而取消自己的 context。

<img src="figures/image-20221017192407377.png" alt="image-20221017192407377" style="zoom:50%;" />

#### 分类

- context.Background：是 context 默认值，一般用在主函数（入口函数）或最初的根 context，其他所有的 context 上下文都是基于它创建出来。
- context.Todo：仅在不知道使用哪种 context 时使用。

但从其实现的源代码来看，Background 和 Todo 可以认为就是互为别名、基本没有差别，很多时候互用也没有任何关系。

#### 操作

- 新建 Context：返回一个空的 Context，这个 Context 一般用来做父 Context。

```go
ctx := context.TODO()
ctx := context.Background()
```

- WithCancel：会根据传入的 Context 生成一个子 Context 和 cancel() 取消函数。当父 Context 有相关取消操作，或直接调用 cancel() 函数，子 Context 就会被取消。

```go
// 一般操作比较耗时等，都会在输入参数里带上一个ctx
func Do(ctx context.Context, ...) {
 ctx, cancel := context.WithCancel(parentCtx)

 // 实现某些业务逻辑

 // 当遇到某种条件，如程序出错，就取消掉子Context，这样子Context绑定的协程也可以跟着退出
 if err != nil {
  cancel()
 }
}
```

- WithTimeout：给 context 附加一个超时控制。当超过指定的时长后，能被自动取消的 context。

```go
func demo(ctx context.Context) {
	ctx, cancel = context.WithTimeout(ctx, 30 * time.Second)
  defer cancel()
  
  go doSomething1(ctx)
  go doSomething2(ctx)
}
```

- WithDeadline：用于创建一个到达指定时间后能被自动取消的 context。WithTimeout 是指定时长，而 WithDeadline 是指定时间点。

```go
func demo(ctx context.Context) {
	// 创建一个1分钟后便会超时取消的context
  t := time.Now().Add(time.Minute)
	ctx, cancel = context.WithDeadline(ctx, t)
  defer cancel()
  
  go doSomething1(ctx)
  go doSomething2(ctx)
}
```

- WithValue：用来保存一些如链路追踪等信息，比如 API 服务里会有来保存一些来源 IP、请求参数等。这个方法比较常用了，如 grpc-go 里的 metadata 就使用这个方法将结构体存储在 ctx 里。在使用 Value() 查找 key 对应的值时，如果没找到，就会从父 context 中查找，直某个父 context 中返回 nil 或找到对应的值。

```go
//  传入父Context和(key, value)，相当于存一个kv
ctx := context.WithValue(parentCtx, "name", 123)
// 用法：将key对应的值取出
v := ctx.Value("name")
```

## Lab

- [goroutine](10_goroutine.go)：goroutine 持续循环，main() 等待外部输入然后中止整个进程

```shell
go run 10_goroutine.go
# 从键盘输入
```

或者使用cat创建输入

```bash
cat "some text 10_goroutine" | go run 10_goroutine.go
```

- [goroutine](12_goroutine-anonym.go)：同上，只是 goroutine 通过一个匿名函数来实现

```shell
go run 12_goroutine-anonym.go 
# 从键盘输入
```

或者使用cat创建输入

```bash
cat "some text 12_goroutine-anonym" | go run 12_goroutine-anonym.go 
```

- [channel](20_channel.go)：当 goroutine 写入 channel 后，因为没有读取操作，所以 goroutine 处于阻塞状态。直到 channel 被读取后，整个 goroutine 才开始循环。

```bash
go run 20_channel.go 
```

- [确保 goroutine 完成](21_channnel.go)：main 函数中起了一个 goroutine，通过 channel 确保在 goroutine 执行结束之前 main 函数不会提前退出。这里缓冲与非缓冲 channel 都可以起到相同的效果。

```bash
go run 21_channnel.go
```

- [channel range](22_channel-range.go)：通过 for range 循环处理 channel 中的所有消息。

```bash
go run 22_channel-range.go
```

- [channel selelct](26_channel-select.go)：当有多个 channel 时，通过 select 自动选择先获得信息的 channel 来处理。

```bash
go run 26_channel-select.go
```

- [多核并行](30_multi-process.go)：不展示任何结果，goroutine 还没开始，main 就会退出。

```bash
go run 30_multi-process.go
```

- [多核并行](31_multi-process.go)：每次执行的结果都不相同，在多个 goroutine 还没执行完成的情况下，因为 goroutine 9 完成了，所以 main 就会退出。

```bash
go run 31_multi-process.go
```

- [多核并行](32_multi-process-channel.go)：需要读 channel 10 次，童儿通过 channel 来确认 10 个 routine 都已完成。

```bash
go run 32_multi-process-channel.go
```

- [多核并行](33_multi-process-wg.go)：通过 waitGroup 来确认 10 个 routine 都已完成。

```bash
go run 33_multi-process-wg.go
```

- [处理 Context.Done()](38_channel-select-ctx.go)：通过 context.Done() 来协同各个 goroutine

```bash
go run 38_channel-select-ctx.go
```

- [Mutex锁](41_sync-mutex.go)

```bash
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
