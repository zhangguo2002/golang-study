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