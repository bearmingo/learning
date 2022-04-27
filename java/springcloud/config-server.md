

```
/{application}/{profile}[/{label}]
[/{label}]/{application}-{profile}{.yml|.properties|.json}
```
- `application`: 应用服务器的名称。
- `profile`: 配置环境，例如： dev/test/prod
- `label`: 仓库分钟名称（git或者svn方式指定，native本地方式无需指定