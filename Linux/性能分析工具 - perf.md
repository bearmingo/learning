# 性能分析工具 -- perf

perf是用来软件性能分析的工具

## perf list

Perf 有三大类采样点事件，可以使用`perf list`命令查看所有的采样点的事件。

- Hardware Event: 是由PMU硬件产生的事件，例如cache命中。
- Software Event: 是由内核软件产生的事件，例如进程切换。
- Tracepoint Event: 是由内核中静态tracepoint所出发的事件，这些tracepoint可用来判断程序运行期间内核的行为细节，例如slab分配器的分配次数等。

## perf stat

程序运行慢，如果是计算量太大，多数事件使用在cpu计算上，这种叫做CPU Bound型。如果是IO过多，但cpu使用率不高，这种叫做IO bound型。

`perf stat`可用于查看程序运行时的各种事件的汇总数据。可以大致给出一个程序的优化方向，确定是对CPU bound调优还是IO Bound调优。

命令：

```bash
$ perf stat </path/to/exe>
```

- Task-clock-msecs：CPU 利用率，该值高，说明程序的多数时间花费在 CPU 计算上而非 IO。
- Context-switches：进程切换次数，记录了程序运行过程中发生了多少次进程切换，频繁的进程切换是应该避免的。
- Cache-misses：程序运行过程中总体的 cache 利用情况，如果该值过高，说明程序的 cache 利用不好
- CPU-migrations：表示进程 t1 运行过程中发生了多少次 CPU 迁移，即被调度器从一个 CPU 转移到另外一个 CPU 上运行。
- cycles：处理器时钟，一条机器指令可能需要多个 cycles，
- instructions: 机器指令数目。
- IPC：是 Instructions/Cycles 的比值，该值越大越好，说明程序充分利用了处理器的特性。
- Cache-references: cache 命中的次数
- Cache-misses: cache 失效的次数。
- branche-misses: cpu分支预测失败次数。

通过指定 -e 选项，您可以改变 perf stat 的缺省事件



## perf top

Perf top 用于实时显示当前系统的性能统计信息。该命令主要用来观察整个系统当前的状态，比如可以通过查看该命令的输出来查看当前系统最耗时的内核函数或某个用户进程。



## perf record



```bash
perf record – e cpu-clock ./t1
perf report

# -g 选项便可以得到需要的信息
perf record – e cpu-clock – g ./t1
```

## perf sched

Linux 调度器数据收集，perf sched有五个子命令：

- perf sched record: low-overhead recording of arbitrary workloads(低负载的开销记录)
- perf sched latency: output per task latency metrics(每个任务的延迟度量)
- perf sched map: show summary/map of context-switching(上下文切换的摘要和映射)
- perf sched trace: output finegrained trace(细颗粒度的跟踪)
- perf sched replay: replay a captured workload using simlated threads(使用模拟线程重播工作负载)

Example:

```bash
# 记录10s的整个系统的活动
perf sched record sleep 10
# 显示任务延迟，使用从大到小的排序。
perf sched latency --sort max
```



# Reference

[linux 性能分析工具——perf](https://blog.csdn.net/u014608280/article/details/80265718)

