# Minik8s

## 启动

```bash
minikube start --hyperv-virtual-switch="minikube"
# 或者
minikube start  image-mirror-country='cn' --registry-mirror=https://registry.docker-cn.com  --memory=4096 --hyperv-virtual-switch="minikube" --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers
# 或者
minikube start  image-mirror-country='cn' --registry-mirror=https://registry.docker-cn.com  --memory=4096 --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers

```

参数说明

- --`vm-driver` 如果不写会自动检测，可选值 virtualbox, vmwarefusion, hyperv, vmware
- `--image-mirror-country` 需要使用的镜像镜像的国家/地区代码。留空以使用全球代码。对于中国大陆用户，请将其设置为 cn。
- `--image-repository` 用来拉取 Kubernetes 集群所需镜像的仓库
- `--registry-mirror` docker registry 的镜像源，集群安装后拉取镜像加速用，可以使用其它加速器地址
- `--memory` 虚机内存大小
- --kubernetes-version=v1.19.0: 指定使用的k8s版本

## 查看器群状态

```bash
minikube kubectl get nodes
```

## 运行dashboard

```bash
minikube.exe dashboard
```

## 登入虚拟机

```bash
minikube.exe ssh
```

## 删除容器

```bash
minikube delete --all
# 完全重建
minikube delete --all --purge
```

## 错误处理

### error: You must be logged in to the server (Unauthorized)

```bash
```

## Reference

https://minikube.sigs.k8s.io/docs/start/