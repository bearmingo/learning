# WireGuard

## CentOS 7

### 安装

```bash
yum update -y
yum install epel-release https://www.elrepo.org/elrepo-release-7.el7.elrepo.noarch.rpm
yum install yum-plugin-elrepo
yum install kmod-wireguard wireguard-tools
reboot
```

### 生成密钥对

```bash
cd /etc/wireguard
# 服务使用
wg genkey | tee privatekey-server | wg pubkey > publickey-server
# 客户端使用，最好在客户端自己生成，这里测试用，就在服务端一起生成了
wg genkey | tee privatekey1-client | wg pubkey > publickey1-client
```

### 配置

```bash
vim /etc/wireguard/wg0.conf
```

配置文件内容

```conf
[Interface]
Address = 10.10.0.1/24
SaveConfig = true
ListenPort = 39989
PrivateKey = <server private key>
PostUp   = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE

[Peer]
PublicKey = <client public key>
AllowedIPs = 10.10.0.2/32
```

## 网卡配置

```bash
vi /etc/sysconfig/network-scripts/ifcfg-wg0
```

配置文件内容

```conf
DEVICE=wg0
TYPE=wireguard
IPADDR=10.10.0.1
NETMASK=255.255.255.0
ONBOOT=yes
NAME=wg0
ZONE=public
```

配置内核转发

```cmd
# firewall-cmd --add-masquerade --zone=public --permanent
# firewall-cmd --reload

sed -i '/net.ipv4.ip_forward/d' /etc/sysctl.conf
echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf
sysctl -p
```

### 启动服务

```bash
wg-quick up wg0
wg-quick down wg0
wg show
```

***如果需要修改wg0.conf文件，先关闭wg服务***

## CentOS 8

```bash
dnf install epel-release elrepo-release
dnf update
dnf install kmod-wireguard wireguard-tools
```

## macOS 客户端

### 安装

```bash
brew info wireguard-tools
brew install wireguard-tools
```

### 配置

添加配置文件

```bash
mkdir -p /usr/local/etc/wireguard
vim /usr/local/etc/wireguard
```

wg0.conf配置内容

```conf
[Interface]
Address = 10.0.0.2
ListenPort = 39919
PrivateKey = <private key>
# DNS = 8.8.8.8

[Peer]
AllowedIPs = 0.0.0.0/0
Endpoint = <service ip and port>
PublicKey = <public key>
PersistentKeepalive = 25
```



## 问题解决

### RTNETLINK answers: Operation not supported

安装完成WireGuard后需要重启系统

```bash
reboot
```

### ping不通对端的IP

关闭firewall

```bash
systemctl stop firewalld.service
```



## Reference

https://www.linode.com/docs/guides/centos-wireguard-installation-and-configuration/

[centos7.x下配置wireguard](https://blog.csdn.net/ytfsksk/article/details/113094197)

[用 wg-quick 调用 wg0.conf 管理 WireGuard.md](https://github.com/wgredlong/WireGuard/blob/master/2.%E7%94%A8%20wg-quick%20%E8%B0%83%E7%94%A8%20wg0.conf%20%E7%AE%A1%E7%90%86%20WireGuard.md)