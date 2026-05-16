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