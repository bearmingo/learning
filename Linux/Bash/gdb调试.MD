# GDB调试

## p

打印指定变量名称的值。可查看以下变量：

- 全局变量
- 静态成员变量
- 局部变量

查看全局变量可以使用`file:variable`或者`function:variable`，例如：

```bash
(gdb) p 'test.c'::g_count
```

> 注：GDB调试优化后的程序时，可能发现某些变量不可访问，或者取值错误。需要在编译时关闭编译优化，gcc可使用`-gstabs`，解决这个问题。

查看数组数据时，使用`p *array@len`形式查看，array是数组的首地址，len是数组的长度。例如：

```bash
(gdb)p array@len
$1 = {1,2,3,4}
```

### 格式化输出

x 按十六进制格式显示变量。
    d 按十进制格式显示变量。
    u 按十六进制格式显示无符号整型。
    o 按八进制格式显示变量。
    t 按二进制格式显示变量。
    a 按十六进制格式显示变量。
    c 按字符格式显示变量。
    f 按浮点数格式显示变量。

可以在输出时指定地址类型，例如：

```bash
(gdb) p (char*)0x23b744a10
$19 = 0x23b744a10 "abc"
```

## x 

查看内存值，语法是：

x/nfu address

n、f、u为可选参数：
1. n 是一个正整数，表示显示内存的长度，也就是说从当前地址向后显示几个地址的内容
2. f 表示显示的格式，参见`print的格式化输出`。如果地址所指的是字符串，那么格式可以是s，如果地十是指令地址，那么格式可以是i。
3. u 表示从当前地址往后请求的字节数，如果不指定的话，GDB默认是4个bytes。u参数可以用下面的字符来代替，

- b表示单字节，
- h表示双字节，
- w表示四字节，
- g表示八字节。

当我们指定了字节长度后，GDB会从指内存定的内存地址开始，读写指定字节，并把其当作一个值取出来。

eg.
```bash
(gdb)x/3uh 0xda3742383
```

## display

设置一个自动显示的值。单步调试时用display观察一个值的变化。
格式：
display  [/i | s] [expression | addr]

/i: 输出格式化为机器指令

可以是用`disable display`和`enable display`来禁用和启用该功能， `info display`查看自动显示的信息设置。



## info

### info args

打印当前的函数的参数名称及其值。

### info locals

打印当前的函数的局部变量及其值。

### info breaks

查看当前的端点。

### info threads

查看运行程序的线程信息

### info registers

查看CPU寄存器值（除了浮点寄存器）

### info all-registers

查看CPU所有寄存器的值，包括浮点寄存器。

### info registers <register name>

查看指定寄存器的值

eg:
```bash
(gdb)info registers $eip
```

### info symbol

### info catch

打印当前函数中的异常出来信息

### info line

查看函数在内存中的开始地址和结束地址。line后面可跟行号、函数名称、文件名：行号、文件名：函数

```bash
(gdb) info line tst.c:func
Line 5 of "tst.c" starts at address 0x8048456 <func+6> and ends at 0x804845d <func+13>.
```

## 设置

### set print address [on | off]

打开地址输出，当程序显示函数信息时，GDB会显出函数的参数地址。

### set print array

打开数组显示，打开后当数组显示时，每个元素占一行，如果不打开的话，每个元素则以逗号分隔。

### set print elements [n]

这个选项主要是设置数组的，如果你的数组太大了，那么就可以指定一个来指定数据显示的最大长度，当到达这个长度时，GDB就不再往下显示了。如果设置为0，则表示不限制。

### set print null-stop

如果打开了这个选项，那么当显示字符串时，遇到结束符则停止显示。这个选项默认为off。

### set print pretty on

如果打开printf pretty这个选项，那么当GDB显示结构体时会比较漂亮。

### set print union

置显示结构体时，是否显式其内的联合体数据。

### set print object

在C++中，如果一个对象指针指向其派生类，如果打开这个选项，GDB会自动按照虚方法调用的规则显示输出，如果关闭这个选项的话，GDB就不管虚函数表了。