#linux #调试
## 映射分类

**文件映射**

磁盘中的数据通过文件系统，映射到物理内存中，再映射到虚拟内存中。用户就可以在用户空间中通过read、write操作文件。至于实际的代码，open,read,write,close,mmap… 操作的虚拟地址都属于文件映射。

**匿名映射**

匿名映射就是用户空间需要分配一定的物理内存来存储数据，这部分数据不属于任何文件。内核使用匿名映射将内存中的某段地址与用户空间一一映射，这样用户就可以直接操作虚拟地址来读写这段物理内存。比如使用malloc申请的内存。

## 输出的信息解释

![[smap-example.png]]

**第一行**

`7fbacd1d2000-7fbacd2d2000` 是虚拟内存地址的开始和结束位置。

`---p` 内存段的权限，分别为：可读（r）、可写（w）、可运行（x)、私有或者共享（p表示私有，s表示共享）

`00003000` 该虚拟内存段在对应映射文件中的偏移量（以页为单位）。对于匿名映射，它等于0或者`vm_start/PAGE_SIZE`

`fd:01` 文件的主设备号和子设备号。对于匿名映射，因为没有对应的文件，所以它都是`00:00`。对于有名映射来说，就是文件文件所在设备的设备号。

`265021` 被映射到虚拟内存的文件的索引节点号。通过该节点可以找到对应的文件。对于匿名映射文件来说，因为没有文件存在磁盘上，所以没有节点号，始终为`00000000`。

`/usr/lib64/libonion_security.so.1.0.19` 被映射到虚拟内存的文件名称。后面带`deleted`是内存数据，可以被销毁。对于有名映射来说，是映射的文件名称。对于匿名映射来说，是此段虚拟内存在进程中的角色。`[stack]` 表示在进程中作为栈使用，`[heap]` 表示堆，其他情况则不显示。

**剩余行**

**Size** 虚拟内存空间大小。但这个内存值不一定是物理内存实际分配的大小。因为在用户态上，虚拟内存总是延迟分配的。这个值计算也非常简单，就是VMA的结束地址减开始地址。

**Rss** 实际分配的内存。这部分物理内存已经分配，不需要缺页中断就可以使用。

Rss的计算公式：
```
Rss = Shared_Clean + Shared_Dirty + Private_Clean + Private_Dirty
```

- shared/private 表示该页面是共享还是私有。
- dirty/clean: 表示该页面是否被修改过，如果修改过（dirty)，该页面被淘汰的时候，就会把该脏页面回写到交换分区（换出，swap out)。有一个标志位用于表示页面是否dirty。

**Pss** （proportional set size)，平摊计算后的实际物理使用内存（有些内存会和其他进程共享，例如mmap进来的）。实际上包含`private_clean + private_dirty`，和安比例均分的 `shared_clean`、`shared_dirty`。

举个计算Pss的例子：  

如果进程A有x个private_clean页面，有y个private_dirty页面，有z个shared_clean仅和进程B共享，有h个shared_dirty页面和进程B、C共享。那么进程A的Pss为：  
$$x + y + z/2 + h/3$$

**Referenced** 当前页面被标记为已引用或者包含匿名映射（The amount of memory currently marked as referenced or a mapping associated with a file may contain anonymous pages)。

在Linux内存管理页面的替换算法中，当某个页面被访问后，`Referenced`标志被设置，这个页面就不能将该页面换出。

**Anonymous** 匿名映射的物理页面，这部分内存不来自文件的内存大小。

**ShMEMPmdMapped** PMD页面已经被映射的共享（shmen/tmpfs）内存量。在官方文档中，这样解释：

	“ShmemPmdMapped” show the amount of shared (shmem/tmpfs) memory backed by huge pages.

`Shared/Private_Huagetlb` 由hugetlbfs页面支持的内存使用量，由于历史原因，该页面未计入`RSS`或者`PSS`字段中。并且这些没有包含在`Shared/Private_Clean/Dirty`字段中。

**Swap** 存在与交换分区的数据大小。如果物理内存有限，可能存在一部分在主存，一部分在交换分区。

**SwapPss** 与Pss类型，只是针对的是交换分区的内存。

**KernalPageSize** 内核一页的大小。

**MMUPageSize** MMU页大小，大多数情况下，和KernalPageSize大小一样。

**Locked** 常驻与物理内存的大小，这些页不会被换出。

**THPeligible** 映射是否符合分配THP的条件。如果为`true`，则为`1`，否则为`0`。它仅显示当前状态。

**VmFlags** 表示与特定虚拟内存区域关联的内核标志。标志如下

```txt
rd - readable
wr - writable
ex - execuatable
sh - shared
mr - may read
mw - may write
me - may execute
ms - may share
gd - stack segment growns down
pf - pure PFN range
dw - disabled write to the mapped file
lo - pages are locked in memory
io - memory maped I/O area
sr - sequential read advise provided
rr - random read advise provided
dc - do not copy area on fork
de - do not expand area on remapping
ac - area is accountable
nr - swap space is not reserved for the arean
ht - area uses huge tlb pages
ar - architecture specific flag
dd - do not include area into core dump
sd - soft-dirty flag
mm - mixed map area
hg - huge page advise flag
nh - no-huage page advise flag
mg - mergable advise flag
```

## 其他

### 物理内存的延迟分配

延迟分配就是当进程申请内存的时候，Linux会给他先分配页，但是并不会区建立页与页框的映射关系，意思就是说并不会分配物理内存，而当真正使用的时候，就会产生一个缺页异常，硬件跳转page fault处理程序执行，在其中分配物理内存，然后修改页表(创建页表项)。异常处理完毕，返回程序用户态，继续执行。

### THP

`THP` 透明大页（Transparent Huage Pages)，RHEL 6 开始引入，目的是使用更大的内存页面，1⃣️适应越来越大的系统内存，让操作系统可以支持现代硬件架构的大页容量功能，与标准大页的区别在于分配机制，标准大页管理是预分配的方式，而透明大页管理则是动态分配的方式。

### TLB

`TLB`  快表(translation lookaside buffer)，直译为旁路块表缓冲。它是一种高速缓存，内存管理硬件使用它来改善虚拟地址到物理地址的转换速度。

由于页面存放在主存中，因此程序每次访问至少需要两次：一次访问物理地址，第二次才获取数据。依靠页表访问的局部性（当一个转换的虚拟页号被使用时，他可能在不久的将来再次被使用到），TLB能提高转换速度。

## 用例

```txt
7fbacd2d3000-7fbacd2d7000 rw-p 00000000 00:00 0 
Size:                 16 kB
KernelPageSize:        4 kB
MMUPageSize:           4 kB
Rss:                   8 kB
Pss:                   8 kB
Shared_Clean:          0 kB
Shared_Dirty:          0 kB
Private_Clean:         0 kB
Private_Dirty:         8 kB
Referenced:            4 kB
Anonymous:             8 kB
LazyFree:              0 kB
AnonHugePages:         0 kB
ShmemPmdMapped:        0 kB
FilePmdMapped:        0 kB
Shared_Hugetlb:        0 kB
Private_Hugetlb:       0 kB
Swap:                  0 kB
SwapPss:               0 kB
Locked:                0 kB
THPeligible:		0
VmFlags: rd wr mr mw me ac sd 
```
