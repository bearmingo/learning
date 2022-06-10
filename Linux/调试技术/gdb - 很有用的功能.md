## .gdbinit
gdb启动时，默认会加载用户目录下的.gdbinit文件。可以在这个文件中添加一些gdb的默认配置。例如
```
set print pretty on
```

## 断点的保存和加载
调试程序时，我们需要经常性的重启gdb。这时如果将常用的断点保存到文件中，下次在加载进来，可以极大的加快调试效率。

保存： `save breakpoints filename`
加载：`source filename`

## 调试宏
`-g`选项编译是无法查看宏的，可以通过`-g3`编译，就可以很轻松的查看宏定义和宏展开聊。
可以通过`p`打印宏，也可以通过例如`macro expand MAX(5. 6)`对宏进行展开。

## 条件断点
```
break [LOCATION] if CONDITION
```

## 断点命令
```gdb
break somebreak
command 1
```

## gdb启动并执行命令
```
gdb --batch -ex `p a` -p <pid>
```
## 转存一段内存到文件
```
(gdb) dump memory filename [START addr] [end addr] 
```
保存下来后，可以使用`hexdump -C`来查看二进制文件 

## 自定义命令
```gdb
(gdb) define [name]

输入多行命令

end # 结束
```
参数预定义变量`$arg0`, `$arg1`...

sizeof: 获取

## 执行shell命令
```gdb
(gdb) shell [command]
```

## 反向调试命令
-   reverse-step，倒退到上一个源代码行
-   *reverse-next，它倒退到上一个源代码行，向后跳过函数调用
-   reverse-finish，倒退到当前函数即将被调用的时刻
-   reverse-continue，它返回到程序中的先前状态，该状态将（现在）触发断点（或其他导致断点停止的状态）