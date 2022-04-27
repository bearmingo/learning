```bash
# TIME_WAIT超时时间，默认是60s
net.ipv4.tcp_fin_timeout = 30 
# 增加tcp支持的队列数，加大队列长度可容纳更多的等待连接
net.ipv4.tcp_max_syn_backlog = 65535
# 减少断开连接时 ，资源回收
net.ipv4.tcp_max_tw_buckets = 8000
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_tw_recycle = 1
net.ipv4.tcp_fin_timeout = 10
# 打开文件的限制
*soft nofile 65535
*hard nofile 65535 

```
