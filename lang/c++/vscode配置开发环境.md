# VSCode配置开发环境
## 安工具和插件
- cmake


## 配置 launch.json

```json
{
	"version": "0.2.0",
        {
            "name": "启动",
            "type": "cppdbg",
            "request": "launch",
            "program": "${command:cmake.launchTargetPath}",
            "args": [],
            "stopAtEntry": false,
            "cwd": "${command:cmake.launchTargetDirectory}",
            "environment": [],
            "externalConsole": false,
			"MIMode": "lldb",
            "logging": {
                "moduleLoad": false,
                "trace": true
            },
            // 调试Qt程序需要指定.natvis文件,比如QString之类的对象可以看到其内容
            // "visualizerFile":"C:/Qt/qt5.natvis"
        },
```
}
```

## 配置tasks.json
```json
```

## 参考
[Windows下VSCode+CMake搭建开发环境](https://zhuanlan.zhihu.com/p/370211322)