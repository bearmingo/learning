# Bash脚本代码片段


## switch 使用

```sh
case "$i" in
    *.sh)
        echo this is a bash script
    ;;
    *)
        echo default
        ;;
esac
```

## 文件夹是否存在

条件语法if中使用
```sh
if [ -d $path ];
then
    echo "dir exist!"
else
    echo "dir does not exist!"
fi
```

单行使用
```sh
[ -d $path ] && echo exist
```


## 脚本所在的路径

```sh
basepath=$(cd `dirname $0`; pwd)
```
- `dirname $0`: 取得当前脚本文件的父路径
- cd `dirname $0`: 进入到这个目录
- pwd: 显示当前的工作目录