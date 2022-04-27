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