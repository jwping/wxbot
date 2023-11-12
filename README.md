# wxbot - 微信聊天机器人
> 适用于微信（WeChat **3.9.8.15** | 3.9.7.29）
> 可在Windows PC微信 **设置** - **关于微信** - **版本信息** 中获取您当前的微信版本，如果您当前的微信版本不在上述可用的版本列表中，请至下方 **3、可用版本微信安装包获取** 选择最新版微信重新安装使用

**未经过大量测试，使用远程线程注入方式可能会被报毒（无毒，请放心使用！），也可以尝试使用例如x64dbg等方式进行注入，注入手段并不重要，只要将wxbot.dll注入到wechat.exe中即可**

## 免责声明
本仓库发布的内容，仅用于学习研究，请勿用于非法用途和商业用途！如因此产生任何法律纠纷，均与作者无关！

## 1、运行
bin目录下有如下两个文件（仅在windows 10 & windows server 2012 R2系统上进行测试）：
* inject.exe (bin/inject.exe)
* wxbot.dll (bin/wxbot.dll)

运行的时请保证这两个文件在同一目录下，直接运行inject.exe即可（运行注入器前请保证微信已登陆！）
默认wxbot.dll为最新版（3.9.8.15），低版本微信注入请选择对应版本的wxbot-xxxx.dll替换为wxbot.dll后注入即可
**运行成功时微信会弹出注入成功弹窗！（http server在弹窗确认后启动）**

## 2、使用
### 2.1、路由列表
**功能类接口**
* /userinfo      - 获取登陆用户信息
* /contacts      - 获取通讯录信息（wxid从这个接口获取）
* /sendtxtmsg    - 发送文本消息

**回调注册类（目前仅用来获取微信实时消息 - 同步消息接口，同时支持WebSocket和http两种方式！）**
* /ws            - 注册websocket回调（支持注册多个ws通道）
* /sync-url  - http回调相关（支持注册多个http接口，注册请带上协议头：http/https）

### 2.2、接口使用例子
```powershell
# 获取登陆用户信息
curl 127.0.0.1:8080/userinfo

# 获取通讯录信息
curl 127.0.0.1:8080/contacts

# 发送文本消息
curl -XPOST -d'{"wxid": "44936712561@chatroom", "content": "测试内容\nHello World"}' 127.0.0.1:8080/sendtxtmsg


# 注册ws回调
# 使用任意程序websocket客户端连接127.0.0.1:8080/ws

# 注册http回调（http协议头不能少！）
curl -XPOST -d'{"url": "http://127.0.0.1:8081/callback", "timeout": 6000}' 127.0.0.1:8080/sync-url

# 获取当前已注册的http回调
curl 127.0.0.1:8080/sync-url

# 删除一个已注册的http回调
curl -XDELETE -d'{"url": "http://127.0.0.1:8081/callback"}' 10.1.12.12:8080/sync-url


# 同步消息回调响应例子（回调消息为JSON格式）
# 下面例子为反序列化后输出
# WebSocket Client Response
msg: {Wxid:wxid_gotub49l54fq29 Content:333}
recv: {"content":"你好","wxid":"34418372934@chatroom"}

# Http Server Response
msg: {Wxid:34418372934@chatroom Content:你好}
[GIN] 2023/11/12 - 12:23:21 | 200 |         135µs |     172.20.96.1 | POST     "/callback"
```

## 3、可用版本微信安装包获取
* **阿里网盘：**
https://www.aliyundrive.com/s/4eiNnE4hp4n
提取码: rt25

* **百度网盘：**
链接：https://pan.baidu.com/s/1cmzXe8AxYvzXWW2WTVCdxQ?pwd=l671 
提取码：l671

## 4、交流
请添加微信：**Anshan_PL**，备注 **wxbot** 拉微信交流群
