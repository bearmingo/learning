
## Kubernetes 网络需求

在深入了解数据包如何在kubernetes集群内流动的细节之前，让我们先了解下Kubernetes网络的需求。

Kubernetes网络模型定义来一组基本的规则：
1. 集群中的POD能够在不使用网络地址转换(NAT)的情况下与任何其他POD自由通信。
2. 在集群节点上运行的任何程序应该在不使用NAT的情况下与同一节点上的任何POD通信。
3. 每一个POD都有自己的IP地址(IP-per-POD)，其他POD都可以用同一个地址访问它。

这些需求并没有将实现限制在单个解决方案中。

相反的，他们用一般的术语来描述集群网络的特性。

为了满足这些限制，您必须解决一下挑战：
1. 你如何确保在同一个POD中的容器表现的像他们在同一台主机上一样？
2. POD能够访问集群中的其他POD吗？
3. POD能够访问Service么？Service负载均衡请求吗？
4. POD是否能够接收集群外部的请求？

在本文中，您将重点关注前三点，从POD的内网或者容器到容器的通信开始。

## Linux的网络名称空间如何在POD中工作

让我们考虑一个托应用程序的主容器和另一个与它一起运行的容器。

在此示例中，您有一个带Nginx容器的POD和另一个带busybox的POD。
```yaml
apiVersion: v1
  kind: Pod
  metadata:
	name: multi-container-pod
	spec:
	  containers:
	  - name: container-1
	    image: busybox
	    command: ['/bin/sh', '-c', 'sleep 1d'] 
	  - name: container-2 
	    image: nginx
```

部署后，会发生一下情况：
1. POD在节点上获得**自己的网络命名空间**。
2. **一个IP地址被分配**给POD，端口在两个容器之间共享。
3. **两个容器共享同一个网络命名空间**，并且可以在本地主机上看看到对方。

网络配置在后台快速进行。

但是让我们后退一步，尝试理解为什么容器运行需要上述内容。

在Linux中，网络命名空间是独立的、隔离的逻辑空间。

您可以将网络名称空间视为获取无力网络接口并将其分割成更小的独立部分。

每个部分都可以单独配置，并具有自己的网络规则和资源，包括防火墙配置、接口（虚拟或者物理）、路由以及与网络相关的其他内容。

> 物理网络接口拥有根网络命名空间
> ![](https://learnk8s.io/a/941e6d55d9aac9cc43e964ad102e9391.svg)
> 您可以使用 Linux 网络命名空间来创建隔离网络。每个网络都是独立的，除非您将其配置为非独立，否则不会与其他网络通信。
> ![](https://learnk8s.io/a/45d5de36a76aa86a898eb1ffb94e55b3.svg)

物理接口最终必须处理所有真实数据包，因此所有虚拟接口都是从中创建的。

网络命名空间可以通过`ip-netns`管理工具进行管理，您可以使用`ip netns list` 来列出主机上的命名空间。

> [!NOTE]
> 请注意，在创建网络空间时，他将出现在`/var/run/netns`下面。但Docker并不总是遵循这一点。

例如，这些来自 Kubernetes 节点的命名空间。

```bash
$ ip netns list
cni-0f226515-e28b-df13-9f16-dd79456825ac (id: 3)
cni-4e4dfaac-89a6-2034-6098-dd8b2ee51dcd (id: 4)
cni-7e94f0cc-9ee8-6a46-178a-55c73ce58f2e (id: 2)
cni-7619c818-5b66-5d45-91c1-1c516f559291 (id: 1)
cni-3004ec2c-9ac2-2928-b556-82c7fb37a4d8 (id: 0
```

> [!NOTE]
> 注意： `cni-` 前缀；这意味着命名空间的创建是由 `CNI` 负责的。

当您创建一个POD，并且该POD被奉陪给一个节点时，`CNI` 将：
1. 分配一个IP地址。
2. 将容器附加到网络。

如果POD包含多个容器（如上面的例子），则两个容器都放到同一个命名空间中。

> 1. 当POD创建时，容器运行时首先为容器创建网络命名空间。
> ![[create-pod-step-1.svg]]
> 2. 然后， CNI 牵头为容器分配一个IP地址。
> 
> ![[create-pod-step-2-4a24a4a646939d95aff02cfd1d7b7ae2.svg]]
>  3. 最后，CNI 将容器附加到网络的中。
>  ![[create-pod-step-3-963c7aa880a953445bed849882e720ae.svg]]

那么当您列出节点上的容器时会发生什么？

您可以通过SSH连接到Kubernetes节点并探索命名空间：
```bash
$ lsns -t net
NS TYPE NPROCS PID USER NETNSID NSFS
4026531992 net 171 1 root unassigned /run/docker/netns/default /sbin/init noembed norestore
4026532286 net 2 4808 65535 0 /run/docker/netns/56c020051c3b /pause
4026532414 net 5 5489 65535 1 /run/docker/netns/7db647b9b187 /pause
```

`lsns` 用于列出主机上所有可用命名空间在哪里。

> [!INFO]
> 请记住，Linux中有多种命名空间类型

*Nginx 容器在哪里？*

*这些`pause`容器是什么？*

## pause容器在POD中创建网络命名空间

让我们列出节点上所有的进程，并检查我们是否可以找到 Nginx 容器：

```bash
$ lsns
...
NS TYPE NPROCS PID USER COMMAND
4026532414 net 5 5489 65535 /pause
4026532513 mnt 1 5599 root sleep 1d
4026532514 uts 1 5599 root sleep 1d
4026532515 pid 1 5599 root sleep 1d
4026532516 mnt 3 5777 root nginx: master process nginx -g daemon off;
4026532517 uts 3 5777 root nginx: master process nginx -g daemon off;
4026532518 pid 3 5777 root nginx: master process nginx -g daemon off;
```

该容器列在 挂载（`mnt`）、Unix分时（`uts`）、PID（`pid`）命名空间中，但不在网络命名空间中（`net`）中。

不幸的是，`lsns`只显示每个进程的最低PID，当您可以根据进程ID进一步过滤。

您可以使用以下方法检索Nginx容器的所有命名空间。

```sh
$ sudo lsns -p 5777
NS TYPE NPROCS PID USER COMMAND 
4026531835 cgroup 178 1 root /sbin/init noembed norestore 
4026531837 user 178 1 root /sbin/init noembed norestore 
4026532411 ipc 5 5489 65535 /pause 
4026532414 net 5 5489 65535 /pause 
4026532516 mnt 3 5777 root nginx: master process nginx -g daemon off; 
4026532517 uts 3 5777 root nginx: master process nginx -g daemon off; 
4026532518 pid 3 5777 root nginx: master process nginx -g daemon off;
```

又是这个`pause`过程，这次它劫持了网络命名空间。

*那是什么？*

**集群中的每一个POD都有一个额外的隐藏容器在后台运行，名为：“pause”**

如果列出节点上运行的容器，并过滤`pause`容器：

```bash
$ docker ps | grep pause 
fa9666c1d9c6 registry.k8s.io/pause:3.4.1 "/pause" k8s_POD_kube-dns-599...
44218e010aeb registry.k8s.io/pause:3.4.1 "/pause" k8s_POD_blackbox-exporter… 5fb4b5942c66 registry.k8s.io/pause:3.4.1 "/pause" k8s_POD_kube-dns-… 
8007db79dcf2 registry.k8s.io/pause:3.4.1 "/pause" k8s_POD_konnectivity-agent-84f87c…
```

您将看到，对于节点上每个分配的POD，都会有一个pause容器将会自动与其配对。

**该`pause`容器负责创建和保存网络命名空间。**

*创建命名空间？*
*是不是呢？*

__网络命名空间的创建有底层容器运行时完成__。通常由`containerd` 或者 `CRI-0`负责。

在部署POD和创建容器之前，它运行责任是创建网络命名空间。

容器运行时会自动执行此操作，而不是手动运行`ip netns`和创建网络命名空间。

回到`pause`容器。

它包含非常少的代码，并在部署后立即进入休眠状态。

然而，==它是必不可少的，并且在Kubernetes生态系统中起着至关重要的作用。==

> 1. 创建 POD 时，容器运行时会创建一个带有 pause 容器的网络命名空间。
> ![[pause-1-201510ad30f9ae89515be0d79f35e597.svg]]
> 2. POD 中的每个其他容器都加入了该容器创建的现有网络命名空间。
> ![[pause-2-822de32c50f94cc5f439d84242a59a83.svg]]
> 3. 此时，CNI 分配 IP 地址并将容器附加到网络。
> ![[pause-3-6400a5aa88f9ef3bdc165a1e2cfeb190.svg]]

_一个进入休眠状态的容器怎么会有作用呢？_

为了理解它的实用性，让我们想象一下“有一个像前面的例子一样的两个容器的POD，但没有`pause` 容器”。

一旦容器启动，CNI：
1. 使`busybox`容器加入到前面的网络命名空间。
2. 分配IP。
3. 将容器附加到网络。

_如果Nginx崩溃了怎么办？_

CNI将不得不再次执行所有的步骤，并且两个容器的网络都将中断。

由于执行sleep的容器不太可能有任何错误，因此由它创建网络命名空间通常是更为安全、更加可靠的选择。

**如果POD中任何一个容器崩溃，其余容器仍然可以响应任何网络请求。**

## POD 被分配了一个 IP 地址

前面提到过POD和两个容器接收相同的IP。

_那是怎么配置的呢？_

**在POD网络命名空间内，创建了一个接口(eth0)，并分配了一个IP地址。**

让我们来验证一下。

首先，找到POD的IP地址：

```bash
$ kubectl get pod multi-container-pod -o jsonpath={.status.podIP}
10.244.4.40
```

接下来，让我们来找到相关的网络命名空间。

由于网络命名空间是从物理接口创建的，因此您必须访问集群节点。

> [!NOTE]
> 如果您运行在`minikube`，您可以尝试`minikube ssh` 访问该节点。如果您运行在云提供商中运行，应该可以使用SSH连接到节点上。

进入后，让我们找到创建的最新网络命名空间：

```bash
ls -lt /var/run/netns
total 0
-r--r--r-- 1 root root 0 Sep 25 13:34 cni-0f226515-e28b-df13-9f16-dd79456825ac
-r--r--r-- 1 root root 0 Sep 24 09:39 cni-4e4dfaac-89a6-2034-6098-dd8b2ee51dcd
-r--r--r-- 1 root root 0 Sep 24 09:39 cni-7e94f0cc-9ee8-6a46-178a-55c73ce58f2e
-r--r--r-- 1 root root 0 Sep 24 09:39 cni-7619c818-5b66-5d45-91c1-1c516f559291
-r--r--r-- 1 root root 0 Sep 24 09:39 cni-3004ec2c-9ac2-2928-b556-82c7fb37a4d8
```

在这种情况下，应该是 `cni-0f226515-e28b-df13-9f16-dd79456825ac` 。

现在您可盈在该命名空间内运行 `exec` 命令。

```bash
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac ip a 
# output truncated 
3: eth0@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default 
    link/ether 16:a4:f8:4f:56:77 brd ff:ff:ff:ff:ff:ff link-netnsid 0 
    inet 10.244.4.40/32 brd 10.244.4.40 scope global eth0 
	   valid_lft forever preferred_lft forever 
	inet6 fe80::14a4:f8ff:fe4f:5677/64 scope link 
	    valid_lft forever preferred_lft forever
```

**那是POD的IP地址**

`12` 让我们通过 `grep` 的部分找出该接口的另一端`@if12`。

```bash
$ ip link | grep -A1 ^12 
12: vethweplb3f36a0@if16: mtu 1376 qdisc noqueue master weave state UP mode DEFAULT group default
    link/ether 72:1c:73:d9:d9:f6 brd ff:ff:ff:ff:ff:ff link-netnsid 1
```

您还可以验证 `Nginx` 容器是否侦听该命名空间的HTTP流量。

```bash
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac netstat -lnp
Active Internet connections (only servers) 
Proto Recv-Q Send-Q Local Address Foreign Address State   PID/Program name 
tcp   0      0      0.0.0.0:80    0.0.0.0:*       LISTEN  692698/nginx: master 
tcp6  0      0      :::80         :::*            LISTEN  692698/nginx: master
```

> [!INFO]
> 如果您无法通过SSH访问集群中的工作节点，您可以使用`kubectl exec shell` 获取 busybox 容器并直接在其中使用 `ip` 命令（或者 `netstat` ）。

优秀～～～

现在我们介绍容器直接的通信，让我懵看看POD到POD的通信是如何建立的。

## 检查集群中 pod 到 pod 的流量

当POD到POD通信出现问题时，有两种可能的情况。
1. POD流量目标是同一个节点上的POD。
2. POD流量的目标是驻留在不同节点上的POD。

为了使整个设置正常，我们需要**虚拟接口对**（我们上面已经讨论过的）和**以太网桥**。

在继续之前，让我们讨论一下他们的功能以及为什么需要它们。

**一个POD要与其他POD通信，它必须首先有权限访问节点根命名空间。**

这是使用连接两个命名空间的虚拟以太网对实现的： `pod` 和 `root`。

这些虚拟接口设备（因此 `veth` 单词中带一个`v`）连接并通道两个命名空间的隧道。

使用此 `veth` 设备，您可以将一端连接到POD的命名空间，将另一端连接到根命名空间（root)。
![[veth-pair-6400a5aa88f9ef3bdc165a1e2cfeb190.svg]]
CNI将会为你完成此操作，但您也可以手动执行此操作。

```bash
$ ip link add veth1 netns pod-namespace type veth peer veth2 netns root
```

现在，您的POD命名空间具有到根命名空间的访问“隧道“。

**节点上每个新创建的POD都将设置一对veth这样的隧道。**

创建接口对是其中的一部分。

另一个是为以太网设备分配地址，并创建默认路由。

`veth1` 让我们探讨如何在POD的命名空间中设置接口：

```bash
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac ip addr add 10.244.4.40/24 dev veth1 
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac ip link set veth1 up 
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac ip route add default via 10.244.4.40
```

在节点终端上，让我们创建另一个`veth2`对。

```bash
$ ip addr add 169.254.132.141/16 dev veth2 
$ ip link set veth2 up
```

您可以像以前一样检查现有的`veth`对。

在POD命名空间中，检索接口后缀`eth0`.

```bash
$ ip netns exec cni-0f226515-e28b-df13-9f16-dd79456825ac ip link show type veth 
3: eth0@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP mode DEFAULT group default 
    link/ether 16:a4:f8:4f:56:77 brd ff:ff:ff:ff:ff:ff link-netnsid 0
```

在这种情况下，您可以`grep` （例如`grep -A1 ^12`)。

```shell
$ ip link show type veth 
# output truncated 
12: cali97e50e215bd@if3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP mode DEFAULT group default 
    link/ether ee:ee:ee:ee:ee:ee brd ff:ff:ff:ff:ff:ff link-netns cni-0f226515-e28b-df13-9f16-dd79456825ac
```

> [!NOTE]
> 您也可以使用`ip -n cni-0f226515-e28b-df13-9f16-dd79456825ac link show type veth`。
> 

注意两个`3: eth0@if12` 和 `12: cali97e50e215bd@if3` 接口上的符号。

从POD命名空间，该 `eth0` 接口连接到根命名空间中的接口号`12`。因此 `@if12`。

在对的另一端`veth`，根命名空间连接到pod命名空间接口编号`3`。

接下来是连接两端`veth`的桥。

## POD网络命名空间连接到以太网桥

该桥会将位于根命名空间中的虚拟接口的每一端“绑定”在一起。

**该网桥将允许流量在虚拟对之间流动，并通过公共根命名空间。**

*Theory time.*

以太网桥位于OSI网络模型的第二层。

您可以将网桥视为接受来自不同名称空间和接口的连接的虚拟交换机。

因此，您可以使用此设置来桥接两个接口，从POD命名空间的veth 到在**同一节点上**的另一个POD的veth。
![[pod-to-pod-by-bridge-72130856cd225fe326bccde12b524814.svg]]
让我们看看实际中的以太网桥和veth对。



## 参考

翻译：https://learnk8s.io/kubernetes-network-packets
https://learnk8s.io/kubernetes-network-packets
https://fedoramagazine.org/network-address-translation-part-2-the-conntrack-tool/

https://kubernetes.io/blog/2019/03/29/kube-proxy-subtleties-debugging-an-intermittent-connection-reset/

[How Container Networking Works: Practical Explanation](https://iximiuz.com/en/posts/container-networking-is-simple/)