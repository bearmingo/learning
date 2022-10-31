# Golang 坑

## 使用指针作为方法的receiver

只要值是可寻址的，就可以在值上直接调用指针的方法。即对一个方法，他的receiver是指针足矣。

但不是所有的值都是可以寻址的，比如map类型的元素、通过interface应用的变量

参考[subtle001](subtle001/main.go)

## 更新map字段的值

如果map一个字段的值是struct类型，则服务直接更新该struct的单个字段。

但slice的元素是可以寻址的

代码参考[subtle002](subtle002/main.go)

## nil interface 和 nil interface的值

虽然interface看起来像指针类型，但它不是。inteface类型的变量只有在类型和值均为nil时才为nil

代码参考[subtle003](subtle003/main.go)

## 堆栈变量

你并不总是清楚你的变量是分配到堆还是栈

在C/C++中使用`new`创建的变量总是分配到堆内存上的，但在Go中即使使用`new()`, `make()`来创建的变量，变量的内存分配位置依旧归Go编译器管。

Go编译器会根据变量的大小及其“Escape analysis”的结果来决定变量的存储位置。故能够放心的返回本地变量的地址。这在C/C++中是不行的。

在`go build`或者`go run`,加入`-m`参数，并打印出能够变量分配的位置。

## GOMAXPROCS、Cocurrency（并发）and Parallelism（并行）

Go 1.4及以下的版本，程序只会使用1个可执行上下文，即任何时间最多只会有1个goroutine在执行。

Go 1.5 版本将可可执行上下文设置为`runtime.NumCPU()`返回的逻辑CPU核心数。这个数语系统实际中的CPU逻辑核数是否一致，取决于你的系统分配给程序的核心数，可以使用`GOMAXPROCS`环境变量或者动态的使用`runtime.GOMAXPROCS()`来调整。

误区：`GOMAXPROCS`表示自行goroutine的CPU核心数。

`GOMAXPROCS`的值是可以超过 CPU 的实际数量的，在 1.5 中最大为 256，1.9中为300

## 读写操作的重新排序

Go可能会重排一些执行顺序，可以保证在一个goroutine中的操作是顺序执行的，但不保证多个goroutine的执行顺序。

## 优先调度

程序可能出现一个goroutine在运行时阻止其他goroutine的运行，比如程序中有一不让调度器执行的for循环。

for循环体不必为空，但如果代码不会触发调度器执行，将会出现问题。

调度器在GC、Go声明、阻塞channel、阻塞系统调用和锁操作后再执行，也会在非内联函数调用时执行。

可以使用`runtime.Gosched()`来手动启用调度器。

### 关闭HTTP的响应体

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

## 关闭 HTTP 连接

一些支持 HTTP1.1 或 HTTP1.0 配置了 `connection: keep-alive` 选项的服务器会保持一段时间的长连接。但标准库 “net/http” 的连接默认只在服务器主动要求关闭时才断开，所以你的程序可能会消耗完 socket 描述符。解决办法有 2 个，请求结束后：

- 直接设置请求变量的 Close 字段值为 true，每次请求结束后就会主动关闭连接。
- 设置 Header 请求头部选项 Connection: close，然后服务器返回的响应头部也会有这个选项，此时 HTTP 标准库会主动断开连接。

也可以创建一个自定义配置的 HTTP transport 客户端，用来取消 HTTP 全局的复用连接。

参考代码[subtle005](subtle005/main.go)

根据需求选择使用场景：

- 若你的程序要向同一服务器发大量请求，使用默认的保持长连接。
- 若你的程序要连接大量的服务器，且每台服务器只请求一两次，那收到请求后直接关闭连接。或增加最大文件打开数 `fs.file-max` 的值。

## 将JSON中的数字解码为 interface 类型

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



## Ref

[Golang 58个坑 – 中级篇：36-51](https://www.toutiao.com/article/7159871673905742377/?log_from=2b722b0ff7602_1667058396912)

[Golang 58个坑 – 高级篇：52-58](https://www.toutiao.com/article/7159872096070353419/?app=news_article&timestamp=1667050661&use_new_style=1&req_id=202210292137400101512090661F9D1DCB&group_id=7159872096070353419&share_token=389150A6-9798-480A-BE91-ABE75DCED1DE&tt_from=copy_link&utm_source=copy_link&utm_medium=toutiao_ios&utm_campaign=client_share&source=m_redirect)