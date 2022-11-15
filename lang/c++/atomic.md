# std::atomic\<T\>

## memory order
**memory order约束的是atomic附近的非atomic的操作行为，而不是atomic自身**
`std::memory_orders`的引入使得用户可以在语言层面对多处理器的环境下多线程共享的行为进行控制，忽略编译器和cpu架构的影响。`memory_order`真正的实现机制是限制了内存操作指令之间的重排。

atmoic有6种memory ordering选项：
- `memory_order_relaxed`
- `memory_order_consume`：这条语句后面所有与这块内存有关的都写操作都无法重排到这个操作之前。
- `momery_order_acquire`：这条语句后面的所有读写操作都无法重排到这个操作之前。
- `memory_order_release`: 这条语句前的所有都写操作都无法呗重排到这个操作之后。
- `memory_order_acq_rel`
- `memory_order_seq_cst`

这6种选项代表三种模型：
- sequentially-consistent ordering: `memory_order_seq_cst`
- acquired-release ordering: `memory_order_consume`、`momery_order_acquire`、`memory_order_release`、`memory_order_acq_rel`
- relaxed ordering: `memory_order_relaxed`

### 例子1
```cpp
#include <iostream>
#include <stdint.h>
#include <atomic>
#include <thread>

std::atomic<bool> x, y;
std::atomic<int> z;

void write_x() {
    x.store(true, std::memory_order_seq_cst); // <-- 1
}

void write_y() {
    y.store(true, std::memory_order_seq_cst); // <-- 2
}

void read_x_then_y() {
    while (!x.load(std::memory_order_seq_cst)) {}
    if (y.load(std::memory_order_seq_cst)) // <-- 3
    ++z;
}

void read_y_then_x() {
    while (!y.load(std::memory_order_seq_cst)) {}
    if (x.load(std::memory_order_seq_cst)) // <-- 4
        ++z;
}  

int main() {
    x = false;
    y = false;
    z = 0;

    std::thread c(read_x_then_y);
    std::thread d(read_y_then_x);
    std::thread a(write_x);
    std::thread b(write_y);

    a.join();
    b.join();
    c.join();
    d.join();

    assert(z.load() != 0); // <-- 5

    printf("finished, %d\n", z.load());

    return 0;
}
```

## 参考
[理解memory order](https://blog.csdn.net/jiang4357291/article/details/110753759)
[内存顺序（Memory Order）](https://zhuanlan.zhihu.com/p/45566448)
[c++ std::atomic类型以及其memory order介绍](https://blog.csdn.net/lqlblog/article/details/53149552)