# GIT

配置pull的默认行为
```bash
git config pull.ff only
```

设置代理
```bash

git config --global http.proxy http://proxyuser:proxypwd@proxy.server.com:8080
git config --global http.proxy socks5://localhost:7088

git config --global --get http.proxy
git config --global --unset http.proxy
```




