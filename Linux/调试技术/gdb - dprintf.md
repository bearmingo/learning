## 简介
`dprintf`是gdb的一个动态打印命令。它是一个特殊的断点，可以使用`info break`命令查看。
命令格式：
```
dprintf location,template,expression, [,expression...]
```
`location`: 需要打印消息的位置，可以是函数名称、文件名称:行号等。
`template`: 打印格式，相当于`printf`函数的第一个参数。
`expression`: 格式化消息的数据源。

## 参考
[Dynamic Printf](https://sourceware.org/gdb/current/onlinedocs/gdb/Dynamic-Printf.html)