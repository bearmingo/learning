# 内存泄漏分析

## valgrind

## eBPF

Berkeley Packat Filtering

## BBC

eBPF toolkit

## ptrace/ltrace/strace

## /smaps

`/proc/<pid>/smaps` 反应了运行时的进程内存使用状况。系统运行时库、堆、栈均可以看到。先看下每个行的意义：

第一行：

- `开始地址`和`结束地址` 在程序进程空间内的地址。
- `rwxp` 前三个是`rwx`(读、写、可执行)，`-`表示没有这个属性。 最后一个是`p/s`（私有/共享）
- 00000 偏移量。如果这段内存是从文件中映射过来的，则偏移量表示这段内容在文件中的偏移量。如果不是从文件里映射过来的则为0
- 03:03 如果是这段内存是从文件中映射过来的，则表示文件所在设备的主子设备号。
- 12121 如果是这段内存是从文件中映射过来的，这个表示文件号
- /bin/bash 如果是这段内存是从文件中映射过来的，这个表示文件的名称。This field is blank for anonymous mapped regions. There are also special regions with names like [heap], [stack], or [vdso]. [vdso] stands for virtual dynamic shared object. It’s used by system calls to switch to kernel mode.

接下来的几行：

- RSS - Resident Set Size 实际占用的物理内存大小（包含共享库占用的内存）。RSS = Shared_Clean + Shared_Dirty + Private_Clean + Private_Dirty
- PSS - Physical Set Size 实际使用的物理内存大小（按比例包含共享库占用的内存）。比如四个进程共享同一个占内存1000MB的共享库，每个进程算进250MB在Pss。
- Shared_Clean/Shared_Dirty/Private_Clean/Private_Dirty - Shared/Private表示共享和私有。Clean表示内存也没有被更改，发生换页时，不用写回。dirty表示次页有更改，当发生换页时需要写回磁盘。此处这四个值是遍历页表中各个页后得到的。

在程序的不同时期记录进程的内存情况，最后通过比较，能过大致的了解内存变动的位置。从而大致了解内存泄漏的位置。对于内存泄漏的问题此方法不一定管用。

大致流程：

```bash
cat /proc/<pid>/smaps > before.txt
...
cat /proc/<pid>/smaps > now.txt

diff -u before.txt now.txt

gdb -p <pid>

gdb> dump memory dump.file <start-address> <end-address>
gdb> quit

strings dump.file
# 或者使用
hexdump -C dump.file
```

