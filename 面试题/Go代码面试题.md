1.开启100个协程，顺序打印1-1000，且保证协程号1的，打印尾数为1的数字
```go
// 同时开启100个协程(分别为1号协程 2号协程 ... 100号协程，
// 1号协程只打印尾数为1的数字，2号协程只打印尾数为2的数，
// 以此类推)，请顺序打印1-1000整数以及对应的协程号；
func main() {
	s := make(chan struct{})
	//通过map的key来保证协程的顺序
	m := make(map[int]chan int, 100)
	//填充map,初始化channel
	for i := 1; i <= 100; i++ {
		m[i] = make(chan int)
	}
	//开启100个协程，死循环打印
	//go func() { 这个协程不加也可以的
	for i := 1; i <= 100; i++ {
		go func(id int) {
			for {
				num := <-m[id]
				fmt.Println(num)
				s <- struct{}{}
			}
		}(i)
	}
	//}()
	//循环1-1000，并把值传递给匹配的map
	//然后通过s限制循序打印
	for i := 1; i <= 1000; i++ {
		id := i % 100
		if id == 0 {
			id = 100
		}
		m[id] <- i
		//通过s这个来控制打印顺序。每次遍历一次i
		//都通过s阻塞协程的打印，最后打印完毕
		<-s
	}
	time.Sleep(10 * time.Second)
}
```