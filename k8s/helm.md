# helm

安装

```bash
wget https://get.helm.sh/helm-v3.7.1-linux-amd64.tar.gz
tar -zxvf helm-v3.2.1-linux-amd64.tar.gz
mv linux-amd64/helm /usr/bin
```

添加仓库

```bash
helm repo add stable http://mirror.azure.cn/kubernetes/charts
helm repo add aliyun https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
helm repo update
```

查看仓库

```bash
helm repo list
```

删除仓库

```bash
helm repo remove aliyum
```



```bash
helm search repo weave
helm install ui stable/weave-scope
helm list
helm status ui
```

```
helm create mychart
cd mychart

helm install web1 mychart/
helm upgrade web1 mychart/
```

## Reference

https://helm.sh/docs/intro/install/