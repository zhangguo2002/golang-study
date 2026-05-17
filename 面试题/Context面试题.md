1.Go语言里的context是什么？
go语言里的context实际上是一个接口，提供了Deadline()、Done()、Err()、Value()四种方法。它在Go1.7标准库被引入。
它本质上是一个信号传递和范围控制的工具。它的核心作用是在一个请求处理链路中，优雅的传递取消信号、超时和截止日期，并能携带一些范围内的键值对数据。
```go
type Context struct{
    Deadline()(deadline time.Time,ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
2、Go语言的Context有什么作用？
Go的Context主要解决三个核心问题：超时控制、取消信号传播和请求级数据传递
在实际项目中，我们最常用的是超时控制。比如一个HTTP请求需要调用多个下游服务，我们通过context.WithTimeout设置整体超时时间，当超时发生时，所有子操作都会收到取消信号并立即退出，避免资源浪费。取消信号的传播是通过Context的层级结构实现的，父Context取消时，所有子Context都会自动取消。
另外Context还能传递请求级的元数据，比如用户ID、请求ID等，这在分布式链路追踪中特别有用。需要注意的是，Context应该作为函数的第一个参数传递，不要存储在结构体中，并且传递的数据应该是请求级别的，不要滥用。

3、Context.Value的查找过程是怎样的
Context.Value的查找过程是一个链式递归查找的过程，从当前Context开始，沿着父Context链一直向上查找直到找到对应的key或到达根Context。
具体流程是：当前调用ctx.Value(key)时，首先检查当前Context是否包含这个key,如果当前层没有，就会调用parent.Value(key)继续向上查找。这个过程会一直递归下去，直到找到匹配的key返回对应的value，或者查找到根Context返回nil。

4、Context如何被取消
Context的取消是通过channel关闭信号实现的，主要有三种取消方式。
首先是主动取消，通过context.WithCancel创建的Context会返回一个channel函数，调用这个函数就会关闭内部的done channel，所有监听这个Context的goroutine都能通过ctx.Done()收到取消信号。
其次是超时取消，context.WithTimeout和context.WithDeadline会启动一个定时器，到达指定时间后自动调用cancel函数触发取消
最后是级联取消，当父Context被取消时，所有子Context会自动被取消，当父Context被取消时，所有子Context会自动被取消，这是通过Context树的结构实现的。