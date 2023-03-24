
# 测试
```bash
# 进行3次性能测试， 每次5秒 
go test -bench . -benchtime=5s -count=3
# 内存性能测试
go test -bench . -benchmem
```

## 性能分析
从服务端获取到性能跟踪文件后，在使用`go`进行性能分析。
```bash
# 实用8083端口，提供性能测试Web服务
go tool pprof -http=":8083" profile
# 终端上进行性能分析
go tool trace mytrace.profile
```
