## LLDB调试



## 常用命令

打印变量。 `p`会返回变量的类型以及命令结果的引用名，`po`只会输出对应的值

```gdb
p <var-name>
po <var-name>
p/x 100
```

查看变量内容（只能查看class数据结构，不能查看struct类型的数据结构）

```
v <var-name>
```

执行代码

```
expression -- self.view.backgroundColor = UIColor.red
```

调用方法

```
call 
```



## 疑难问题记录

使用p查看`Date?`类型的变量，答应的结果一直是`nil`， 原因未知。`po`命令可以正常打印出内容