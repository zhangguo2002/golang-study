1.Slice底层结构是怎样的？
slice的底层数据其实也是数组，slice是对数组的封装，它描述一个数组的片段。slice实际上是一个结构体，包含三个字段：长度、容量、底层数组

```go
//runtime/slice.go
type slice struct{
    array unsafe.Pointer //元素指针
    len int //长度
    cap int //容量
}
```

2.Go语言里slice是怎么扩容的？
Go1.17及以前
1、如果期望容量大于当前容量的两倍就会使用期望容量
2、如果当前切片的长度小于1024就会将容量翻倍
3、如果当前切片的长度大于1024就会每次增加25%的容量，直到新容量大于期望容量
Go1.18及以后，引入了新的扩容规则
当原slice容量小于256的时候，新slice容量为原来的2倍；原slice容量超过256，新slice容量newcap=oldcap+(oldcap+3*256)/4

3.从一个切片截取出另一个切片，修改新切片的值会影响原来的切片内容吗？
在截取完之后，如果新切片没有触发扩容，则修改切片元素会影响原切片，如果触发了扩容则不会
```go
package main

import "fmt"

func main(){
    slice:=[]int{0,1,2,3,4,5,6,7,8,9}
    s1:=slice[2:5]
    s2:=s1[2:6:7]

    s2=append(s2,100)
    s2=append(s2,200)

    s1[2]=20

    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(slice)
    //运行结果
    // [2 3 20]
    // [4 5 6 7 100 200]
    // [0 1 2 3 20 5 6 7 100 9]
```


4、Slice作为函数参数传递，会改变原slice吗？
当slice作为函数参数时，因为会拷贝一份新的slice作为实参，所以原来的slice结构并会被函数中的操作改变，也就是说，slice其实是一个结构体，包含了三个成员：len、cap、array并不会变化。但是需要注意的是，尽管slice结构不会变，但是其底层数组的数据如果有修改的话，则会发生变化。若传的是slice的指针，则原slice结构会变，底层数组的数据也会变
```go
package main

func main(){
    s:=[]int{1,1,1}
    f(s)
    fmt.Println(s)
}

func f(s []int){
    //i只是一个副本，不能改变s中元素的值
    // for _,i:=range s{
    //     i++
    // }
    for i :=range s{
        s[i]+=1
    }
}

//程序输出
[2,2,2]
```
果真改变了原始slice的底层数据。这里传递的是一个slice的副本，在f函数中，s只是main函数中s的一个拷贝。在f函数内部，对s的作用并不会改变外层main函数的s的结构

要想真的改变外层slice，只有将返回的新slice赋值到原始slice，或者向函数传递一个指向slice的指针
```go

```