# Golang 坑

## 001 使用指针作为方法的receiver

只要值是可寻址的，就可以在值上直接调用指针的方法。即对一个方法，他的receiver是指针足矣。

但不是所有的值都是可以寻址的，比如map类型的元素、通过interface应用的变量

参考[subtle001](subtle001/main.go)

## 002 更新map字段的值

如果map一个字段的值是struct类型，则服务直接更新该struct的单个字段。

但slice的元素是可以寻址的

代码参考[subtle002](subtle002/main.go)

## 003 nil interface 和 nil interface的值

虽然interface看起来像指针类型，但它不是。inteface类型的变量只有在类型和值均为nil时才为nil

代码参考[subtle003](subtle003/main.go)

## 004 堆栈变量

你并不总是清楚你的变量是分配到堆还是栈

在C/C++中使用`new`创建的变量总是分配到堆内存上的，但在Go中即使使用`new()`, `make()`来创建的变量，变量的内存分配位置依旧归Go编译器管。

Go编译器会根据变量的大小及其“Escape analysis”的结果来决定变量的存储位置。故能够放心的返回本地变量的地址。这在C/C++中是不行的。

在`go build`或者`go run`,加入`-m`参数，并打印出能够变量分配的位置。

## 005 GOMAXPROCS、Cocurrency（并发）and Parallelism（并行）

Go 1.4及以下的版本，程序只会使用1个可执行上下文，即任何时间最多只会有1个goroutine在执行。

Go 1.5 版本将可可执行上下文设置为`runtime.NumCPU()`返回的逻辑CPU核心数。这个数语系统实际中的CPU逻辑核数是否一致，取决于你的系统分配给程序的核心数，可以使用`GOMAXPROCS`环境变量或者动态的使用`runtime.GOMAXPROCS()`来调整。

误区：`GOMAXPROCS`表示自行goroutine的CPU核心数。

`GOMAXPROCS`的值是可以超过 CPU 的实际数量的，在 1.5 中最大为 256，1.9中为300

## 006 读写操作的重新排序

Go可能会重排一些执行顺序，可以保证在一个goroutine中的操作是顺序执行的，但不保证多个goroutine的执行顺序。

## 007 优先调度

程序可能出现一个goroutine在运行时阻止其他goroutine的运行，比如程序中有一不让调度器执行的for循环。

for循环体不必为空，但如果代码不会触发调度器执行，将会出现问题。

调度器在GC、Go声明、阻塞channel、阻塞系统调用和锁操作后再执行，也会在非内联函数调用时执行。

可以使用`runtime.Gosched()`来手动启用调度器。

## 007 关闭HTTP的响应体

使用HTTP标准库发起请求，获取响应时，即使你补充响应中读取任何数据或者响应为空，都需要手动关闭响应体。
新手很容易忘记手动关闭，或者写在错误的位置。

在返回错误时，resp也有可能不为空（例如因为重定向导致的错误）。会导致内存泄漏。

参考代码[subtle004](subtle004/main.go)

`resp.Body.Close()`早先版本的实现时读取响应体的数据后丢弃，保证了keep-alive的HTTP连接能够重用处理不止一个请求。但Go的最新版本将读取并丢弃数据的任务交给来调用者。如果不处理，HTTP连接可能会直接关闭而非重用，请参考Go 1.5的文档。

如果需要大量从用Http长连接，需要在不读取完毕Body的逻辑（只读取了部分body数据）中加入：

```golang
_, err := io.Copy(ioutil.Discard, resp.Body) // 手动丢弃读取完毕的数据
```

如果你需要完整读取响应，上边的代码是需要写的。比如在解码 API 的 JSON 响应数据：

```golang
json.NewDecoder(resp.Body).Decode(&data)
```

`Decode`有可能未读取完整的数据？？？？

## 008 关闭 HTTP 连接

一些支持 HTTP1.1 或 HTTP1.0 配置了 `connection: keep-alive` 选项的服务器会保持一段时间的长连接。但标准库 “net/http” 的连接默认只在服务器主动要求关闭时才断开，所以你的程序可能会消耗完 socket 描述符。解决办法有 2 个，请求结束后：

- 直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
- 设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，此时 HTTP 标准库会主动断开连接。

也可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接。

参考代码[subtle005](subtle005/main.go)

根据需求选择使用场景：

- 若你的程序要向同一服务器发大量请求，使用默认的保持长连接。
- 若你的程序要连接大量的服务器，且每台服务器只请求一两次，那收到请求后直接关闭连接。或增加最大文件打开数 `fs.file-max` 的值。

## 008 将JSON中的数字解码为 interface 类型

在 encode/decode JSON 数据时，Go 默认会将数字当作 float64 处理。[subtle006](subtle006/main.go)的examplePanic会造成panic。

如果你尝试 decode 的 JSON 字段是整型，你可以：

- 将 int 值转为 float 统一使用
- 将 decode 后需要的 float 值转为 int 使用(example2)
- 使用 Decoder 类型来 decode JSON 数据，明确表示字段的值类型(example3、example4)
- 使用 struct 类型将你需要的数据映射为数值型（example5)
- 可以使用 struct 将数值类型映射为 json.RawMessage 原生数据类型。适用于如果 JSON 数据不着急 decode 或 JSON 某个字段的值类型不固定等情况(example6)

## 007 struct、array、slice 和 map 的值比较

可以使用相等运算符 == 来比较结构体变量，前提是两个的成员都是可比较类型。

如果两个结构体中有任意成员是不可比较的，将会编译错误。数组成员只有在数组元素可比较时才可比较。

不可比较：

- 函数指针
- map
- []byte

Go提供了一些库函数来比较那些无法使用 == 比较的变量，不如“reflect”包的`DeepEqual`

参考代码[](subtle007/main.go)

这种比较方式可能比较慢，根据你的程序需求来使用。DeepEqual还有其他用法：

```golang
func main() {
    var b1 []byte = nil
    b2 := []byte{}

    fmt.Println("b1 == b2: ", reflect.DeepEqual(b1, b2)) // false
}
```

注意：DeepEqual并不适合用于比较slice

```golang
func code2() {
    var str = "one"
    var in interface{} = "one"
    fmt.Println("str == in: ", reflect.DeepEqual(str, in)) // str == in: true

    v1 := []string{"one", "two"}
    v2 := []string{"two", "one"}
    fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // v1 == v2: false

    data := map[string]interface{}{
        "code":  200,
        "value": []string{"one", "two"},
    }

    encoded, _ := json.Marshal(data)
    var decoded map[string]interface{}
    json.Unmarshal(encoded, &decoded)
    fmt.Println("data == decoded: ", reflect.DeepEqual(data, decoded)) // data == decoded:
}
```

如果要大小写不敏感比较 byte 或者 string 中的英文文本，可以使用 `bytes` 或者 `strings` 包的 `ToUpper()` 和 `ToLower()` 函数。比较其他语言的 byte 或 string，应该使用 `bytes.EqualFold()` 和 `strings.EqualFold()`

如果 byte slice 中含有验证用户身份的数据（密文哈希，token等），不应使用 `reflect.DeepEqual()`、`bytes.Equal()`、`bytes.Compare()`。这三个函数容易对程序造成 *timing attacks*， 此时应该使用 `crypto/subtle` 中包含的 `subtle.ConstantTimeCompare()` 等函数。

- reflect.DeepEqual() 认为空的 slice 与 nil slice 并不相等，但注意 bytes.Equal() 会认为两者相等：

```golang
func code3() {
    var b1 []byte = nil
    b2 := []byte{}

    // b1 与 b2 长度相等、有相同的字节序
    // nil 与 slice 在字节上是相同的
    fmt.Println("b1 == b2: ", bytes.Equal(b1, b2)) // b1 == b2: true
}
```

## 008 从 panic 中恢复

在 defer 延迟执行的函数中调用 recover() ，它便能捕获 中断 panic

```golang
// 错误的 recover 调用示例
func code1() {
    recover()         //
    panic("not good") // 发生 panic，主程序退出
    recover()         // 不会被执行
    println("ok")
}

// 正确的 recover 调用示例
func code2() {
    defer func() {
        fmt.Println("recovered: ", recover())
    }()

    panic("not good")
}
```

从上边可以看出，recover() 仅在 defer 执行的函数中调用才会生效。

```golang
// 错误的调用示例
func code3() {
    defer func() {
        doRecover()
    }()

    panic("not good")
}

func doRecover() {
    fmt.Println("recoverd: ", recover())
}
```

## 009 在 range 迭代 slice、array、map 时通过更新引用来更新元素

在range迭代中，得到的值其实时元素的一份值拷贝，更新拷贝并不会更改原来的元素，即拷贝的地址并不是原有元素的地址。

```golang
func code1() {
    data := []int{1, 2, 3}
    for _, v := range data {
        v *= 10 // data中原有的元素是不会被修改的
    }

    fmt.Println("data: ", data) // data: [1, 2, 3]
}
```

如果要修改的原有元素的值，应该使用索引直接访问：

```golang
func code2() {
    data := []int{1, 2, 3}
    for i, v := range data {
        data[i] = v*10
    }

    fmt.Println("data: ", data) // data: [10, 20, 30]
}
```

如果你的集合保存的是指向值的指针，需稍作修改。依旧需要使用索引访问元素，不过可以使用 range 出来的原则直接更新原有值：

```golang
func code3() {
    data := []*struct{ num int }{{1}, {2}, {3} }
    for _, v := range data {
        v.num *= 10   // 直接使用指针更新
    }

    fmt.Println("data: ", data) // data: &{10} &{20} &{30}
}
```

## 010 slice 中隐藏的数据

从 slice 中重新切出新的 slice 时，新的 slice 会引用原 slice 底层数组。如果跳了这个坑，程序可能会分配大量的临时 slice 来指向原底层数组部分数据，将导致难以预料的内存使用。

```golang
func get() []byte {
    raw := make([]byte, 100000)
    fmt.Println(len(raw), cap(raw), &raw[0])

    return raw[:3]
}

func code1() {
    data := get() // data的容量还是10000，

    fmt.Println(len(data), cap(data), &data[0])
}
```

可以通过拷贝临时的 slice 的数据，而不是重新切片来解决：

```golang
func get2() (res []byte) {
    raw := make([]byte, 100000)
    fmt.Println(len(raw), cap(raw), &raw[0])
    res = make([]byte, 3)
    copy(res, raw[:3])
    return
}

func code2() {
    data := get2()
    fmt.Println(len(data), cap(data), &data[0])
}
```

## 011 Slice 中数据的误用

举个简单的例子，重新文件路径（存储在slice中）

分割路径来指向每个不同级的目录，修改第一个目录名，在重组子目录名，创建新路径：

```golang
func code_error() {
    path := []byte("AAAA/BBBBBBBB")
    sep := bytes.IndexByte(path, '/')
    println("sep: ", sep) // sep:  4

    dir1 := path[:sep]
    dir2 := path[sep+1:]
    println("dir1: ", string(dir1)) // dir1:  AAAA
    println("dir2: ", string(dir2)) // dir2:  BBBBBBBB

    dir1 = append(dir1, "suffix"...)
    println("current dir1: ", string(dir1)) // current dir1:  AAAAsuffix

    path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
    println("dir1: ", string(dir1)) // dir1:  AAAAsuffix
    println("dir2: ", string(dir2)) // dir2:  uffixBBB

    println("new path: ", string(path)) // new path:  AAAAsuffix/uffixBBB
}
```

拼接的结果不是正确的 AAAASuffix/BBBBBBBB ，因为 dir1、 dir2 两个 slice 引用的数据都是 path 的底层数组。
修改了 dir1 的同时也修改了 path， 也导致 dir2 的修改。

解决办法：

- 重新分配新的 slice ，并拷贝你需要的数据。
- 使用完整的 slice 表达式： `input[low:high:max]`，容量便调整为 `max - low`

```golang
// 使用 full slice expression
func code() {
    path := []byte("AAAA/BBBBBBBB")
    sep := bytes.IndexByte(path, '/')
    println("sep: ", sep) // sep:  4

    dir1 := path[:sep:sep] // 此时 cap(dir1) 指定为4， 而不是先前的 16
    dir2 := path[sep+1:]

    path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
    println("dir1: ", string(dir1))     // dir1:  AAAAsuffix
    println("dir2: ", string(dir2))     // dir2:  BBBBBBBB
    println("new path: ", string(path)) // new path:  AAAA/BBBBBBBB
}
```

第7行中的第三个参数时用来控制 dir1 的新容量，再往 dir1 中 append 元素时，
将分配新的 buffer 来保持。而不是覆盖原来的 path 底层数组。

## 012 旧 slice

当你从一个已经存在的 slice 创建新 slice 时，二者的数据指向相同的底层数组。如果你的程序使用这个特性，那就需要注意*旧* slice 问题。

某些情况下，向一个 slice 中追加元素，而它指向的底层数组容量不足时，将会重新分配一个新数组来存储数据。而其他 slice 还指向原来的旧底层数组。

示例代码 [subtle012](subtle012/main.go)

## 013. 类型声明与方法

从一个现有的非 interface 类型创建新类型时，并不回继承原有的方法：

```golang
type myMutex sync.Mutex

func main() {
    var mtx myMutex
    mtx.Lock() // mtx.Lock undefined (type myMutex has no field or method Lock)…
    mtx.UnLock() //
}
```

如果你需要使用原类型的方法，可以将原类型以匿名字段的形式嵌入到你定义的新 struct 中：

```golang
type myLocker struct {
    sync.Mutex
}
```

interface 类型声明也保留它的方法集：

```golang
type myLocker sync.Locker

func main() {
    var locker myLocker
    locker.Lock()
    locker.Unlock()
}
```

## 014 for-switch 和 for select 代码块

没有指定标签的 break 只会跳出 switch/select 语句， 若不能使用 return 语句跳出的话，可以为 break 跳出标签指定的代码块：

```golang
func main() {
loop:
    for {
        switch {
        case true:
            fmt.Println("breaking out...")
            //break    // 死循环，一直打印 breaking out...
            break loop
        }
    }
    fmt.Println("out...)
}
```

goto 虽然也能跳转到指定位置，但依旧会再次进入 for-switch，死循环。

## 015 for 语句的迭代变量和闭包函数

for 语句中的迭代变量在每次迭代中都会重用，即 for 中创建的闭包函数接收到的参数始终是同一个变量，在 goroutine 开始执行时都会得到同一个迭代值。

```golang
func main() {
    data := []string{"one", "two", "three"}

    for _, v := range data {
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
    // 输出： three three three
}
```

最简单的解决方法：无需修改 goroutine 函数，在 for 内部使用局部变量保持迭代值，再传参：

```golang
func main() {
    data := []string{"one", "two", "three"}

    for _, v := range data {
        vCopy := v
        go func() {
            fmt.Println(vCopy)
        }()
    }

    time.Sleep(3 * time.Second)
    // 输出： one two three
}
```

另一种解决方法： 直接将当前的迭代值以参数的形式传递给匿名函数：

```golang
func main() {
    data := []string{"one", "two", "three"}

    for _, v := range data {
        go func(in string) {
            fmt.Println(vCopy)
        }(v)
    }

    time.Sleep(3 * time.Second)
    // 输出： one two three
}
```

注意下面这些复杂的3个示例的区别：

```golang
type field struct {
    name string
}

func (p *field) print() {
    fmt.Println(p.name)
}

// 错误示例
func code_error() {
    data := []field{{"one"}, {"two"}, {"three"}}
    for _, v := range data {
        go v.print()
    }
    
    time.Sleep(3 * time.Second)
    // 输出： three three three
}

func code_right() {
    data := []field{{"one"}, {"two"}, {"three"}}
    for _, v := range data {
        v := v
        go v.print()
    }
    
    time.Sleep(3 * time.Second)
    // 输出： one two three
}

func code_right2() {
    data := []*field{{"one"}, {"two"}, {"three"}}
    for _, v := range data { // // 此时迭代值 v 是三个元素值的地址，每次 v 指向的值不同
        go v.print()
    }
    
    time.Sleep(3 * time.Second)
    // 输出： one two three
}
```

## 016 defer 函数参数的值

对 defer 延迟执行的函数，它的参数会在声明时就会求出具体值，而不是在执行时才求值：

```golang
// 在 defer 函数中参数会提前求值
func main() {
    var i = 1
    defer fmt.Println("result: ", func() int { i * 2}())
    i++
    // result: 2
}
```

## 017 defer 函数的执行时机

对 defer 延迟执行的函数，会在调用它的函数结束时执行，而不是在调用它的语句块结束时执行，注意区分开。

比如在一个长时间执行的函数里，内部 for 循环中使用 defer 来清理每次迭代产生的资源调用，就会出现问题：

[subtle017](subtle017/main.go)

创建

```bash
#!/bin/bash
for n in {1..10000}; do 
    echo content > "file${n}.txt"
done
```

运行后，会异常 too many open file

解决办法： defer 延迟执行的函数写入匿名函数中：

```golang
func main() {
    // ...

    for _, target := range targets {
        func() {
            f, err := os.Open(target)
            if err != nil {
                fmt.Println("bad target: ", target, "error: ", err) // error:too many open files
                break
            }

            defer f.Close() // 在每次 for 语句块结束时，不会关闭文件资源

            // 使用 f 
        }
    }
}
```

当然你也可以去掉 defer， 在文件资源使用完毕后，直接调用 `f.Close()` 来关闭

## 018 失败的类型断言

在类型断言语句中，断言失败则会返回目标类型的“零值”。断言变量与原来变量混用可能会出现异常情况：

```golang

func code_error() {
    var data interface{} = "great"
    
    // data 混用
    if data, ok := data.(int); ok {
        fmt.Println("[is an int], data: ", data)
    } else {
        // 外面的data 被隐藏
        fmt.Println("[not an int], data: ", data) // [isn't a int], data:  0
    }

    fmt.Println("data: ", data) // data:  great
}

// 正确示例
func main() {
    var data interface{} = "great"

    if res, ok := data.(int); ok {
        fmt.Println("[is an int], data: ", res)
    } else {
        fmt.Println("[not an int], data: ", data)    // [not an int], data:  great
    }
}
```

## 019 阻塞的 goroutine 与资源泄漏

在2012年Google I/O 大会上，Rob Pike 的GoConcurrency Patterns 演讲讨论 Go 的几种基本并发模式，如完整代码中从数据集中获取第一条数据的函数：

```golang
func First(query string, replicas []Search) Result {
    c := make(chan Result)
    replicaSearch := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go replicaSearch(i)
    }

    return <-c
}
```

在搜索重复时依旧会每次起一个 goroutine 去处理， 每个 goroutine 都把它的搜索结果发送到结果 channel 中，channel 中收到第一条数据会直接返回。

返回完第一条数据后，其他 goroutine 的搜索结构怎么处理？他们的协程如何处理？

在 First() 中的结果 channel 是无缓冲的，这意味着只有一个 goroutine 能够返回。由于没有 receiver，其他的 goroutine 会在发送上一直柱塞。如果你的大量调用，则可能造成资源泄漏。

为了避免泄露，你应该确保所有的 goroutine都能够正确退出，有两个解决办法。

- 使用带缓冲的 channel, 确保能接收完全部 goroutine 的返回结果。

```golang
func First(query string, replicas []Search) Result {
    c := make(chan Result, len(replicas))
    replicaSearch := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go replicaSearch(i)
    }

    return <-c
}
```

- 使用 select 语句，配合能保存一个缓冲值的 channel default 语句：default 的缓冲 channel 保证了即使结果 channel 收不到数据，也不会柱塞 goroutine。

```golang
func First(query string, replicas []Search) Result {
    c := make(chan Result)
    searchReplica := func (i int) {
        select {
        case c <- replica[i](query):
        default:
        }
    }
    replicaSearch := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go replicaSearch(i)
    }

    return <-c
}
```

- 使用特殊的废弃(cancellation) channel 来中断剩余 goroutine 的执行：

```golang
func First(query string, replicas []Search) Result {
    c := make(chan Result)
    done := make(chan struct{})
    searchReplica := func (i int) {
        select {
        case c <- replica[i](query):
        case <- done:
        }
    }
    replicaSearch := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go replicaSearch(i)
    }

    return <-c
}
```

Rob Pike 为了简化演示，没有提及演讲代码中存在的这些问题。不过对于新手来说，可能会不加思考直接使用。

## 020 左大括号 { 不能单独放一行

在其他大多数语言中，`{` 的位置由你自行决定。Go 比较特别，遵循*分号注入规则(automatic semicolon injection)*:
编译器会在每行代码尾部特定分隔符后加";" 来分隔多条语句，比如噢会在“）”后加分号：

```golang
// 错误示例
func main()
{
    println("aaaaaa")
}

// 等效于
func main();  // 无函数体
{
  println("aaaaaa")
}
```

```bash
./main.go: missing function body
./main.go: syntax error: unexpected semicolon or newline before {
```

## 021 未使用的变量

如果在函数体代码中有未使用的变量，则无法通过编译，不过全局变量声明但不使用是可以的。即使变量声明后为变量赋值，依旧无法通过编译，需在某处使用它：

## 022 未使用的 import

如果你 import一个包，但包中的变量、函数、接口和结构体一个都没有用到的话，将编译失败。可以使用 _下划线符号作为别名来忽略导入的包，从而避免编译错误，这只会执行 package 的 init()

## 023 简短声明的变量只能在函数内部使用

```golang
// 错误示例
myvar := 1
func main() {
}

// 正确示例
var myvar = 1
func main() {
}
```

## 024 使用简短声明来重复声明变量

不能用简短声明方式来单独为一个变量重复声明，:=左侧至少有一个新变量，才允许多变量的重复声明：

```golang
// 错误示例
func main() {  
  one := 0
  one := 1 // error: no new variables on left side of :=
}

// 正确示例
func main() {
  one := 0
  one, two := 1, 2    // two 是新变量，允许 one 的重复声明。比如 error 处理经常用同名变量 err
  one, two = two, one    // 交换两个变量值的简写
}
```

## 0025 不能使用简短声明来设置字段的值

struct 的变量字段不能使用 := 来赋值以使用预定义的变量来避免解决：

```golang
// 错误示例
type info struct {
    result int
}

func work() (int, error) {
    return 3, nil
}

func main() {
    var data info
    data.result, err := work()    // error: non-name data.result on left side of :=
    fmt.Printf("info: %+v\n", data)
}


// 正确示例
func main() {
    var data info
    var err error    // err 需要预声明

    data.result, err = work()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("info: %+v\n", data)
}
```

## 0026 不小心覆盖了变量

对从动态语言转过来的开发者来说，简短声明很好用，这可能会让人误会 := 是一个赋值操作符。如果你在新的代码块中像下边这样误用了 :=，编译不会报错，但是变量不会按你的预期工作：

```golang
func main() {
  x := 1
  println(x)        // 1
  {
    println(x)    // 1
    x := 2
    println(x)    // 2    // 新的 x 变量的作用域只在代码块内部
  }
  println(x)        // 1
}
```

这是 Go 开发者常犯的错，而且不易被发现。可使用 vet工具来诊断这种变量覆盖，Go 默认不做覆盖检查，添加 -shadow 选项来启用：

```bash
> go tool vet -shadow main.go
main.go:9: declaration of "x" shadows declaration at main.go:5
```

注意 vet 不会报告全部被覆盖的变量，可以使用 go-nyet 来做进一步的检测：

```bash
> $GOPATH/bin/go-nyet main.go
main.go:10:3:Shadowing variable `x`
```

## 显式类型的变量无法使用 nil 来初始化

nil 是 interface、function、pointer、map、slice 和 channel 类型变量的默认初始值。但声明时不指定类型，编译器也无法推断出变量的具体类型。

参考：go语言nil：空值/零值

```golang
// 错误示例
func main() {
  var x = nil    // error: use of untyped nil
  _ = x
}

// 正确示例
func main() {
  var x interface{} = nil
  _ = x
}
```

## 028

允许对值为 nil 的 slice 添加元素，但对值为 nil 的 map添加元素则会造成运行时 panic

```golang
// map 错误示例
func main() {
    var m map[string]int
    m["one"] = 1        // error: panic: assignment to entry in nil map
    // m := make(map[string]int)// map 的正确声明，分配了实际的内存
}

// slice 正确示例
func main() {
    var s []int
    s = append(s, 1)
}
```

## 029 map 容量

在创建 map 类型的变量时可以指定容量，但不能像 slice 一样使用 cap() 来检测分配空间的大小：

```golang
// 错误示例
func main() {
  m := make(map[string]int, 99)
  println(cap(m))     // error: invalid argument m1 (type map[string]int) for cap  
}
```

## 030 string 类型的变量值不能为 nil

对那些喜欢用 nil 初始化字符串的人来说，这就是坑：

```golang
// 错误示例
func main() {
  var s string = nil    // cannot use nil as type string in assignment
  if s == nil {    // invalid operation: s == nil (mismatched types string and nil)
    s = "default"
  }
}


// 正确示例
func main() {
  var s string    // 字符串类型的零值是空串 ""
  if s == "" {
    s = "default"
  }
}
```

## 031 Array 类型的值作为函数参数

在 C/C++ 中，数组（名）是指针。将数组作为参数传进函数时，相当于传递来数组内存地址的引用，在函数内部可以改变该数组的值。

在 Go 中，数组时值。作为参数传进函数时，传递的是数组的原始值拷贝，此时在函数内部时服务更新该数组的：

```golang
func main() {
    x := [3]int{1, 2, 3}

    func(arr [3]int) {
        arr[0] = 7
        fmt.Println(arr) // [7 2 3]
    }(x)

    fmt.Println(x) // [1 2 3]
}
```

如果想修改参数数组：

- 直接场地指向这个数组的指针类型：

```golang
func main() {
    x := [3]int{1, 2, 3}

    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) // &[7 2 3]
    }(&x)

    fmt.Println(x) // [7 2 3]
```

- 直接使用slice：即使函数内部得到的时 slice 的值拷贝，但依旧会更新 slice 的元素数据（底层array）

```golang
func main() {
    x := [3]int{1, 2, 3}

    func(arr []int) {
        arr[0] = 7
        fmt.Println(arr) // [7 2 3]
    }(&x)

    fmt.Println(x) // [7 2 3]
```

## 032 range 遍历 slice 和 array  时混淆了返回值

与其他编程语言中的 for-in， foreach 遍历语句不通， Go 中的range在遍历时会生成2个值，第一个时索引，第二个时元素的值：

```golang
func exampleError() {
    x := []string{"a", "b", "c"}
    for v := range x {
        fmt.Println(v) // 1 2 3
    }
}

func exampleRight() {
    x := []string{"a", "b", "c"}
    for _, v := range x {
        fmt.Println(v) // 1 2 3
    }
}
```

## 033 slice 和 array 其实是一位数据

看起来 Go 支持多维的array 和 slice， 可以创建数组的数组，切片的切片，但其实并不是。

对依赖动态计算多位数组的应用来说，就性能和复杂度而言，用 Go 实现效果并不理想。

可以使用原始的一位数组、“独立”切片、“共享底层数据”的切片来创建动态的多维数组。

1. 使用原始的一维数组：要做好索引检查、溢出检查、以及当数组满时在添加值时需要重新做内存分配。

2. 使用“独立”的切片分两步：
   1. 创建外部slice。
   2. 对那个内部slice进行内存分配。注意内部的slice相互独立，使得任一内部 slice 增缩都不会影响到其他的 slice

```golang
func main() {
    x, y := 2, 4

    table := make([][]int, x)
    for i := range(table) {
        table[i] = make([]int, y)
    }
}
```

3. 使用“共享底层数组”的切片
    1. 创建一个存放原始数据的容器 slice。
    2. 创建其他的 slice
    3. 切割原始的slice来初始化其他的 slice

```golang
func main() {
    h, w := 2, 4
    raw := make([]int, h*w)

    for i := range raw {
        raw[i] = i
    }

    // 初始化原生的 slice
    fmt.Println(raw, &raw[4])    // [0 1 2 3 4 5 6 7] 0xc420012120

    table := make([][]int, h)
    for i := range table {
        // 等间距切割原始 slice，创建动态多维数组 table
        // 0: raw[0*4: 0*4 + 4]
        // 1: raw[1*4: 1*4 + 4]
        table[i] = raw[i*w : i*w + w]
    }

    fmt.Println(table, &table[1][0]) // [[0 1 2 3] [4 5 6 7]] 0xc420012120
}
```

## 033 访问 map 中不存在的 key

和其他编程语言类似，如果访问了 map 中不存在的 key 则希望能返回 nil，比如在 PHP 中：

```php
> php -r '$v = ["x"=>1, "y"=>2]; @var_dump($v["z"]);'
NULL
```

Go 则会返回元素对应数据类型的零值，比如 nil、""、 false 和 0，取值操作总有值返回，故不能通过取出来的值来判断 key 是不是在 map 中。

检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可：

```golang
// 错误的 key 检测方式
func main() {
    x := map[string]string{"one": "2", "two": "", "three": "3"}
    if v := x["two"]; v == "" {
        fmt.Println("key two is no entry")  // 键 two 存不存在都会返回的空字符串
    }
}

// 正确示例
func main() {
    x := map[string]string{"one": "2", "two": "", "three": "3"}
    if _, ok := x["two"]; !ok {
        fmt.Println("key two is no entry")
    }
}
```

## 034 string 类型的值是常量，不可更改

尝试使用索引遍历字符串，来更新字符串中个别字符，是不允许的。

string 类型的值是只读的二进制 byte slice，如果真需要修改字符串中的字符，将 string 转为 []byte 修改后，再转为 string 即可：

```golang
// 修改字符串的错误示例

func main() {
    x := "text"
    x[0] = 'T' // error: cannot assign to x[0]
    fmt.Println(x)
}

// 正确示例
func main() {
    x := "text"
    fmt.Println(&x)     // 0xc000010250
    xBytes := []byte(x)
    xBytes[0] = 'T' // 注意此时的T 是 rune 类型
    x = string(xBytes)
    fmt.Println(x) // Text 0xc000010250
}
```

问题：为什么前后的 &x 地址会不会变化

注意： 上面示例并不是更新字符串的正确姿势，因为一个UTF8 编码的字符串可能会占用多个字节，比如汉子就需要3～4个字节来存储，此时更新其中一个字符是错误的。

更新字符串的正确姿势： 将 string 转换为 rune slice （此时 1 个rune 可能会占多个 byte），直接更新 rune 中的字符：

```golang
func main() {
    x := "text"
    xRunes := []rune(x)
    xRune[0] = '我'
    x := string(xRunes)
    fmt.Println(x)
}
```

## 035 string 与 byte slice 之间的转换

当进行 string 和 byte slice 相互转换时，参与转换的时拷贝原始值。这种转换过程，与其他编程语言的强制类型转换操作不同， 也和新 slice 与旧 slice 共享底层数组不同。

Go 在 string 与 byte slice 相互转换上优化来两点，避免额外的内存分配：

- 在 map[string] 中查找 key 时， 使用了对应的 []byte 避免m[string(key)]的内测分配。
- 使用 for-range 迭代 string 转换为 []byte 的迭代：`for i, v := range []byte(str) {...}`

## 036 string 与索引操作符

对字符串用索引返回的不是字符，而是一个 byte 值。

```golang
func main() {
    x := "ascii"
    fmt.Println(x[0])
    fmt.Printf("%T\n", x[0]) // uint8
}
```

如果需要使用 for-range 迭代访问字符串中的字符（unicode code point / rune），标准库中有“unicode/utf8”包来最UTF8相关的解码编码。另外 utf8string 也有像 `func (s *String) At(i int) rune` 等很方便的库函数。

## 036 字符串并不都是 utf8 文本

string 的值不必是 utf8 文本，可以包含任意的值。只有字符串是文字字面值时才是 UTF8 文本，字符串可以通过转义来包含其他数据。

判断字符串是否是UTF8文本，可以使用 "unicode/utf8" 包中的 `ValidString()` 函数。

```golang
func main() {
    str1 := "ABC"
    fmt.Println(utf8.ValidString(str1)) // true

    str2 := "A\feC"
    fmt.Println(utf8.ValidString(str2)) // false

    str3 := "A\\xfeC"
    fmt.Println(utf8.ValidString(str3)) // true     // 把转义字符转义成字面值
}
```

## 037 字符串的长度

在 Python 中：

```python
data = u'♥'
println(len(data)) # 1
```

然而在 Go 中

```golang
func main() {
    c := "♥"
    fmt.Println(len(c)) // 3
}
```

Go 的内建函数 len() 返回的是字符串的 byte 数量，而不像 Python 中那样是计算 Unicode 字符数。

如果想得到字符串的字符数，可以使用 `unicode/utf8` 包中的 `RuneCountInString(str string) (n int)`

```golang
func main() {
    c := "♥"
    fmt.Println(utf8.RuneCountInString(char)) // 1
}
```

注意：`RuneCountInString` 并不总是返回我们看到的字符数，因为有的字符会占用两个 Rune:

```golang
func main() {
    char := "é"
    fmt.Println(len(char))    // 3
    fmt.Println(utf8.RuneCountInString(char))    // 2
    fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合
}
```

## 038 在多行 array、slice、map 语句中缺少","号

## 039 log.Fatal 和 log.Panic 不只是 log

log 标准库提供了不同的日志等级，与其他语言的日子库不同，Go 的 log 包在调用 `Fatal*()`、`Panic*()` 时能做日志外的事情，如中断程序的执行等：

```golang
func main() {
    log.Fatal("Fatal level log: log entry") // 输出信息后，程序终止执行
    log.Panic("Normal level log: log entry")
}
```

## 040 对内建数据结构的操作并不是同步的

尽管 Go 本身有大量的特性来执行并发，但并不保证并发的数据安全，用户需要自己保证变量等数据以原子操作更新。

goroutine 和 channel 是进行原子操作的好方法，或者使用 sync 包中的锁。

## 041 range 迭代 string 得到的值

range 得到的索引是字符值（unicode point/rune）的第一字节的位置，与其他编程语言不同，这个索引并不直接是字符在字符串中的位置。

range string 的索引可能是不连续的。

注意一个字符可能占多个 rune，比如法文单词 café 的 é 。操作特殊字符可使用`norm`包。

for-range 迭代会尝试将 string 翻译成 UTF8 文本，对任何无效的码点都是直接使用 `0XFFFD` rune(�) Unicode 代替字符来表示。如果string中有任何非 UTF8 的数据，应该将 string 保存为 byte slice 再进行操作。

```golang
func main() {
    data := "Aé\xfe\x02\xff\x04"
    // data := "我的"
    fmt.Println(data)
    for i, v := range data {
        fmt.Printf("%d: %#x\n", i, v)
    }

    for i, v := range []byte(data) {
        fmt.Printf("%d: %#x\n", i, v)
    }
}
```

## 042 range 迭代 map

如果你希望以特定的顺序（如按key排序）来迭代 map，要注意*每次迭代都可能产生不一样的结果*

Go 运行时是有意打乱迭代顺序的，所以你得到的迭代结果可能不一致。但也并不总会打乱，得到连续相同5个迭代结果也是有可能的，如：

```golang
func main() {
    m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```

如果你去 Go Playround 重复运行上边的代码，输出是不会变的，只有你更新代码它才会重新编译。重新编译后迭代顺序是被打乱的。

## 043 switch 中的 fallthrough 语句

switch 语句中的 case 代码块会默认带上 break，但可以使用 fallthrough 来强制执行下一个 case 代码块。

```golang
func main() {
    isSpace := func(char byte) bool {
        switch char {
        case ' ': // 空格符会直接 break，返回 false // 和其他语言不一样
            // fallthrough
        case '\t':
            return true
        }
        return false
    }

    fmt.Println(isSpace('\t')) // true
    fmt.Println(isSpace(' ')) // false
}
```

不过你可以在 case 代码块末尾使用 fallthrough，强制执行下一个 case 代码块。

也可以改写 case 为多条件判断：

```golang
func main() {
  isSpace := func(char byte) bool {
    switch char {
    case ' ', '\t':
      return true
    }
    return false
  }
  fmt.Println(isSpace('\t'))    // true
  fmt.Println(isSpace(' '))    // true
}
```

## 044 自增和自减运算

很多编程语言都自带前置后置的 ++、–- 运算。但 Go 特立独行，去掉了前置操作，同时 ++、–- 只作为运算符而非表达式。

```golang
// 错误示例
func main() {
    data := []int{1, 2, 3}
    i := 0
    ++i            // syntax error: unexpected ++, expecting }
    fmt.Println(data[i++])    // syntax error: unexpected ++, expecting :
}

// 正确示例
func main() {
    data := []int{1, 2, 3}
    i := 0
    i++
    fmt.Println(data[i])    // 2
}
```

## 045 按位取反

很多编程语言使用 “~” 作为一元按位取反（NOT）操作符， Go 重用 ^ XOR 操作符来按位取反：

```golang
// 错误的取反操作
func main() {
    fmt.Println(~2)        // bitwise complement operator is ^
}


// 正确示例
func main() {
    var d uint8 = 2
    fmt.Printf("%08b\n", d)        // 00000010
    fmt.Printf("%08b\n", ^d)    // 11111101
}
```

同时 ^ 也是按位异或（XOR）操作符。

一个操作符能重用两次，是因为一元的 NOT 操作 NOT 0x02，与二元的 XOR 操作 0x22 XOR 0xff 是一致的。

Go 也有特殊的操作符 AND NOT &^ 操作符，不同位才取1。

```golang
func main() {
  var a uint8 = 0x82
  var b uint8 = 0x02
  fmt.Printf("%08b [A]\n", a)
  fmt.Printf("%08b [B]\n", b)

  fmt.Printf("%08b (NOT B)\n", ^b)
  fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

  fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
  fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
  fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
  fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}
```

## 046 运算符的优先级

除了位清除（bit clear）操作符，Go 也有很多和其他语言一样的位操作符，但优先级另当别论。

```golang
func main() {
  fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n", 0x2&0x2+0x4)    // & 优先 +
  //prints: 0x2 & 0x2 + 0x4 -> 0x6
  //Go:    (0x2 & 0x2) + 0x4
  //C++:    0x2 & (0x2 + 0x4) -> 0x2

  fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n", 0x2+0x2<<0x1)    // << 优先 +
  //prints: 0x2 + 0x2 << 0x1 -> 0x6
  //Go:     0x2 + (0x2 << 0x1)
  //C++:   (0x2 + 0x2) << 0x1 -> 0x8

  fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n", 0xf|0x2^0x2)    // | 优先 ^
  //prints: 0xf | 0x2 ^ 0x2 -> 0xd
  //Go:    (0xf | 0x2) ^ 0x2
  //C++:    0xf | (0x2 ^ 0x2) -> 0xf
```

优先级列表：

```txt
Precedence    Operator
   5             *  /  %  <<  >>  &  &^
   4             +  -  |  ^
   3             ==  !=  <  <=  >  >=
   2             &&
   1             ||
```

## 047 不导出的 struct 字段无法被 encode

以小写字母开头的字段成员是无法被外部直接访问的，所以 struct 在进行 json、xml、gob 等格式的 encode 操作时，这些私有字段会被忽略，导出时得到零值：

```golang
func main() {
  in := MyData{1, "two"}
  fmt.Printf("%#v\n", in)    // main.MyData{One:1, two:"two"}

  encoded, _ := json.Marshal(in)
  fmt.Println(string(encoded))    // {"One":1}    // 私有字段 two 被忽略了

  var out MyData
  json.Unmarshal(encoded, &out)
  fmt.Printf("%#v\n", out)     // main.MyData{One:1, two:""}
}
```

## 048 程序退出时还有 goroutine 在执行

程序默认不等所有 goroutine 都执行完才退出，这点需要特别注意：

```golang
// 主程序会直接退出
func main() {
  workerCount := 2
  for i := 0; i < workerCount; i++ {
    go doIt(i)
  }
  time.Sleep(1 * time.Second)
  fmt.Println("all done!")
}

func doIt(workerID int) {
  fmt.Printf("[%v] is running\n", workerID)
  time.Sleep(3 * time.Second)        // 模拟 goroutine 正在执行 
  fmt.Printf("[%v] is done\n", workerID)
}
```

常用解决办法：使用 “WaitGroup” 变量，它会让主程序等待所有 goroutine 执行完毕再退出。

如果你的 goroutine 要做消息的循环处理等耗时操作，可以向它们发送一条 kill 消息来关闭它们。或直接关闭一个它们都等待接收数据的 channel：

```golang
// 等待所有 goroutine 执行完毕
// 进入死锁
func main() {
  var wg sync.WaitGroup
  done := make(chan struct{})

  workerCount := 2
  for i := 0; i < workerCount; i++ {
    wg.Add(1)
    go doIt(i, done, wg)
  }

  close(done)
  wg.Wait()
  fmt.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}, wg sync.WaitGroup) {
  fmt.Printf("[%v] is running\n", workerID)
  defer wg.Done()
  <-done
  fmt.Printf("[%v] is done\n", workerID)
}
```

看起来好像 goroutine 都执行完了，然而报错：

    fatal error: all goroutines are asleep – deadlock!

为什么会发生死锁？goroutine 在退出前调用了 wg.Done() ，程序应该正常退出的。

原因是 goroutine 得到的 “WaitGroup” 变量是 var wg WaitGroup 的一份拷贝值，即 doIt() 传参只传值。所以哪怕在每个 goroutine 中都调用了 wg.Done()， 主程序中的 wg 变量并不会受到影响。

```golang
func main() {
  var wg sync.WaitGroup
  done := make(chan struct{})
  ch := make(chan interface{})

  workerCount := 2
  for i := 0; i < workerCount; i++ {
    wg.Add(1)
    go doIt(i, ch, done, &wg)    // wg 传指针，doIt() 内部会改变 wg 的值
  }

  for i := 0; i < workerCount; i++ {    // 向 ch 中发送数据，关闭 goroutine
    ch <- i
  }

  close(done)
  wg.Wait()
  close(ch)
  fmt.Println("all done!")
}

func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
  fmt.Printf("[%v] is running\n", workerID)
  defer wg.Done()
  for {
    select {
    case m := <-ch:
      fmt.Printf("[%v] m => %v\n", workerID, m)
    case <-done:
      fmt.Printf("[%v] is done\n", workerID)
      return
    }
  }
}
```

## 049 向无缓冲的 channel 发送数据，只要 receiver 准备好了就会立刻返回

只有在数据未被 receiver 处理时，sender 才会阻塞。当 receiver 开始接收过数据后, 发送消息不会被阻塞 。因运行环境而异，在 sender 发送完数据后，receiver 的 goroutine 可能没有足够的时间处理下一个数据。如：

```golang
func main() {
  ch := make(chan string)

  go func() {
    for m := range ch {
      fmt.Println("Processed:", m)
      time.Sleep(1 * time.Second)    // 模拟需要长时间运行的操作
    }
  }()

  ch <- "cmd.1"
  ch <- "cmd.2" // 不会被接收处理
}
```

## 050 向已经关闭的 channel 发送数据会造成 panic

从已关闭的 channel 接收数据时安全的。

接收状态值 ok 是 false 时 表明 channel 中已经没有数据可以接收了。类似的，从有缓冲的 channel 中接收数据，缓存的数据获取完再没有数据可以接收时，状态也是 false

向已关闭的 channel 中发送数据，会造成 panic。

可以使用一个 channel done 来告诉剩余的 goroutine 无需向 channel 发送数据。此时 `<- done` 的结果是 `{}`:

```golang
func main() {
    ch := make(chan int)
    done := make(chan struct{})

    for i := 0; i < 3; i++ {
        go func(idx int) {
            select {
            case ch <- (idx + 1) * 2:
                fmt.Println(idx, "Send result")
            case <- done:
                fmt.Println(idx, "Existing")
            }
        }(i)
    }

    fmt.Println("Result: ", <- ch)
    close(done)
    time.Sleep(3 * time.Second)
```

## 051 使用了值为 nil 的channel

在一个值为 nil 的 channel 三发送的接收数据将永久性阻塞。

```golang
func main() {
    var ch chan int // 未初始化
    for i := 0; i < 3; i++ {
        go func(i int) {
            ch <- i
        }(i)
    }

    fmt.Println("Result: ", <- ch)
    time.Sleep(2 * time.Second)
}
```

利用这个死锁的特性，可以在 sleep 中动态的打开和关闭 case 语句块：

```golang
func main() {
    inCh := make(chan int)
    outCh := make(chan int)

    go func() {
        var in <-chan int = inCh
        var out chan<- int
        var val int

        for {
            select {
            case out <- val:
                fmt.Println("--------")
                out = nil
                in = inCh
            case val = <- in:
                fmt.Println("+++++++++")
                out = outCh
                in = nil
            }
        }
    }()

    go func() {
        for r := range outCh {
            fmt.Println("Result: ", r)
        }
    }()

    time.Sleep(0)
    inCh <- 1
    inCh <- 2
    time.Sleep(3 * time.Second)
}
```

## 052 若函数 receiver 传参是传值的方式，则无法修改参数的原有值

方法 receiver 的参数与一般的函数的参数类似： 如果声明为值，那方法体得到的是一份参数的值拷贝，此时对参数的任何修改都不会对原有值产生影响。

除非 receiver 参数是 map 或者 slice 类型的变量，并且是以指针方式更新 map 中的字段、slice的元素，才会更新原有值。

## Ref

[Golang 58个坑 – 初级篇：1-35](https://www.toutiao.com/article/7159478735492710924/?log_from=e88ffc9c1968c8_1667322172313)
[Golang 58个坑 – 中级篇：36-51](https://www.toutiao.com/article/7159871673905742377/?log_from=2b722b0ff7602_1667058396912)

[Golang 58个坑 – 高级篇：52-58](https://www.toutiao.com/article/7159872096070353419/?app=news_article&timestamp=1667050661&use_new_style=1&req_id=202210292137400101512090661F9D1DCB&group_id=7159872096070353419&share_token=389150A6-9798-480A-BE91-ABE75DCED1DE&tt_from=copy_link&utm_source=copy_link&utm_medium=toutiao_ios&utm_campaign=client_share&source=m_redirect)