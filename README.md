# ssnt2geoip
ss -nt output with ip geo info（ss -nt命令输出时带上IP地理位置）

---

geo data from ip.sb

地理位置数据来源于ip.sb

#### demo

```bash
pi@raspberrypi: $ ./go_build_ssnt2geoip_go_armv6
2020/05/20 09:00:01 State      Recv-Q Send-Q Local Address:Port               Peer Address:Port
ESTAB      0      0      192.168.1.254:4001 局域网               80.82.17.10:59269 波兰 S-NET Sp. z o.o.
ESTAB      0      0      192.168.1.254:4001 局域网               167.99.132.247:4001 德国黑森州法兰克福 Digital Ocean
ESTAB      0      0      192.168.1.254:4001 局域网               217.24.230.181:4001 德国萨尔州 VSE NET GmbH
ESTAB      0      0      192.168.1.254:4001 局域网               110.189.192.244:1169 中国四川德阳 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               104.196.0.102:31957 美国南卡罗来纳州蒙克斯科纳 Google Cloud
ESTAB      0      0      192.168.1.254:4001 局域网               111.231.108.95:4001 中国上海 Tencent cloud computing
ESTAB      0      0      192.168.1.254:4001 局域网               117.80.194.140:41363 中国江苏苏州 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               118.114.82.203:23854 中国四川成都 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               45.55.43.166:4001 美国纽约州纽约 Digital Ocean
ESTAB      0      0      192.168.1.254:4001 局域网               138.201.67.220:4001 德国萨克森自由州法尔肯施泰因 Hetzner Online GmbH
ESTAB      0      0      192.168.1.254:4001 局域网               63.142.249.110:4001 美国德克萨斯州达拉斯 QuadraNet
ESTAB      0      0      192.168.1.254:4001 局域网               54.39.248.243:4001 加拿大魁北克省博阿努瓦 OVH SAS
ESTAB      0      0      192.168.1.254:4001 局域网              138.201.68.74:4002 德国萨克森自由州法尔肯施泰因 Hetzner Online GmbH
CLOSE-WAIT 32     0      192.168.1.254:4001 局域网              128.95.160.156:443 美国华盛顿州西雅图 University of Washington
ESTAB      0      0      192.168.1.254:4001 局域网               91.121.231.234:4001 法国上法兰西大区鲁贝 OVH SAS
ESTAB      0      0      192.168.1.254:4001 局域网               180.115.100.5:1570 中国江苏常州 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               180.115.103.160:1060 中国江苏常州 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               139.178.88.35:4001 美国加利福尼亚州森尼韦尔 Packet Host
ESTAB      0      0      192.168.1.254:4001 局域网              129.204.103.68:7000 中国广东广州 Tencent cloud computing
ESTAB      0      0      192.168.1.254:4001 局域网               184.161.147.218:57274 加拿大魁北克省蒙特利尔 Videotron Ltee
ESTAB      0      0      192.168.1.254:4001 局域网               118.113.194.255:25769 中国四川成都 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               182.138.123.26:58446 中国四川成都 China Telecom
ESTAB      0      0      192.168.1.254:4001 局域网               94.176.233.122:443 立陶宛维尔纽斯县维尔纽斯 UAB Interneto vizija
ESTAB      0      0      192.168.1.254:4001 局域网               18.18.248.83:4001 美国马萨诸塞州剑桥 Massachusetts Institute of Technology
ESTAB      0      0      192.168.1.254:4001 局域网               147.75.70.221:4001 美国加利福尼亚州森尼韦尔 Packet Host
ESTAB      0      0      192.168.1.254:4001 局域网               165.227.32.248:4001 加拿大安大略省多伦多 Digital Ocean
```
