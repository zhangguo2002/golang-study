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

2.三个goroutine交替打印abc 10次
```go
package main

import(
	"fmt"
	"sync"
)
func main(){
	//定义3个channel
	ch1:=make(chan struct{})
	ch2:=make(chan struct{})
	ch3:=make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)
	//打印a
	go func(){
		defer wg.Done()
		for i:=0;i<10;i++{
			<-ch1
			fmt.Println("a")
			ch2<- struct{}{}
		}
		//第10次的时候，打印c的goroutine写入了ch1
		//为了防止阻塞，要消费以下ch1
		<-ch1
	}()
	//打印b
	go func(){
		defer wg.Done()
		for i:=0;i<10;i++{
			<-ch2
			fmt.Println("b")
			ch3<- struct{}{}
		}	}()
	//打印c
	go func(){
		defer wg.Done()
		for i:=0;i<10;i++{
			<-ch3
			fmt.Println("c")
			ch1<- struct{}{}
		}
	}()
	//启动
	ch1<-struct{}{}
	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)
	fmt.Println("end")
}
```

3.用不超过10个goroutine不重复的打印slice中100个元素
```go
package main

import(
	"fmt"
	"sync"
)
//用不超过10个goroutine不重复打印slice中的100个元素
//容量为10的有缓冲channel实现
//每次启动10个，累计启动100个goroutine,且无序打印
func main(){
	var wg sync.WaitGroup
	//创建切片
	ss:=make([]int,100)
	for i:=0;i<100;i++{
		ss[i]=i
	}
	ch:=make(chan struct{},10)
	for i:=0,i<100;i++{
		wg.Add(1)
		ch <- struct{}{}
		//写10个就阻塞了，此时goroutine中打印
		go func(idx int){
			defer wg.Done()
			fmt.Printf("val:%d \n",ss[idx])
			<-ch
		}(i)
	}
	wg.Wait()
	//关闭channel
	close(ch)
	fmt.Println("end")
}
//用不超过10个goroutine不重复的打印slice中的100个元素
//创建10个无缓冲channel和10个goroutine
//固定10个goroutine，且顺序打印
func test9(){
	var wg sync.WaitGroup
	//创建切片
	ss:=make([]int,100)
	for i:=0,i<100,i++{
		ss[i]=i
	}
	//创建channel和goroutine
	hashMap:=make(map[int]chan int)
	sort:=make(chan struct{})
	for i:=0;i<10;i++{
		hashMap[i]=make(chan int)
		wg.Add(1)
		go func(idx int){
			defer wg.Done()
			for val:=range hashMap[idx]{
				fmt.Printf("go id:%d,val:%d \n",idx,val)
				sort<-struct{}{}
			}
		}(i)
	}
	//循环切片，对10取模，找到对应channel的key，写入值
	for _,v:=range ss{
		id:=v%10
		hashMap[id]<-v
		//有序
		<-sort
	}
	//循环结束关闭channel，删除map的key
	for k,_:=range hashMap{
		close(hashMap[k])
		close(hashMap,k)
	}
	wg.Wait()
	close(sort)
	fmt.Println("end")
}
```

4.两个协程交替打印奇偶数
```go
package main

import (
	"fmt"
	"time"
)
func main(){
	//golang交替打印奇偶数
	//交替打印，可以通过channel来实现
	chan1:=make(chan struct{})
	//偶数
	go func(){
		for i:=0;i<10;i++{
			chan1<-struct{}{}
			if i%2==0{
				fmt.Println("打印偶数：",i)
			}
		}
	}()
	//奇数
	go func(){
		for i:=0,i<10;i++{
			<-cha1
			if i%2==1{
				fmt.Println("打印奇数:",i)
			}
		}
	}()
	//阻塞
	select{
		case <-time.After(time.Second*10):
	}
}
```