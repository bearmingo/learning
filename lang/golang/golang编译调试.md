# golang编译调试

## 环境变量

Golang有和多环节变量用用于控制程序运行的行为。全部内容可通过`go help environment`查看所有提供的环境变量。

限制Golang运行时P的个数，默认个数与核数一致

```bash
GOMAXPROCS=4
```

监控进程GMP的变化

```bash
GODEBUG=schedtrace=1000
GODEBUG=scheddetail=1,schedtrace=1000
```

## 编译参数

开启逃逸分析日志

```
-gcflags '-m -l'
```

### 跨平台编译
```bash
# win x64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
# linux x64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# max os
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```
## 工具

### pprof - 内存泄漏分析

pprof工具方便对运行中的go程序进行采样分析，支持对多种类型的采样分析
- goroutine - stack traces of all current goroutines
- heap - a sampling of all heap allocations
- threadcreate - stack traces that led to blocking on synchronization primitives
- mutex - stack traces of holders of contented mutexes
- profile - cpu profile
- trace - allows collecting all the profiles for certain duration

现在有很多rpc框架有内置管理模块，允许访问管理端口的`/debug/pprof`对服务进行采样分析。

集成pprof只需要在工程中引入下面的代码

```golang
import _ "net/http/pprof"

...
go func() {
    log.Println(http.ListenAndServe("localhost:9980", nil))
}
...

```

然后使用`go tool pprof`采样

```bash
go tool pprof -seconds=10 -http=:9999 http://localhost:9980/debug/pprof/heap?seconds=30 > heap.out
```

当不能直接访问目标机器时，或者没有安装go，可以使用`curl`工具下载数据，例如

```bash
curl http://localhost:9980/debug/pprof/heap?seconds=30 > heap.out

go tool pprof heap.out
```

反编译

```bash
go tool compile -S main.go
```

## 代码注解


## 其他

不安全的方式安装依赖库

```bash
go get -insecure github.com/xxxx
```