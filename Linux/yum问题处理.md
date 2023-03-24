#linux #yum 

## Module yaml error: Unexpected key in data: static_context \[line 9 col 3\]

```bash
yum clean all
rpm --rebuilddb
yum update
yum list
yum makecache
```

## 删除本地下载的包

```bash
yum clean packages
```