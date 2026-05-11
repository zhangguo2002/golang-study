1、什么是CSP?
CSP(通信顺序进程)并发编程模型，它的核心思想是：通过通信共享内存，而不是通过共享内存通信。Go语言的Goroutine和Channel机制，就是CSP的经典实现，具有以下特点：
a.避免共享内存：协程（Goroutine）不直接修改变量，而是通过channel通信
b.天然同步：channel的发送/接受自带同步机制，无序手动加锁
c.易于组合：channel可以嵌套使用，构建复杂并发模式（如管道、超时控制）

2、channel的底层实现原理是怎么样的？
Channel的底层是一个名为hchan的结构体，核心包含几个关键组件：
环形缓存区：有缓冲channel内部维护一个固定大小的环形队列，用buf指针指向缓冲区，sendx和recvx分别记录发送和接收的位置索引。这样设计能高效利用内存，避免数据搬移
两个等待队列sendq和recvq:用来管理阻塞的goroutine。sendq存储因channel满而阻塞的发送者，recvq存储因channel空而阻塞的接收者。这些队列用双向链表实现，当条件满足时会唤醒对应的goroutine
互斥锁:hchan内部有个mutex，所有的发送、接收操作都需要先获取锁，用来保证并发安全。虽然看起来可能影响性能，但Go的调度器做了优化，大多数情况下锁竞争并不激烈。
hchan定义如下
```go
type hchan struct{
    //chan 里元素数量
    qcount int
    //chan 底层循环数组长度
    dataqsiz uint
    //指向底层循环数组的指针
    //只针对有缓冲的channel
    buf unsafe.Pointer
    //chan中元素大小
    elemsize uint16
    //chan 是否被关闭的标志
    closed uint32
    //chan 中元素类型
    elemtype *_type //element type
    //已发送元素在循环数组中的索引
    sendx uint //send index
    //已经接收元素在循环数组中的索引
    recvx uint //receive index
    //等待接收的goroutine队列
    recvq waitq //list of recv waiters
    //等待发送的goroutine队列
    sendq waitq //list of send waiters
}
```

3、向channel发送数据的过程是怎么样的？
向channel发送数据的整个过程都会在mutex保护下进行，保证并发安全。会经历几个关键步骤：
a、首先是检查是否有等待的接收者。如果recvq队列不为空，说明有goroutine在等待接收数据，这时会直接把数据传递给等待的接收者，跳过缓冲区，这是最高效的路径。同时会唤醒对应的goroutine继续执行。
b、如果没有等待接收者，就尝试写入缓存区。检查缓存区是否还有空间，如果qcount<dataqsiz，就把数据复制到buf[sendx]的位置，然后更新sendx索引和qcount计数。这是无缓冲或缓存区未满时的正常流径。
c、当缓存区满了就需要阻塞等待。创建一个sudog结构体包装当前goroutine和要发送的数据，加入到sendq等待队列中，然后调用gopark让当前goroutine进入阻塞状态，让出CPU给其它goroutine。
被唤醒后继续执行。当有接收者从channel读取数据后，会从sendq中唤醒一个等待的发送者，被唤醒的goroutine会完成数据发送并继续执行
还有个特殊情况是向已关闭的channel发送数据会直接panic。这是Go语言的设计原则，防止向已关闭的通道写入数据
```go
package main

import(
    "fmt"
    "time"
)
func goroutineA(a <-chan int){
    val:=<-a
    fmt.Println("goroutine A received data:",val)
    return
}
func goroutineB(b <-chan int){
    val:=<-b
    fmt.Println("goroutine B received data:",val)
    return
}
func main(){
    ch:=make(chan int)
    go goroutineA(ch)
    go goroutineB(ch)
    ch<-3
    time.Sleep(time.Second)
    ch1:=make(chan struct{})
}
```