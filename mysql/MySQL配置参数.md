## 参数
### `max_connections`
允许最大连接的客户端数量。增加该值时，同时需要增加mysqld的文件描述符数量。否则，会报`too many connections`。默认值为100个。
查看对应信息的方式：
```sql
SHOW VARIABLES LIKE '%max_connections%';
```

### `back_log`
当MySQL(处理大量请求时，由于满负荷)暂时停止回复新请求前，能过对方的请求数。如果用户希望段时间内有更大的连接数，应该增大`back_log`的值。默认值为50。
在MySQL主线程短时间内收到大量连接请求时会涉及`back_log`，然后主线程在一段时间（或者瞬间）后检查连接数，并启动新线程。
换而言之，`back_log`的值表示监听的TCP/IP连接队列的大小。当然操作系统也有自己的队列大小限制。如果需要，可以检查操作系统文档中关于该最大值的说明。`back_log`不能大于操作系统限制的最大值。

### `wait_timeout`
服务在关闭之前在一个连接上等待的秒数。默认值为28800。

### `interactive_timeout`
服务器在关闭它前，在一个交互连接上等待的秒数。默认值为28800
如果客户端连接时，`mysql_read_connect`使用量`client_interactive`选项，这个连接将被标记为交互客户短。


## 缓存相关的配置参数
### `query_cache_size`
查询缓冲大小，当打开时，执行查询语句会进行缓存，读写都会带来额外的内存开销。下次再次查询若命中该缓存，会立刻返回结果。该选项默认关闭，打开则需要调整参数项`query_cache_type=ON`。默认值为64M。

### `thread_cache_size`
可以复用的保存在缓存中的线程数量。如果有需要，新的线程从缓存中获取。
通过比较`connections`和`threads_created`的状态，可以卡看到这个变量的作用。

### `sort_buffer`
每个需要进行排序的线程分配该大小的一个缓冲区。增加这个值能过加速`order by`或者`group by`操作。默认值为2M。
这是个一次性分配设置的内存，并不是越大越好，由于时connections级别的参数，过大的设置+高并发可能会耗尽系统内存资源。

### `table_cache`
所有线程打开表的总数量。增加该值时需要同时增加mysqld的最大文件描述符数量。MySQL对每个打开的表需要2个文件描述符。默认值为64。

### `record_buffer`
进行顺序扫描的线程为其扫描的每张表分配这个大小的一个缓冲。如果有很多顺序扫描，可以增加该值。默认值为128k。

### `key_buffer_size`
索引块是缓存的并且被所有线程共享。`key_bufer_size`是用于索引块的缓存区大小，增加它可更好处理的索引（对所有读和多重写）。如果太大，系统将开始换页并且变慢。默认值时8m。
只对MyISAM表起作用。即使不用MyISAM表，但内部临时的磁盘表是MyISAM表，也要使用该值。

### `tmp_table_size`
临时表的大小，做类似的`GROUP BY`操作时生成临时表。提高连表查询的效果，调整该值直到`created_tmp_disk_tables / created_tmp_tables * 100% < 25%`，处于这样的状态之下，效果较好。范围设置为64M~256M最佳。不宜过大，会导致内存不足、I/O阻塞。

### `read_buffer_size`
MySQL读入按**数据文件存储顺序数据**的缓冲区大小，，默认为2M。
将对表进行顺序扫描的请求将分配一个读入缓冲区，MySQL会为它分配一段内存缓冲区，`read_buffer_size`变量控制这一缓冲区的大小。如果对表的顺序扫描非常频繁，并你认为频繁扫描进行的太慢，可以通过增加内存缓冲区大小提高性能。

### `read_rnd_buffer_size`
随机读缓存大小。
当按任意顺序读取行时（例如按照排序顺序）将分配一个随机读取缓冲区。进行排序查询时，MySQL会先扫描一遍该缓存，以避免磁盘搜索，提高查询数据。如果需要大量数据可适当的调整该值。但MySQL会为每个客户端连接分配该缓冲区，所以尽量适当设置该值，以避免内存开销过大。

### `join_cache_size`
多表参与`join`操作时分配的缓存，

### `thread_stack`
每个连接线程被创建时，MySQL给他分配的内存大小。当MySQL创建一个新的连接线程时，需要给它分配一定大小的内存堆栈空间，一边存放客户端的请求Query及自身的各种状态和处理信息。Thead_Cache_Hit = (Connections - Threads_created) / Connections * 100%;命中率处于90%才算正常配置，当出现“mysql-debug: Thread stack overrun“的错误提示的时候需要增加该值。

### `binlog_cache_size`
为每个session分配的内存，在事务过程中用来存储二进制日志的缓存。作用时提高记录binlog的效率。没有什么大事务，dml也不是很频繁的情况下可以设置小一点，如果事务大而且多，dml操作也频繁，则可以适当的调大一点。

### `innodb_buffer_pool_size`
主要作用是缓存innodb的索引、数据、插入数据时的缓冲。
专用MySQL服务器设置的大小：操作系统内存 70% ～80%。
设置过大会使得系统的swap空间被占用，导致操作系统变慢，从而减低sql查询效率。

### `innodb_additional_mem_pool_size`
用来存放Innodb内部的目录，这个值不用分配的太大，系统可以自动调。一般默认16M够用。如果表比较多，可以适当的增大。如果这个值自动增加时，会在error log日志中有显示。
### `innodb_log_buffer_size`
日志缓存大小，默认值为：8M。
InnoDB的写操作，将数据写入到内存中的日志缓存中，由于InnoDB再事务提交前，并不将改变的日志写入到磁盘中，因此可在大事务中，可以减轻磁盘I/O的压力。通常情况下，如果不是写入大量的超大二进制数据（a lot of huge blobs），4MB-8MB已经足够了。此处我们设置为8M。

## 参数查看
查看以上参数的方式：
```sql
UPDATE performance_shema.setup_instruments SET enabled = 'yes' WHERE name LIKE 'memory%'
```

### 查看是否有长期运行或者阻塞的SQL
```sql
show full processlist
```


```conf
max_connections = 151
# 同时处理最大连接数，建议设置最大连接数是上限连接数的80%左右,一般默认值为151，可以做适当调整。
sort_buffer_size = 2M
# 查询排序时缓冲区大小，只对order by和group by起作用，建议增大为16M
open_files_limit = 1024 
# 打开文件数限制，如果show global status like 'open_files'查看的值等于或者大于open_files_limit值时，程序会无法连接数据库或卡死
```

innerDB配置参数
```conf
innodb_buffer_pool_size = 128M
# 索引和数据缓冲区大小，建议设置物理内存的70%左右（这个前提是这个服务器只用做Mysql数据库服务器）
innodb_buffer_pool_instances = 1    
# 缓冲池实例个数，推荐设置4个或8个
innodb_flush_log_at_trx_commit = 1  
# 关键参数，0代表大约每秒写入到日志并同步到磁盘，数据库故障会丢失1秒左右事务数据。1为每执行一条SQL后写入到日志并同步到磁盘，I/O开销大，执行完SQL要等待日志读写，效率低。2代表只把日志写入到系统缓存区，再每秒同步到磁盘，效率很高，如果服务器故障，才会丢失事务数据。对数据安全性要求不是很高的推荐设置2，性能高，修改后效果明显。
sync_binlog=1
 
innodb_file_per_table = ON  
# 是否共享表空间，5.7+版本默认ON，共享表空间idbdata文件不断增大，影响一定的I/O性能。建议开启独立表空间模式，每个表的索引和数据都存在自己独立的表空间中，可以实现单表在不同数据库中移动。
innodb_log_buffer_size = 8M  
# 日志缓冲区大小，由于日志最长每秒钟刷新一次，所以一般不用超过16M 
```