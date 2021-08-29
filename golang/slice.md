# slice

一个切片数据结构中有三个成员：内存地址指针、长度、容量。



当向切片中添加元素(append)时

- 如果容量满足，直接添加到最后。`append`返回的`slice`的地址支持不变。
- 如果容量不够，会新分配一个更大的内存空间，并从旧的`slice`拷贝到新的`slice`中，再添加追加新的元素。`append`返回的`slice`的地址不同。



## 常用操作

### 初始化

```go
// 第一个参数是类型，
// 第二个参数表示初始时的slice中元素的个数，
// 第三个参数表示初始化时slice的容量。在容量可以预估的情况下使用该值，可以有效减少slice重新分配的次数
x := make([]int, 0, 100) 

x := []int{1, 2, 3}
```

### 追加

```go
x = append(x, 1)
```

### 复制

```go
x := []int{1, 2, 3}

a := make([]int, 0, 3)
copy(y, x)

b := append([]int, x...)
```

### 删除

```go
x := []int{1, 2, 3, 4, 5}
// 删除部分第二位和第三位
x := append(x[:1], x[3:])

// 删除其中一位
x := append(x[:i], x[i+1:])
```

### 限制新切片的容量

在创建切片时，使用第三个索引选项引可以用来控制新切片的容量。其目的并不是要增加容量，而是要限制容量。允许限制新切片的容量为底层数组提供了一定的保护，可以更好地控制追加操作。

```go
fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
// slice[i:j:k]
// 第一个参数指定新切片的开始位置，即Plum的位置
// 第二个参数指定新切片的结束位置(不包含)，即Banana的位置
// 第三个参数指定新切片的容量的结束位置，即Grape的位置，新切片的容量是2。如果没有给定，则表示切到底层数组的最尾部
myFruit := fruit[2:3:4]

myFruit = append(myFruit, "Lemon")
fmt.Printf("%v\n", fruit)
// print: [Apple Orange Plum Lemon Grape]
fmt.Printf("%v\n", myFruit)
// print: [Plum Lemon]
```

*在myFruit中添加元素后，fruit中的Banna也变成了Lemon。可以看出myFruit和fruit是使用了同地址空间。*

```go
a := x[:0:0] // 没有元素，且容量空0的切片
```

## slice 的内存泄露

```go
x := []*int{new(int), new(int)}
x := x[:1]
```

这个例子中，新的x和旧的slice指向同一个地址空间。新的x长度变成了1，不能够看到第二个值，但它仍然指向*int的。指针的被引用着，不能够被gc释放。只要x一直不释放，第二个值就不能被释放。