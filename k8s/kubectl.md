# kubectl

get

```bash
kubectl get po
```

- po
- depolyment
- sc
- pv
- sv
- all

```bash
kubectl cluster-info
```

获取容器

```bash
# -l 根据标签过滤
kubectl get pods -l app=v1
```

创建发布service

```bash
kubectl expose deployment/kubernetes-bootcamp --type=“NodePort” --port 8080 --name=my-service
```



伸缩

```bash
kubectl scale deployments/XXX --replicas=N
```



添加标签

```bash
kubectl label pod XXX app=v1
```

升级

```bash
kubectl set image deployments/kubernetes-bootcamp kubernetes-bootcamp=jocatalin/kubernetes-bootcamp:v2
```

滚动回滚

```bash
kubectl rollout undo deployments/xxxxx
```

查看资源

```bash
kubectl api-resources
```

???

```bash
kubectl kubectl drain
```

