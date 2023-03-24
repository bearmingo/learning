# Minik8s

## 启动

```bash
minikube start --hyperv-virtual-switch="minikube"
# 或者
minikube start --image-mirror-country='cn' \
	--registry-mirror=https://registry.docker-cn.com  \
	--memory=4096 \
	--hyperv-virtual-switch="minikube" 
	--image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
	--driver=none \
	--dns-domain=xxxx.com
# 或者
minikube start --image-mirror-country='cn' \
	--registry-mirror=https://registry.docker-cn.com  \
	--memory=4096 \
	--image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers \
	--dns-domain=xxxx.com

```

参数说明

- --`vm-driver` 如果不写会自动检测，可选值 virtualbox, vmwarefusion, hyperv, vmware
- `--image-mirror-country` 需要使用的镜像镜像的国家/地区代码。留空以使用全球代码。对于中国大陆用户，请将其设置为 cn。
- `--image-repository` 用来拉取 Kubernetes 集群所需镜像的仓库
- `--registry-mirror` docker registry 的镜像源，集群安装后拉取镜像加速用，可以使用其它加速器地址
- `--memory` 虚机内存大小
- `--kubernetes-version=v1.19.0`: 指定使用的k8s版本
- `--dns-domain`: 集群内部使用的域名

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

添加ingress

```
minikube addons enable ingress
```



## 错误处理

### error: You must be logged in to the server (Unauthorized)

```bash
```

## pod长时间启动后，服务容器被重启
由于minikube的pv放在服务器的tmp目录下，部分文件容易被宿主机删除（/tmp目录的特性）。

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: k8s.io/minikube-hostpath
parameters:
  type: pd-ssd
```

### Exiting due to DRV_AS_ROOT: The "docker" driver should not be used with root privileges.

错误信息：

```bash
[root@VM-24-17-centos ~]# minikube start
😄  Centos 8.4.2105 (amd64) 上的 minikube v1.24.0
✨  自动选择 docker 驱动。其他选项：none, ssh
🛑  The "docker" driver should not be used with root privileges.
💡  If you are running minikube within a VM, consider using --driver=none:
📘    https://minikube.sigs.k8s.io/docs/reference/drivers/none/

❌  Exiting due to DRV_AS_ROOT: The "docker" driver should not be used with root privileges.
```

**解决方法1**
添加 `--force` 参数，强制启动
```bash
minikube start --force
```

**解决方法2**
1. 添加新用户
```bash
useradd docker
usermod -aG sudo docker
su - docker
```

2. 登录新创建的用户
```bash
su - docker
```

3. 将用户添加到docker组中
```bash
sudo groupadd docker
sudo usermod -aG docker $USER
```

4. 使用以下命令重新登录并启动 minikube
```bash
minikube start --driver=docker
```

## Reference

https://minikube.sigs.k8s.io/docs/start/
[从零开始的K8S学习笔记（二）K8S本地开发环境——minikube安装部署及实践](https://zhuanlan.zhihu.com/p/574759499)