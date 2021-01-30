# cocoapods使用



## 安装

```bash
brew install cocoapods
brew link cocoapods
```

## 替换源

```bash
cd ~/.cocoapods/repos 
pod repo remove master
git clone https://mirrors.tuna.tsinghua.edu.cn/git/CocoaPods/Specs.git master
```

在工程文件在自己工程的`Podfile`第一行加上

```conf
source 'https://mirrors.tuna.tsinghua.edu.cn/git/CocoaPods/Specs.git'
```

参考：https://mirrors.tuna.tsinghua.edu.cn/help/CocoaPods/

恢复默认源

```bash
git remote set-url origin https://github.com/CocoaPods/Specs.git
```





## 基本命令

初始化工程

```bash
pod install
```

查看有更新的依赖库

```
pod outdated
```

更新依赖库

```
pod update
```



