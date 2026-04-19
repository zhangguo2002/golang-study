1.Go语言的GMP模型是什么？
GMP是Go运行时的核心调度模型
GMP含义：G是goroutine协程；M是machine系统线程，真正干活的；P是processor，逻辑处理器，它是G和M之间的桥梁。它负责调度G

调度逻辑是这样的，M必须绑定P才能执行G。每个P维护一个自己的本地G队列（长度256），M从P的本地队列取G执行。当本地队列空时，M会按优先级从全局队列、网络轮询器、其他P队列中窃取goroutine，这是work-stealing机制