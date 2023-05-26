
## 功能库
github.com/jinzhu/copier
结构体复制

golang.org/x/lint/golint
用于检测代码中规范

[panjf2000/ants](https://github.com/panjf2000/ants)
高性能的goroutine线程池。

go-enum

carbon 
时间转换库

BigCache 
是一个快速，支持并发访问，自淘汰的内存型缓存，可以在存储大量元素时依然保持高性能。BigCache 将元素保存在堆上却避免了 GC 的开销。

[uNetworking/SuperCereal](https://github.com/uNetworking/SuperCereal)
json解析，据说比标准库快17倍

[monochromegane/go-avx](https://pkg.go.dev/github.com/monochromegane/go-avx)
使用AXV指令集的库，前提是CPU需要支持该指令集。

[stretchr/testify](github.com/stretchr/testify)
单元测试工具库

### 网络库
quic-go
是完全用 go 写的 QUIC 协议栈，开发很活跃，已在 Caddy 中使用，MIT 许可，目前看是比较好的方案。

gnet
高性能网络库

hertz
HTTP框架

chi
是一个轻量级、通用和可组合的路由器，用于构建Go HTTP服务。它特别擅长帮助您编写大型REST API服务，这些服务随着项目的增长和变化而保持可维护。chi建立在Go 1.7中引入的新上下文包之上，用于处理处理请求链上的信号、取消和请求范围值等。

### 命令行

cobra
一个命令行程序库，可以用来编写命令行程序。

pflag
命令行参数解析库

cli

### 依赖注入

**google/wire**
**uber-go/dig**
**facebookgo/inject**
**uber-go/fx**

## 框架

### web服务

**gin**

**echo**

**beego**
国产开源的高性能Web框架，让你快速的开发Go Web应用服务，谢大主笔。

**martini**
也是一款不错的Web框架

caddy
快速的，跨平台的HTTP/2 Web服务器。

go micro
go kit
goframe
dubbogo
go-zero
kratos
kitx
hertz


## 工具
[https://github.com/muesli/duf](https://github.com/muesli/duf)
Linux磁盘空间整理工具，代替linux的du命令。

shadowsocks-go
网络代理

delv
调试工具

[misspell/cmd/misspell](github.com/client9/misspell/cmd/misspell)
用于检测代码中的拼写错误

**hub**
一款更便捷使用github的工具，包装并且扩展了git，提供了很多特性和功能，使用和git差不多。

**hugo**
一款极速的静态页面生成器，让你可以很快的搭建个人网站，提供了多套主题可供使用，并且可以自己定制，和NodeJS的Hexo是一样的。

**pholcus**
支持分布式的高并发、重量级爬虫软件，定位于互联网数据采集

### 数据库

**etcd**
一款分布式的，可靠的K-V存储系统，使用简单，速度快，又安全。

influxdb
可伸缩的数据库，使用场景主要用来存储测量数据，事件点击以及其他等实时分析数据，用来做监控性能很不错。

rqlite
轻量级分布式数据库，完全用Go语言实现。

nutsdb
buntdb
嵌入式数据库

codis
一个分布式Redis解决方案,其实就是一个数据库代理，让你在使用Redis集群的时候，就像使用单机版的Redis是一样的，对开发者透明

## 其他

**awesome-go**
这不是一个go项目，他是一个学习go的资料网站，属于著名的awesome系列，里面关于go的资源非常详细。