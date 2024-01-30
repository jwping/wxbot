# wxbot - 微信聊天机器人
> 适用于微信（WeChat **3.9.8.15** | 3.9.7.29）
> 可在Windows PC微信 **设置** - **关于微信** - **版本信息** 中获取您当前的微信版本，如果您当前的微信版本不在上述可用的版本列表中，请至下方 **3、可用版本微信安装包获取** 选择最新版微信重新安装使用

**未经过大量测试，使用远程线程注入方式可能会被报毒（无毒，请放心使用！），也可以尝试其它方式进行注入，注入手段并不重要，只要将wxbot.dll注入到wechat.exe中即可**

## 免责声明
**`DLL`注入是已被废弃的方式，请使用`wxbot-sidecar`！**

**本仓库发布的内容，仅用于学习研究，请勿用于非法用途和商业用途！如因此产生任何法律纠纷，均与作者无关！**
**无任何后门、木马，也不获取、存储任何信息，请大家在国家法律、法规和腾讯相关原则下学习研究！**
**不对任何下载和使用者的任何行为负责，请于下载后24小时内删除！**


## 1、运行
bin目录下有如下两个文件（仅在windows 10 & windows server 2012 R2系统上进行测试）：
* **injector.exe (bin/injector.exe)**
* **wxbot.dll (bin/wxbot.dll)**

简单方式可直接运行 `injector.exe` 即可自动拉起微信并完成注入！
默认wxbot.dll为最新版（3.9.8.15），低版本微信注入请选择对应版本的wxbot-xxxx.dll替换为wxbot.dll后注入即可

### Linux下Docker部署
> 在 `Linux` 下使用 `Docker` 部署 `Wechat` + `wxbot` 全部流程已经跑通了，后面我会构建成一个公共镜像供大家使用（但使用 `wine` 运行 `WeChat` 的稳定性如何到时还需要大家帮忙一起测试了）

## 2、使用
> **如果您在运行注入器（`injector.exe`）时遇到了缺少运行库的报错**
> **如：由于找不到 `MSVCP140.dll`，无法继续执行代码。重新安装程序可能会解决此问题**
> **如果您遇到了此类问题可通过文档最下方的网盘链接中下载 *微软常用运行库.exe* 进行安装**
> [或通过此链接下载最新微软常用运行库合集解决](https://www.lanzoux.com/b0dptvb0f)

### 2.1、注入器
> 注入器（`injector.exe`）目前支持：
> * 直接运行可启动微信并完成自动注入
> * 按微信PID（进程号）注入
> * 注入卸载（当您开启了隐藏注入时，此功能不可用！）
> * 指定注入的DLL路径
> * 开启微信多开（取消微信多开限制）
> 
> 具体使用方式您可以使用`inject.exe --help`命令查看！
```powershell
$ injector.exe --help
Usage: injector.exe [ OPTIONS ] [ VALUE ]

Options:
        -p, --pid  [ pid ]              Specify the process number for injection
        -d, --dll  [ path ]             Specify the DLL path to be injected
        -i, --inject                    inject DLL
        -u, --uninject                  uninject DLL
        -m, --multi                     Remove WeChat multi instance restrictions (allowing multiple instances)
        -s, --silence                   Enable silent mode(without popping up the console)
        -h, --help                      Output Help Information
```
**Tips：**
* 如果您只有一个微信实例在运行并需要注入，那么您无需关心其它参数，直接运行即可注入
* 如果您需要多开微信，那么请使用`inject.exe -m`解除微信多开限制（执行时机并不重要，您可以在任何情况下去解除多开限制），**但是您需要注意，如果您已经运行了多个微信实例，那么此时请不要尝试直接运行inject.exe进行注入了，而是使用`-p`参数对指定微信进程号进行注入！**
* 如果您已经解除了多开限制，并希望对运行中的多个微信实例进行注入，那么您需要使用`-p`参数对每个微信的进程号进行注入

### 2.2、配置文件
> 配置文件支持两种方式分别是：
> * **[wxid].json：** 支持登陆用户wxid的专属配置文件，如你登陆的微信用户wxid是abc，且微信根目录下有abc.json配置文件的话则优先读取此配置文件！
> * **wxbot.json：** 这是默认的配置文件（如果有的话）
>
> **Tips：**
> **配置文件路径为 `WeChat.exe` 所在的同级目录（也就是微信的安装目录）**
> **配置文件为json格式，默认不自动创建！**
> **配置文件优先级：[wxid].json > wxbot.json > 无配置文件时的默认值**
> 
> **这样设计配置文件优先级是为了适配微信多开而不那么优雅的实现方式，具体您可以看 `4. 多开高级用法` 了解更多**

#### 2.2.1、配置文件示例
```json
{
    "addr": "0.0.0.0:8080",
    "sync-url": {
        "msg": [
            {
                "timeout": 3000,
                "url": "http://localhost:8081/callback"
            }
        ],
        "msg2": [
            {
                "timeout": 3000,
                "url": "http://localhost:8082/callback"
            }
        ]
    },
    "hide-module": false,
    "authorization": {
        "enable": false,
        "users": [
            {
                "user": "admin",
                "password": "123",
                "token": "token123"
            },
            {
                "user": "user",
                "password": "321",
                "token": "token321"
            }
        ]
    },
    "root-dir": "D:\\WeChat",
    "log": {
        "level": "info"
    }
}
```
* **addr:** wxbot服务监听地址（固定为ip:port形式）
* **sync-url：** http回调地址列表（建议通过下面的/sync-url接口修改，不要手动修改）
  * **msg{2}：**
    * **url：** 回调url
    * **timeout：** 回调超时时间

**关于`msg`和`msg2`的区别：**
**`msg`：为通用回调：好友、群聊、公众号、微信提示，反正各种乱七八糟的都有**
**`msg2`：为普通回调：仅支持好友、群聊消息等，但这个回调可以拿到pc微信客户端界面发送的消息！**

**Tips：这里的`http://localhost:8081/callback`只是一个例子，而并非必须的，如果您未启动此回调地址，那么请删除它！**
* **hide-module：** 是否隐藏注入，当开启隐藏注入时 inject.exe 的注入卸载将不可用，此时您只能通过重启微信的方式来卸载DLL！
* **authorization：** 鉴权 **鉴权使用方法请您至下方 `5、鉴权` 了解更多**
  * **enable：** 是否开启鉴权
  * **users：** 用户列表（这是一个对象数组）
    * **user：** 用户名
    * **password：** 密码
    * **token：** 登陆后的token
* **root-dir：** 发送和获取的泛文件（图片、语音、视频、文件）存放路径，同时也会启动一个文件服务端，默认值是微信安装目录下
* **log** 
  * **level：** 日志级别：分为 **trace**、**debug**、**info**、**warn**

**实际上配置文件中的所有字段都是非必填项，它们都可以独立存在，如果您不需要配置任何项，那么请不要手动创建配置文件！**

### 2.2、路由列表
> 响应信息
> **固定为JSON格式响应：** {"code": 200, data: xxxx, "message": "xxx"}
> * **code：** 固定200
> * **message：** 成功为success，失败为faild，或是其它错误提示信息
> * **data:**  根据请求接口不同数据不同，无特别描述时下面的请求接口返回字段全部为该data字段的子字段

**路由列表概览：**
* **功能类**
  * **/api/checklogin**        - 检查当前是否已登录（兼容接口，没什么用）
  * **/api/userinfo**          - 获取登陆用户信息
  * **/api/contacts**          - 获取通讯录信息（wxid从这个接口获取）
  * **/api/sendtxtmsg**        - 发送文本消息（好友和群聊组都可通过此接口发送，群聊组消息支持艾特）
  * **/api/sendimgmsg**        - 发送图片消息（支持json和form-data表单上传两种方式，json方式请将二进制数据使用base64编码后发送）
  * **/api/sendfilemsg**       - 发送文件消息（支持json和form-data表单上传两种方式，json方式请将二进制数据使用base64编码后发送）
  * **/api/chatroom**          - 获取群聊组信息，包括：管理员、公告、成员列表等
  * **/api/accountbywxid**     - WXID反查微信昵称（支持好友和群聊组等）
  * **/api/sendcardmsg**       - 发送卡片消息
  * **/api/getgeneralfile**    - 获取泛文件（图片、语音、视频、文件）
  * **/api/forwardmsg**        - 消息转发
  * **/api/close**             - 注入卸载（仅卸载 `wxbot.dll` 的注入，不会关闭或重启微信）

* **回调注册类（目前仅用来获取微信实时消息 - 同步消息接口，同时支持WebSocket和http两种方式！）**
  * **/ws/msg**             - 注册websocket回调（支持注册多个ws通道）：通用消息回调
  * **/ws/msg2**            - 注册websocket回调（支持注册多个ws通道）：普通消息回调
  * **/api/syncurl**        - http回调相关（支持注册多个http接口，注册请带上协议头：http/https，注册成功会持久化到配置文件中）

#### 2.2.1、功能类接口
> **以`[]`中括号括起来的字段为可选字段**
> **目前所有请求和响应字段均按大驼峰命名法规范**

##### 2.2.1.1、检查当前是否已登录
**协议信息**

GET /api/checklogin

**别名**

/api/checkLogin

/api/check-login

/api/check_login

**响应字段**

* status *uint64*: 当前登陆状态：0 未登陆 - 1 已登陆

##### 2.2.1.2、登陆信息
**协议信息**

GET /api/userinfo

**别名**

/api/userInfo

/api/user-info

/api/user_info

**响应字段**

* customAccount *string*: 微信号
* nickname *string*： 微信昵称
* phone *string*： 手机号
* phoneSystem *string*： 手机系统
* profilePicture *string*： 头像
* profilePictureSmall *string*： 小头像
* wxid *string*

##### 2.2.1.3、通讯录
**协议信息**

GET /api/contacts

**响应字段**

* contacts *array*
  * customAccount *string*： 微信号
  * nickname *string*： 昵称
  * note *string*： 备注
  * pinyin *string*： 昵称拼音首字母大写
  * pinyinAll *string*： 昵称拼音全
  * type1 *uint64*： 用户类别1
  * type2 *uint64*： 用户类别2
  * wxid *string*
* total *uint64*： 通讯录成员总数

##### 2.2.1.4、发送文本消息
> 对于群聊组消息发送支持艾特

**协议信息**

POST /api/sendtxtmsg

**别名**

/api/sendTxtMsg

/api/send-txt-msg

/api/send_txt_msg

**请求字段**

* wxid *string*
* content *string*：发送消息内容（如果是群聊组消息并需要发送艾特时，**此content字段中需要有对应数量的`@[自定义被艾特人的昵称，不得少于2个字符] [每个艾特后都需要一个空格以进行分隔（包括最后一个艾特！）]`，这一点很重要！ 如果您不理解，请继续看下面的Tips！**）
* [atlist] *array\<string\>*：如果是群聊组消息并需要发送艾特时，此字段是一个被艾特人的数组

**Tips：如果是群聊艾特消息，那么`content`字段中的`@`艾特符号数量需要和`atlist`中的被艾特人数组长度一致，简单来说，就是`atlist`中有多少个被艾特人的`wxid`，那么`content`字段中就需要有多少个艾特组合，位置随意，例如：**
`{"wxid": "xx@chatroom", "content": "这里@11 只是@22 想告诉你@33 每个被艾特人的位置并不重要", "atlist": ["wxid_a", "wxid_b", "wxid_c"]}`
**每个被艾特人在`content`中 固定为`@[至少两个字符的被艾特人名] + 一个空格`！**
**如果是发送`@所有人`消息，那么请在`atlist`字段中仅传入一个`notify@all`字符串，`content`字段中仅包含一个`@符号规范（最少两字符+一个空格）`， 一般建议是`@所有人`见名知意**

**响应示例**

{"code":200,"msg":"success"}

##### 2.2.1.5、发送图片消息
**协议信息**

POST /api/sendimgmsg

**别名**

/api/sendImgMsg

/api/send-img-msg

/api/send_img_msg

/api/sendimagemsg

/api/sendImageMsg

/api/send-image-msg

/api/send_image_msg
> 支持JSON和form-data表单两种方式提交

**请求头**

* **JSON：`Content-Type: application/json`**
* **form-data表单：`Content-Type: multipart/form-data`**

**请求字段**

* **JSON：**
    * wxid *string*
    * path *string*：图片路径（注意，这里的图片路径是bot登陆系统的路径！）
    * image *string*： 图片二进制数据base64编码后字符串 **（不需要加 `data:image/jpeg;base64,` 前缀）**
    * clear *bool*： 指定图片发送后是否需要删除，默认删除 **（需要注意的是，图片文件保存后并没有后缀，这意味着如果您需要查看历史发送图片，那么您需要至`[微信根目录]/temp`自行查看判断图片格式并添加后缀）**

* **form-data表单**
    符合标准`form-data`数据格式，需要参数分别是`wxid`、`path`和`image`

`path`和`image`二选一即可，当`path`和`image`同时存在时，`path`优先

##### 2.2.1.6、发送文件消息
**协议信息**

POST /api/sendfilemsg

**别名**

/api/sendFileMsg

/api/send-file-msg

/api/send_file_msg
> 支持JSON和form-data表单两种方式提交

**请求头**

* **JSON：`Content-Type: application/json`**
* **form-data表单：`Content-Type: multipart/form-data`**

**请求字段**

* **JSON：**
    * wxid *string*
    * path *string*：文件路径（注意，这里的文件路径是bot登陆系统的路径！）
    * file *string*： 文件二进制数据base64编码后字符串
    * filename *string*： 文件名
* **form-data表单**
    符合标准`form-data`数据格式，需要参数分别是`wxid`、`path`和`image`

**Tips：** 当文件大小大于`5M`时则建议使用`path`文件路径的方式传参，但这并不意味着`file`不支持大文件发送，只是它需要更久的调用时间，可能是分钟级！`path`和`file`二选一即可，当`path`和`file`同时存在时，`path`优先，当使用`JSON`格式和`file`参数直接传递文件数据时`filename`是必填项！

#### 2.2.1.7、获取群聊组信息
**协议信息**
> 同时支持GET和POST

GET /api/chatroom?wxid=xxxx&account=true
POST /api/chatroom

**别名**

/api/chatRoom

/api/chat-room

/api/chat_room

**请求字段**

* **JSON：**
    * wxid *string*
    * account *bool*：为 `true` 时输出中会将每个成员的微信昵称反查带出

**响应字段**

* admin1 *string*：群聊组管理员
* admin2 *string*：一般同上
* adminNickname *string*：管理员昵称
* notice *string*： 群公告
* pinyinAll *string*： 昵称拼音全
* member *array\<string\>*： 成员wxid列表
* memberNickname *map\<string, string\>*： 成员群聊昵称（MAP类型）
* [memberAccount] *map\<string, Object\>*： 当请求参数中`account`为`true`时存在此字段，map类型，成员微信昵称，value是一个对象，字段为`wxid反查昵称的所有字段`
* xml
* wxid

#### 2.2.1.8、WXID反查微信昵称
**协议信息**

> 同时支持GET和POST

GET /api/accountbywxid?wxid=xxxx
POST /api/accountbywxid

**别名**

/api/accountByWxid

/api/account-by-wxid

/api/account_by_wxid

**请求字段**

* **JSON：**
    * wxid *string*

**响应字段**

* customAccount *string*：微信号
* nickname *string*：微信昵称
* pinyin *string*：拼音
* pinyinAll *string*：拼音全
* profilePicture *string*：头像链接
* profilePictureSmall *string*：小头像（群聊组仅有小头像，没有`profilePicture`）
* v3 *string*
* wxid *string*

#### 2.2.1.9、发送卡片消息
**协议信息**

POST /api/sendcardmsg

**别名**

/api/sendCardMsg

/api/send-card-msg

/api/send_card_msg

**请求字段**

* **JSON/form-data：**
    * wxid *string*
    * title *string*：卡片标题
    * url *string*：卡片链接
    * [digest *string*]：卡片简介
    * [image *string*]：卡片右侧小图（传url）
    * [subscriptionAccountId *string*]：订阅号id **此字段在高级版中支持**
    * [subscriptionAccountName *string*]：订阅号昵称 **此字段在高级版中支持**

#### 2.2.1.10、获取泛文件
> 为什么叫泛文件呢，因为这个接口可以用来获取图片（大图、原图）、语音、视频、以及通常我们所说的文件
**协议信息**
> 同时支持GET和POST

GET /api/getgeneralfile?msgId=xxxxxxxxxxxx
POST /api/getgeneralfile

**别名**

/api/getGeneralFile

/api/get-general-file

/api/get_general_file

**请求字段**
> 这里说明一下，因为前端精度问题，有些大佬可能传递`msgId`字段时存在精度丢失或自动转字符串的问题，所以这里我将`msgId`字段设置为了同时支持`uint64`和`string`两种类型！

* **JSON：**
    * msgId *uint64|string*：消息id（通常可以用消息回调或者`websocket`回调获取到，`msg`和`msg2`都可以）

**响应字段**

* thumb *string*：缩略文件（比如图片这里可能是缩略图、视频可能是封面，小程序消息可能也是一个封面）
* file *string*：完整文件

**但是这里注意一下，这里返回的是一个文件路径，我们该如何获取呢？**
**这一版开始wxbot会启动一个文件服务，将配置文件中指定的`root-dir`目录映射出来（如果未指定则为微信安装目录下），那么你只需要进行拼接就可以获取了，例如：**
**"http://127.0.0.1:8080/"+file**

#### 2.2.1.11、转发消息
**协议信息**
> 同时支持GET和POST

GET /api/forwardmsg?wxid=xxxxxxxxxxx&msgId=xxxxxxxxxxxx
POST /api/forwardmsg

**别名**

/api/forwardMsg

/api/forward-msg

/api/forward_msg

**请求字段**
> 这里说明一下，因为前端精度问题，有些大佬可能传递`msgId`字段时存在精度丢失或自动转字符串的问题，所以这里我将`msgId`字段设置为了同时支持`uint64`和`string`两种类型！

* **JSON：**
    * wxid *string*：本次转发消息的接收对象
    * msgId *uint64|string*：消息id（通常可以用消息回调或者`websocket`回调获取到，`msg`和`msg2`都可以）

##### 2.2.1.10、卸载
> 这个接口是用来卸载已注入的 `wxbot.dll`，而不关闭微信，可以算是 `injector.exe -u` 的平替

**协议信息**

GET /api/close

#### 2.2.2、回调注册类
> 目前仅用来同步微信消息

**响应字段**

* wxid *string*： 发送消息的消息人/群聊组wxid
* customAccount *string*： 发送消息的消息人微信号（如果是群聊组此字段为空）
* nickname *string*：发送消息的消息人/群聊组昵称
* content *string*： 消息内容
* toUser *string*： 消息接收人（一般为登陆用户wxid）
* msgid *uint64*： 消息唯一标识
* originMsg *string*： 原始消息（如：wxid:\nxxxxxxx）
* chatRoomSourceWxid *string*： 如果为群聊消息，则为消息发送人wxid
* chatRoomSourceCustomAccount *string*： 如果为群聊消息，则为消息发送人微信号
* chatRoomSourceNickname *string*： 如果为群聊消息，则为消息发送人微信昵称
* msgSource *string*： 加密消息
* type *uint32*： 消息类型
* displayMsg *string*： 展示消息（一般群聊消息下可用）
* [imgData] *string*： base64后的图片数据字符串

**Tips：当`type`为`3`时表示当次消息是图片消息，本次消息中会新增`imgData`字段**

##### 2.2.2.1、websocket协议消息
**协议信息**

GET ws://xxxxx/ws/msg{1|2}

> websocket没什么好说的，基本上第三方库都有直接可用的实现，协议升级后就是一条全双工通道，目前只用来接收同步微信的实时消息，不要发送消息到服务端，服务端不会响应。

##### 2.2.2.2、http协议
> 需要你自己起一个Http Server服务用来接收微信的实时消息，你自己的Http Server启动之后通过接口注册到wxbot即可

##### 2.2.2.2.1、注册接口
POST /api/syncurl

**别名**

/api/syncUrl

/api/sync-url

/api/sync_url

**请求字段**

* url *string*： 你自己启动的Http Server地址路由（**ip:port/[subpath]**）
* timeout *int*： 超时时间（当有一条新消息通过wxbot发送到你的回调地址时的最长连接等待时间）
* *type： string* 此版本开始同时支持`msg`和`msg2`两种回调（默认值：`msg`）

**Tips: **
**`msg`：为通用回调：好友、群聊、公众号、微信提示，反正各种乱七八糟的都有**
**`msg2`：为普通回调：仅支持好友、群聊消息等，但这个回调可以拿到pc微信客户端界面发送的消息！**

##### 2.2.2.2.2、获取已注册接口列表
GET /api/syncurl

**别名**

/api/syncUrl

/api/sync-url

/api/sync_url

##### 2.2.2.2.3、删除接口
DELETE /api/syncurl

**别名**

/api/syncUrl

/api/sync-url

/api/sync_url

**请求字段**

* url： 已注册的Http Server地址（**ip:port/[subpath]**）
* *type： string* 此版本开始同时支持`msg`和`msg2`两种回调（默认值：`msg`）

### 2.3、接口使用例子
**Windows**
**所有`powershell`或者是使用`cmd`测试发送的例子都可能有编码问题！建议直接用程序测试！**
```powershell
# 发送文本
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "content": "测试内容\nhello world!"}' http://127.0.0.1:8080/sendtxtmsg

# 发送图片
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "path": "D:\\gopath\\wxbot\\测试.txt"}' http://127.0.0.1:8080/sendfilemsg

# 发送卡片消息（例子仅仅是卡片消息，订阅号消息等待高级版支持）
curl -Method POST -ContentType "application/json" -Body '{"wxid": "47331170911@chatroom", "title": "测试标题", "url": "https://www.baidu.com"}' http://127.0.0.1:8080/sendcardmsg
```

**Linux**
```bash
# 获取登陆用户信息
curl 127.0.0.1:8080/userinfo

# 获取通讯录信息
curl 127.0.0.1:8080/contacts

# 发送文本消息
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "content": "测试内容\nHello World"}' 127.0.0.1:8080/sendtxtmsg

# 发送图片消息1（使用form-data表单方式提交）
curl -XPOST -F "wxid=47331170911@chatroom" -F "image=@/home/jwping/1.jpg" 127.0.0.1:8080/sendimgmsg
# 发送图片消息2（使用json方式提交）
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "image": "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"}' 127.0.0.1:8080/sendimgmsg

# 发送文件消息1（使用form-data表单方式提交）
curl -XPOST -F "wxid=47331170911@chatroom" -F "file=@/home/jwping/1.txt" 127.0.0.1:8080/sendfilemsg
# 发送文件消息2（使用json方式提交）
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "filename": "1.txt", "file": "aGVsbG8gd29ybGQh"}' 127.0.0.1:8080/sendfilemsg

# 发送卡片消息
curl -XPOST -d'{"wxid": "47331170911@chatroom", "title": "测试标题", "url": "www.baidu.com"}' 127.0.0.1:8080/sendcardmsg

# 注册ws回调
# 使用任意程序websocket客户端连接127.0.0.1:8080/ws

# 注册http回调（http协议头不能少！）
curl -XPOST -d'{"url": "http://127.0.0.1:8081/callback", "timeout": 6000}' 127.0.0.1:8080/sync-url

# 获取当前已注册的http回调
curl 127.0.0.1:8080/sync-url

# 删除一个已注册的http回调
curl -XDELETE -d'{"url": "http://127.0.0.1:8081/callback"}' 127.0.0.1:8080/sync-url


# 同步消息回调响应例子（回调消息为JSON格式）
# 下面例子为反序列化后输出
# WebSocket Client Response
{Wxid:34418372934@chatroom Content:你好 ToUser:wxid_gotub49l54fq29 Msgid:7438040783824576403 OriginMsg:wxid_o4jinvsgz6lp31:
你好 ChatRoomSourceWxid:wxid_o4jinvsgz6lp31 MsgSource:<msgsource>
        <pua>1</pua>
        <silence>1</silence>
        <membercount>193</membercount>
        <signature>v1_8TXCDRkh</signature>
        <tmp_node>
                <publisher-id></publisher-id>
        </tmp_node>
</msgsource>
 Type:1 DisplayMsg:}
```

## 3、赞助码
**如果觉得本项目对你有帮助，可以打赏一下作者，毕竟开源不易**

<img src=https://raw.githubusercontent.com/jwping/wxbot/main/public/wechat_collection.jpg width=40% />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<img src=https://raw.githubusercontent.com/jwping/wxbot/main/public/alipay_collection.jpg width=40% />

## 4、微信多开高级用法
这里仅仅给出一种为每个wxbot指定端口和回调地址等使用思路：
当您使用`inject.exe -m`解开微信多开限制后，可以在微信根目录下为每个wxbot生成一个[wxid].json的配置文件，以此来为不同的wxbot定义不同的监听地址，例如：

假设您有两个wxid为`wxid_a`和`wxid_b`的两个微信号希望实现多开注入，
那么您可以在您的微信根目录下分别生成`wxid_a.json`和`wxid_b.json`两个配置文件：
```powershell
# wxid_a.json配置文件内容如下：
{"addr": "0.0.0.0:8080"}

# wxid_b.json配置文件内容如下：
{"addr": "0.0.0.0:8081"}

# 配置文件生成好之后，您可以使用注入器对两个微信bot分次注入
# 第一次注入
inject.exe -p [wxid_a的微信PID]

# 第二次注入
inject.exe -p [wxid_b的微信PID]
```
至此，您就完成了对两个微信号的注入，并且这两个wxbot分别监听在`8080`和`8081`端口
其中`wxid_a`监听在`8080`端口
其中`wxid_b`监听在`8081`端口

## 5、鉴权
> 当您在配置文件中开启了鉴权之后则 **您后续的每个api请求都需要包含鉴权信息！**
> 这里使用的是`Http Basic Authentication`，您可以先百度去了解一下它，当然，如果您不想了解也没关系，因为它真的很简单

```json
# 假设您定义一个如下用户：
{
    "user": "user2",
    "password": "321",
    "token": "token321"
}
```
那么您需要在您后续的每次请求的请求头中加上`Authorization`字段：`Authorization: Basic base64(username:password)`
例如用curl命令请求的话，它可能长这样：
```bash
curl -H "Authorization: Basic dXNlcjI6MzIx" 127.0.0.1:8080/login -v

# response：
{"code":200,"data":{"token":"token321","user":"user2"},"msg":null}
```
这里**引入了一个新的路由`/login`**，但我并不想将他写到上面的路由列表中，因为它真的没什么用，仅仅是在您登陆成功之后返回一个当前的登陆用户名和`token`
**您在以后每个请求都加上`Authorization: Basic dXNlcjI6MzIx`这个请求头就可以了！**

如果您希望使用`cookie`的方式，那么您可以在`cookie`中指定`token`
例如，您也可以这样做：
```bash
curl --cookie "access_token=token321" 127.0.0.1:8080/userinfo -v
```
**如果您不想用设置请求头的方式，那么您也可以在后续的所有请求的`cookie`中指定`access_token`字段即可。** *实际上`cookie`也是`request header`中的一个字段*


## 6、wxbox.dll、注入器、可用版本微信安装包等获取
* **阿里网盘：**
https://www.aliyundrive.com/s/4eiNnE4hp4n
提取码: rt25

* **百度网盘：**
https://pan.baidu.com/s/1cmzXe8AxYvzXWW2WTVCdxQ?pwd=l671 
提取码：l671

## 7、交流
请添加微信：**Anshan_PL**，备注 **wxbot** 拉微信交流群

**Tips：此群仅限学习和交流，无其他用处**