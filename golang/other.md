
# Golang 

## 定义匿名全局变量的用途

在`grpc`的`serverreflection.go`看到文件里，看到这么一句代码

```golang
var _ GRPCServer = (*grpc.Server)(nil)
```

一直没看明白，网上搜索了才知道含义。用途是检测`grpc.Server`是否实现了`GRPCServer`接口，如果没有实现，那么编译时会报错。