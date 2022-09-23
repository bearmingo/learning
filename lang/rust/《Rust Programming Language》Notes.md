来源：[Rust Programming Language]([Variables and Mutability - The Rust Programming Language](https://doc.rust-lang.org/book/ch03-01-variables-and-mutability.html)

# 编程观念

## 变量和可变性

1. Rust中变量默认是不可修改的。

2. 变量声明使用`let`。

3. 可修改的变量声明需要加`mut`，例如 `let mut a = 5;`

4. 变量可以被隐藏，可以是不同的变量类型。这点和C\C++是有区别的，C\C++隐藏只能发生在子作用于中。

5. 可变变量的类型不可被修改。

### 数据类型

数据类型分为：标量<mark>（英文：scaler，不太清楚标准的翻译是什么）</mark>和复合类型

#### 标量

标量有：整型、浮点、布尔和字符。

##### 整型

整型分有符号和无符号两种。无符号的以`u`开始；有符号的以`i`开始。整型的长度有8、16、32、64、128和arch（`isize`、`usize`)几种。其中arch在与运行平台架构有关。

常量书写时，可以使用`_`作为分割符，用于提高可读性（例如：`1_00000`)。

整型运算溢出，在debug模式下，程序会报panic。release模式下会忽略（与C/C++的处理方式一致）。可以使用标准库里的函数（wrapping_\*, check_\*, overflowing_\*, saturating_\*）处理溢出。

##### 字符

rust的字符长度时4个字节的Unicode标量。但Rust字符并不是真正的Unicode中字符的概念<mark>（两者有所区别，具体我还不知道）</mark>。

#### 复合类型

复合类型分：组元和数组。

##### 组元

组合的元素数量是不可变的，不同的元素的类型可以是不同的。声明方式：

```rust
let tup: (i32, f64, u8) = (500, 6.4, 1);
```

使用`tup.1`这种方式提取其中的元素。也可以使用下面的方式提取元素

```rust
let (x, y, z) = tup;
```

##### 数组类型

数组类型中有array和vector两种实现。它们的数据是存储在堆上面。

数组在声明时可以指定数据类型和数组大小：

```rust
let a = [u32: 5]
```

声明时也可以指定数组大小和初始化值：

```rust
let a = [3; 5] // [3, 3, 3, 3, 3]
```

## 函数

函数命名惯例：`snake case`，

函数格式如下：

```rust
fn function_example(x: i32) -> i32 {
    x * 5
}
```

返回值的方式与C/C++有区别，在C/C++中，使用return指定函数返回的值。在Rust中最后一行代码不带有`;`结束，就表示这个是返回值（上面例子中的`x * 5`就是`funcation_example`的返回值）。

Rust函数体是有声明和表达式（表达式不一定要有）。Rust中区分声明和表达式。

声明：执行某些操作但不返回值的指令。

表达式：计算结果值。

## 注释

只使用`//`, 没有C/C++中的多行注释`/*`。

## 控制流

### if

if的英文解释：“If this condition is met, run this block of code, If the condition is not met, do not run this block of code”。这个解释简单明了。

与if的条件关联的代码有时被称为包裹条件（arms）。

if条件中的值必须是布尔型，否则编译报错。Rust不像C/C++，自动将非布尔型转换为布尔型。

if的条件内容可以不使用括号包裹，像Golang一样书写。

在声明中可以使用if，例如：

```rust
let condition = true;
let number = if condition { 5 } else { 6 };
```

### 循环重复

Rust有三种循环：loop、while、for。

#### loop

loop可以可以在循环中返回一个值，例如：

```rust
let mut counter = 0;
let result = loop {
    counter += 1;
    if counter == 10 {
        break counter * 2;
    }
}
```

多个loop嵌套时，可以使用标签区分不同的loop。在break时，可以指明结束的loop。这点比C/C++中`while (true)`的break方便。用法：

```rust
let mut count = ;
'counting_up': loop {
    let mut remaining = 10;
    
    loop {
        if remaining == 9 {
            break; // 这里默认结束当前的循环
        }
        if count == 2 {
            break 'counting_up'; // 这里指明结束的时外面的counting_up循环
        }
        remaining -= 1;
    }
    count += 1;
}
```

#### for

for的书写方式：

```rust
for e in （1..4) {
    
}
```

## 所有权
Rust通过一些列编译器的规则来管理内存。编译器在编译时，检查有是否违反规则。所有全的特性不会使得程序运行变慢。

所有权规则：
- 每一个值都有一个所有者。
- 同时只能有一个所有者。
- 但所有者超出作用域，该值将被删除（Dropped）。

拷贝&转移：
- 栈上的值只拷贝。
- 堆上的值默认转移。深拷贝需要显示调用`clone`方法。

	shallow copying: 浅层拷贝
	deep copying: 深层拷贝

如果一个类型（或者任何他的一部份实现了`Drop`特性，Rust将不允许我门在这个类型上使用`Copy`注解。

如果组元只含有实现Copy的类型，那么组元也是`Copy`特性。

### 引用和借用
	Reference & Borrowing

```rust
let s1 = String::from("Hello World")
let leng = calcualte_length(&s1)

fn calculate_lenght(s: &String) ->usize {
	s.len()
}
```

使用语法`&s1`，创建对`s1`的引用，但没有拥有他。

借用(Borrowing): 创建引用的动作叫做借用。

#### 可修改引用
普通引用是不可修改的，创建可修改的引用，需要添加修饰符`mut`，例如：`&mut s`。函数参数需要一个可修改引用作为参数，声明中也需需要带`mut`，例如：`value: &mut String`。

```rust
let mut s = String::from("hello")
change(&mut s)
...
fn change(value: &mut String) {
	value.push_str(", World")
}
```

可修改应用的限制：
1. 当你有一个值的可修改应用，你就不能有其他对这个值的应用。
2. 当你有一个不可修改的引用时，你也不能创建可修改的引用。

		产生数据竞争的三种情况：
		1. 两个或两个以上的指针同时访问同一个数据。
		2. 至少以一个指针在用于修改数据。
		3. 没有使用任何同步访问数据的机制。

#### 悬空引用
在Rust中，编译器会确保不产生一个悬空引用。

#### 引用的规则
1. 在任何时间里，你都可以创建一个可变引用或者任意个数的不可变引用。
2. 引用一定是有效的。

### 切片
英文名称叫：slice

切片能够让你引用集合中连续的元素序列，而不是引用整个集合。它也是一种引用，但他没有所有权。

对字符串创建切片：
```rust
let s = String::from("Hello world");
let len = s.len();

let slice = &s[0..2];
let slice = &s[..2]; // 与&s[0..2]相同

let slice = &s[3..len];
let slice = &s[3..]; // 与&s[3..len]相同

let slice = &s[0..len];
let slice = &s[..]; // 与&s[0..len]相同
```

	使用`&str`作为参数的函数比使用`&String`的通用性更强。


