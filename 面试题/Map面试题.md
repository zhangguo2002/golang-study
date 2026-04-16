1.Go语言Map的底层实现原理是怎样的？
map就是一个hmap的结构。Go Map的底层实现是一个哈希表。它在运行时表现为指向一个hmap结构体的指针，hmap中记录了桶数组指针buckets、溢出桶指针以及元素个数等字段。每个桶是一个bmap结构体，能存储8个键值对tophash，并有指向下一个溢出桶指针overflow。为了内存紧凑，bmap中采用的是先存8个键再存8个值的存储方式。

hmap结构定义：
```go
type hmap struct{
    count int //map中元素个数
    flags uint8 //状态标志位，标记map的一些状态
    B uint8 //桶数以2为底的对数，即B=log_2(len(buckets))，比如B=3,那么桶数为2^3=8
    noverflow uint16 //溢出桶数量近似值
    hash0 uint32 //哈希种子
    buckets unsafe.Pointer //指向buckets数组的指针
    oldbuckets unsafe.Pointer //是一个指向buckets数组的指针，在扩容时，oldbuckets指向老的buckets数组（大小为新buckets数组的一半），非扩容时，oldbuckets为空
    nevacuate uintptr //表示扩容进度的一个计数器，小于该值的桶已经完成迁移
    extra *mapextra //指向mapextra结构的指针，mapextra存储map中的溢出值
}
```


2.Go语言中map的遍历是有序的还是无序的？
Go语言中Map的遍历完全是随机的，并没有固定的顺序。map每次遍历，都会从一个随机值序号的桶，在每个桶中，再从按照之前选定随机槽位开始遍历，所以是无序的。

3.Map如何实现顺序读取？
如果业务上确实需要有序遍历，最规范的做法就是将Map的键（key）取出来放入一个切片（slice）中，用sort包对切片进行排序，然后根据这个有序的切片去遍历map
```go
package main

import(
    "fmt"
    "sort"
)

func main(){
    keyList:=make([]int,0)
    m:=map[int]int{
        3:200,
        4:200,
        1:100,
        8:800,
        5:500,
        2:200,
    }
    for key:=range m{
        keyList=append(keyList,key)
    }
    sort.Ints(keyList)
    for _,key:=range keyList{
        fmt.Println(key,m[key])
    }
}
```

4.Go语言的Map是否是并发安全的？
map不是线程安全的
在查找、赋值、遍历、删除的过程中都会检测写标志，一旦发现写标志置位（等于1），则直接panic。赋值和删除函数在检测完写标志是复位之后，先将写标志位置位，才会进行之后的操作
检测写标志：
```go
if h.flags&hasWriting==0{
    throw("concurrent map writes")
}
```
设置写标志
```go
h.flags!=hasWriting
```

5.Map的key一定是可比较的吗？ 为什么？
Map的key必须要可比较
首先，Map会对我们提供的key进行哈希运算，得到一个哈希值。这个哈希值决定了这个键值对大概存储在哪个位置（也就是哪个桶里）。然而，不同的key可能会产生相同的哈希值，这就是哈希冲突。当多个key被定位到同一个桶里时，Map就没法只靠哈希值来区分它们了。此时，它必须在桶内进行逐个遍历，用我们传入的key和桶里已有的每一个key进行相等（==）比较。这样才能确保我们操作的是正确的键值对。