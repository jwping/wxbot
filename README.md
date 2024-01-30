# wxbot - 微信聊天机器人
> 适用于微信（WeChat **3.9.8.25** | 3.9.8.15 | 3.9.7.29）
> 可在Windows PC微信 **设置** - **关于微信** - **版本信息** 中获取您当前的微信版本，如果您当前的微信版本不在上述可用的版本列表中，请至下方 **3、可用版本微信安装包获取** 选择最新版微信重新安装使用

**未经过大量测试，存在封号风险！**

## 免责声明
**本仓库发布的内容，仅用于学习研究，请勿用于非法用途和商业用途！如因此产生任何法律纠纷，均与作者无关！**
**无任何后门、木马，也不获取、存储任何信息，请大家在国家法律、法规和腾讯相关原则下学习研究！**
**不对任何下载和使用者的任何行为负责，请于下载后24小时内删除！**

## 关于免注
> 因注入版本(注入DLL，功能多点)被举报了后续不再更新了，和使用方式（`READNE.md`）一起放在`wxbot-injector`目录下了
> 现在只更新免注版本了 **（免注只是不注入DLL了，并不是真的不需要注入！）**


## 1、运行
bin目录下`wxbot-sidecar.exe`，直接运行即可
* **wxbot-sidecar.exe (bin/wxbot-sidecar.exe)**

多种情况说明：
* 当无微信进程在运行时会主动拉起微信（会从注册表找微信安装目录，如果是非Setup方式安装的可能拉不起来，那么请手动启动微信并使用`-p`参数指定`pid`）
* 如微信已运行（非多开模式下）会获取当前运行中的微信进程号
* 您也可以使用`-p`参数手动指定微信进程号

```
> .\wxbot-sidecar.exe -p 30568
> .\wxbot-sidecar.exe --help
Usage: wxbot-sidecar.exe -p [pid] [ OPTIONS ] [ VALUE ]

Options:
        -p, --pid  [ pid ]              Specify the process number for injection
        -c, --config  [ path ]          Specify the configuration file path
        -a, --address                   Specify listening address (e.g. 0.0.0.0:8080)
        -b, --auto-click-login          Automatically click login when there is a login button
        -q, --qrcode-callback           If there is a QR code when WeChat is not logged in, specify the callback address for the QR code URL, using English commas as as separators (e.g. http://127.0.0.1:8081/callback,http://127.0.0.1:8082/callback)
        -w, --wechat [wechat_path,optional]     Ignoring the -p parameter will pull up a WeChat instance. Please use this -w parameter after lifting the multi opening restriction!
        -s, --silence                   Enable silent mode(without popping up the console)
        -m, --multi                     Remove WeChat multi instance restrictions (allowing multiple instances)
        -v, --version                   Output Version Information
        -h, --help                      Output Help Information
```
**目前免注入版本仅支持 3.9.8.25，请务必确认版本号正确**

### Linux下Docker部署
> 在 `Linux` 下使用 `Docker` 部署 `Wechat` + `wxbot` 全部流程已经跑通了，目前构建的公共镜像测试时间不长稳定性未知，建议各位使用时先测试
> 构建方式在`docker`路径下查看`README.md`和`Dockerfile`
```shell
ENV：
  WXBOT_ARGS：wxbot的运行参数，默认已经指定了-w和-b参数
  WINEPREFIX：指定wine运行的目录（一些驱动文件、程序目录的存放地址），基本无需变动

# 运行
# -e 指定环境变量
# -p 指定端口映射，这里是将本地环境的8080端口映射到容器内的8080端口，可按需更改
$ docker run -itd --name wxbot -e WXBOT_ARGS="-q xxx" -p 8080:8080 registry.cn-shanghai.aliyuncs.com/jwping/wxbot:v1.10.1-9-3.9.8.25
# 如果希望将微信的数据持久化出来（包括指定wxbot.json），请使用-v参数把/home/wxbot目录映射出来
$ docker run -itd --name wxbot -e WXBOT_ARGS="-q xxx" -p 8080:8080 -v xxxx:/home/wxbot registry.cn-shanghai.aliyuncs.com/jwping/wxbot:v1.10.1-9-3.9.8.25

# 可能会刷一些错误日志，不用担心，不影响使用
# 查看登陆URL，可通过百度随便找个二维码在线生成工具将登陆URL转为二维码后使用手机扫码登陆
# 再次提醒，URL不是直接访问出二维码，或者手机微信访问就可以登陆，是需要百度一个二维码在线生成，将此URL生成为二维码后使用手机微信扫码登陆！
# wxbot第一次运行时需要进行一些初始化动作，时间会较久（大约3-5分钟），此命令会占用shell不断输出，5分钟内无输出也可能是正常的，等待它初始化完成后输出登陆地址即可，期间任何报错都可以忽略，只要容器没有退出运行都可以继续等待
# 二维码扫描确认登陆后也需要一段时间才会触发服务端口监听（因为微信也有一些初始化动作），期间他也会一致刷登陆URL，这是正常的，后面再次启动就不需要了
# 直到日志出现Http Server Listen 0.0.0.0:8080，那么就可以进行访问验证了
# 正常来说初始化动作不会超过五分钟（具体视你的环境配置），如果出现了login url但是没有具体的地址并且一直无输出了那么大概率挂掉了，请尝试使用docker rm -f wxbot后重新run一个
$ docker logs -f wxbot

# 如果最终日志输出报错：
Manager Init Base faild: -3
Please review the logs and provide feedback
Manager init faild: -1

那么这可能是您当前的docker版本较低（> 20.10.8），或者在docker run时添加--security-opt seccomp=unconfined参数，完整运行命令如下：
$ docker run -itd --name wxbot -e WXBOT_ARGS="-q xxx" -p 8080:8080 --security-opt seccomp=unconfined registry.cn-shanghai.aliyuncs.com/jwping/wxbot:v1.10.1-9-3.9.8.25

# 如果您还是无法成功运行，那么请尝试registry.cn-shanghai.aliyuncs.com/jwping/wxbot:v1.10.1-8-3.9.8.25镜像，启动命令如下：
$ docker run -itd --name wxbot -e WXBOT_ARGS="-q xxx" -p 8080:8080 registry.cn-shanghai.aliyuncs.com/jwping/wxbot:v1.10.1-8-3.9.8.25

# docker运行微信是加了自动登陆点击的，所以对于二次及以上的运行会自动触发登陆，只需要等待手机上弹出登录框即可，如果启动长时间无响应请使用下面的命令重启
$ docker restart wxbot
```

### 

## 2、使用
> 如果您在使用时遇到了缺少运行库的报错
> 如：由于找不到 `MSVCP140.dll`，无法继续执行代码。重新安装程序可能会解决此问题
> 如果您遇到了此类问题可通过文档最下方的网盘链接中下载 *微软常用运行库.exe* 进行安装
> [或通过此链接下载最新微软常用运行库合集解决](https://www.lanzoux.com/b0dptvb0f)

**如果您不希望在启动时弹出命令行窗口（CMD黑框框），那么您可以使用`-s`参数以静默方式运行！**

### 2.1、自动登陆
目前已支持自动登陆，这里有两种情况：
* 之前微信已登陆过，再次启动时出现登陆按钮，如果您启动`wxbot-sidecar`时指定了`-b`参数，则会自动触发登陆，您只需在手机端确认登陆即可（**请注意，当您使用`-s`/`--silence`参数进行静默运行时`-b`/`--auto-click-login`参数不生效！**）
* 之前微信未登陆过或被取消登陆，此时会渲染出二维码图片，您可以在命令行参数中指定`-q`参数，例如：`-q http://127.0.0.1:8081/qrcode-callback` 来指定一个二维码图像的回调地址，或者您也可以在配置文件中指定（参考下方的`配置文件示例`），当然如果您就在屏幕前，您可以直接扫码登陆

配置二维码图片回调时，**命令行参数中指定的回调地址优先级要高于配置文件**


### 2.2、多开
* 如果您只有一个微信实例在运行并需要注入或快速上手，那么您无需关心其它参数，直接双击运行即可
* 如果您需要多开微信，那么请先使用`wxbot-sidecar.exe -m`解除微信多开限制（执行时机并不重要，您可以在任何情况下去解除多开限制）
* 如果您已经解除了多开限制，并希望对运行中的多个微信实例进行注入：
  * （可选）使用`-a`参数指定每个实例监听的地址（格式为：`ip:port`），**如果您的配置文件中未指定`addr`或地址冲突，那么此`-a`参数是必选项**，否则会导致端口冲突
  * （可选）使用`-c`参数指定第二个配置文件地址，如果您不指定一个新的配置文件地址，那么可能会导致您多个`wxbot-sidecar`实例共用一份配置文件（包括但不限于回调地址操作会同时影响该配置文件），推荐用法是对于多开实例每份实例使用`-c`参数去指定不同的配置文件运行
* **请注意！如果您在多开模式下希望使用`wxbot-sidecar.exe`拉起新的微信实例，那么您需要为每个微信新实例加上`-w`参数，例如：`wxbot-sidecar.exe -w` 或是 `wxbot-sidecar.exe -w -a 0.0.0.0:8081`**

### 2.3、配置文件
> 配置文件支持两种方式分别是：
> * **wxbot.json：** 这是默认的配置文件（如果有的话）
>
> **Tips：**
> **配置文件路径为 `wxbot-sidecar.exe` 所在的同级目录，或使用`-c`参数指定配置文件路径**
> **配置文件为json格式，默认不自动创建！**
> **配置文件优先级：`-c` 参数指定的配置文件 > wxbot.json > 无配置文件时的默认值**

#### 2.3.1、配置文件示例
```json
{
    "addr": "0.0.0.0:8080",
    "sync-url": {
        "qrcode": [
          {
            "timeout": 3000,
            "url": "http://localhost:8081/qrlogin-callback"
          }
        ],
        "public-msg": [
            {
                "timeout": 3000,
                "url": "http://localhost:8082/callback"
            }
        ],
        "general-msg": [
            {
                "timeout": 3000,
                "url": "http://localhost:8081/callback"
            }
        ]
    },
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
  * **qrcode：** 登陆界面出二维码图像时，二维码URL的回调地址
    * **url：** 回调url
    * **timeout：** 回调超时时间
  * **general-msg：** 普通类型消息回调地址
    * **url：** 回调url
    * **timeout：** 回调超时时间

**Tips：这里的`http://localhost:8081/callback`只是一个例子，而并非必须的，如果您未启动此回调地址，那么请删除它，否则一个不可达地址将会影响后续每个回调的到达时间！**
* **authorization：** 鉴权 **鉴权使用方法请您至下方 `5、鉴权` 了解更多**
  * **enable：** 是否开启鉴权
  * **users：** 用户列表（这是一个对象数组）
    * **user：** 用户名
    * **password：** 密码
    * **token：** 登陆后的token
* **root-dir：** 发送和获取的泛文件（图片、语音、视频、文件）存放路径，同时也会启动一个文件服务端，默认值是微信安装目录下
* **log** 
  * **level：** 日志级别：分为 **trace**、**debug**、**info**、**warn**，如果您在使用中遇到了闪退、崩溃等问题，那么请将此`log.level`指定为`trace`后收集日志反馈

**实际上配置文件中的所有字段都是非必填项，它们都可以独立存在，如果您不需要配置任何项，那么请不要创建它！**

### 2.2、路由列表
> 响应信息
> **固定为JSON格式响应：** {"code": 200, data: xxxx, "message": "xxx"}
> * **code：** 固定200
> * **message：** 成功为success，失败为faild，或是其它错误提示信息
> * **data:**  根据请求接口不同数据不同，无特别描述时下面的请求接口返回字段全部为该data字段的子字段

**路由列表概览：**
* **功能类**
  * **/api/checklogin**        - 检查登陆状态
  * **/api/userinfo**          - 获取登陆用户信息
  * **/api/contacts**          - 获取通讯录信息（wxid从这个接口获取），**不建议使用，请使用下面的`/api/dbcontacts`**
  * **/api/dbcontacts**        - 从数据库中获取通讯录信息（wxid从这个接口获取）
  * **/api/sendtxtmsg**        - 发送文本消息（好友和群聊组都可通过此接口发送，群聊组消息支持艾特）
  * **/api/sendimgmsg**        - 发送图片消息（支持json和form-data表单上传两种方式，json方式请将二进制数据使用base64编码后发送）
  * **/api/sendfilemsg**       - 发送文件消息（支持json和form-data表单上传两种方式，json方式请将二进制数据使用base64编码后发送）
  * **/api/chatroom**          - 获取群聊组成员列表，**不建议使用，请使用下面的`/api/dbchatroom`**
  * **/api/dbchatroom**        - 从数据库中获取群聊组信息和成员列表
  * **/api/accountbywxid**     - WXID反查微信昵称（支持好友、群聊组和群聊组内成员等），**不建议使用，请使用下面的`/api/dbaccountbywxid`**
  * **/api/dbaccountbywxid**   - 从数据库中通过WXID反查微信昵称（支持好友、群聊组和群聊组内成员等）
  * **/api/forwardmsg**        - 消息转发
  * **/api/dbs**               - 获取支持查询的数据库句柄
  * **/api/execsql**           - 通过数据库句柄执行`SQL`语句
  * **/api/lz4decode**         - 将lz4压缩的数据进行解码（请求数据需要base64）
  * **/close**                 - 停止 `wxbot-sidecar`（此命令用来停止`http server`，并中止程序运行）

* **回调注册类（目前仅用来获取微信实时消息 - 同步消息接口，同时支持WebSocket和http两种方式！）**
  * **/ws/generalMsg**             - 注册websocket回调（支持注册多个ws通道）：通用消息回调
  * **/ws/publicMsg**              - 注册websocket回调（支持注册多个ws通道）：订阅号（公众号）消息回调
  * **/api/syncurl**        - http回调相关（支持注册多个http接口，注册请带上协议头：http/https，注册成功会持久化到配置文件中）

#### 2.2.1、功能类接口
> **以`[]`中括号括起来的字段为可选字段**
> **目前所有请求和响应字段均按大驼峰命名法规范**

##### 2.2.1.0、检查登陆状态
**协议信息**

GET /api/checklogin

**别名**

/api/checkLogin

/api/check-login

/api/check_login

**响应字段**

* status *string*: 1: 登陆正常，其余非1值都为异常状态
* wxid *string*: 登陆用户wxid

##### 2.2.1.1、登陆用户信息
**协议信息**

GET /api/userinfo

**别名**

/api/userInfo

/api/user-info

/api/user_info

**响应字段**

* customAccount *string*: 微信号
* city *string*: 城市
* country *string*：国家
* dbKey *string*：数据库加密key，可解密读取数据库
* nickname *string*： 微信昵称
* phone *string*： 手机号
* phoneSystem *string*： 手机系统
* privateKey *string*：私钥
* profilePicture *string*： 头像
* province *string*：省
* publicKey *string*：公钥
* signature *string*：个性签名
* wxid *string*

##### 2.2.1.2、通讯录
> **不建议使用，请使用下面的从数据库中获取通讯录接口**

**协议信息**

GET /api/contacts

**响应字段**

* contacts *array*
  * customAccount *string*： 微信号
  * nickname *string*： 昵称
  * v3 *string*
  * note *string*： 备注
  * notePinyin *string*： 备注拼音首字母大写
  * notePinyinAll *string*： 备注拼音全
  * pinyin *string*： 昵称拼音首字母大写
  * pinyinAll *string*： 昵称拼音全
  * profilePicture *string*：头像
  * profilePictureSmall *string*：小头像
  * reserved1 *string*
  * reserved1 *string*
  * type *string*
  * verifyFlag *string*
  * wxid *string*
* total *uint64*： 通讯录成员总数

##### 2.2.1.3、从数据库中获取通讯录
**协议信息**

GET /api/dbcontacts

**别名**

/api/dbContacts

/api/db-contacts

/api/db_contacts

**响应字段**

* contacts *array*
  * Alias *string*： 微信号
  * NickName *string*： 昵称
  * EncryptUserName *string*：v3
  * Remark *string*： 备注
  * RemarkPYInitial *string*： 备注拼音首字母大写
  * RemarkQuanPin *string*： 备注拼音全
  * PYInitial *string*： 昵称拼音首字母大写
  * QuanPin *string*： 昵称拼音全
  * profilePicture *string*：头像
  * profilePictureSmall *string*：小头像
  * type *string*
  * UserName *string*：wxid
  * 其余字段请自测
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
> 如果您发送图片失败，那么可能是权限问题，如果您的程序工作目录（`wxbot-sidecar`所在的目录）是在C盘，那么请尝试移动到其他分区中，例如D盘，如果还未解决，请您在github上提issue或加交流群反馈

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

**`path`和`image`二选一即可，当`path`和`image`同时存在时，`path`优先**

##### 2.2.1.6、发送文件消息
> 如果您发送文件失败，那么可能是权限问题，如果您的程序工作目录（`wxbot-sidecar`所在的目录）是在C盘，那么请尝试移动到其他分区中，例如D盘，如果还未解决，请您在github上提issue或加交流群反馈

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
    * filename *string*： 文件名（使用`file`字段时，此字段必填）
* **form-data表单**
    符合标准`form-data`数据格式，需要参数分别是`wxid`、`path`和`image`

**Tips：** 当文件大小大于`5M`时则建议使用`path`文件路径的方式传参，但这并不意味着`file`不支持大文件发送，只是它需要更久的调用时间，可能是分钟级！`path`和`file`二选一即可，当`path`和`file`同时存在时，`path`优先，当使用`JSON`格式和`file`参数直接传递文件数据时`filename`是必填项！

**`path`和`file`二选一即可，当`path`和`file`同时存在时，`path`优先**


#### 2.2.1.7、获取群聊组信息
> **不建议使用，请使用下面的从数据库中获取群聊组信息**

**协议信息**
> 同时支持GET和POST

GET /api/chatroom?wxid=xxxx
POST /api/chatroom

**别名**

/api/chatRoom

/api/chat-room

/api/chat_room

**请求字段**

* **JSON：**
    * wxid *string*

**响应字段**
* data *map*
  * wxid *string*：
    * customAccount *string*：微信号
    * nickname *string*：昵称
    * note *string*：备注
    * pinyin *string*： 昵称拼音首字母大写
    * pinyinAll *string*： 昵称拼音全
    * profilePicture *string*：头像
    * profilePictureSmall *string*：小头像
    * v3 *string*

#### 2.2.1.8、从数据库中获取群聊组成员信息

**协议信息**
> 同时支持GET和POST

GET /api/dbchatroom?wxid=xxxx
POST /api/dbchatroom

**别名**

/api/dbChatRoom

/api/db-chat-room

/api/db_chat_room

**请求字段**

* **JSON：**
    * wxid *string*

**响应字段**
* data *map*
  * Announcement *string*：群公告
  * AnnouncementEditor *string*：群公告编辑人
  * AnnouncementPublishTime *string*：群公告编辑秒级时间戳
  * wxid *string*：
    * Alias *string*： 微信号
    * NickName *string*： 昵称
    * EncryptUserName *string*：v3
    * Remark *string*： 备注
    * RemarkPYInitial *string*： 备注拼音首字母大写
    * RemarkQuanPin *string*： 备注拼音全
    * PYInitial *string*： 昵称拼音首字母大写
    * QuanPin *string*： 昵称拼音全
    * profilePicture *string*：头像
    * profilePictureSmall *string*：小头像
    * type *string*
    * UserName *string*：wxid
  * 其余字段请自测

#### 2.2.1.9、WXID反查信息
> **不建议使用，请使用下面的从数据库中通过WXID反查信息**
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
* nickname *string*：昵称
* note *string*：备注
* pinyin *string*： 昵称拼音首字母大写
* pinyinAll *string*： 昵称拼音全
* profilePicture *string*：头像
* profilePictureSmall *string*：小头像
* v3 *string*

#### 2.2.1.10、从数据库中通过WXID反查信息
**协议信息**

> 同时支持GET和POST

GET /api/dbaccountbywxid?wxid=xxxx
POST /api/dbaccountbywxid

**别名**

/api/dbAccountByWxid

/api/db-account-by-wxid

/api/db_account_by_wxid

**请求字段**

* **JSON：**
    * wxid *string*

**响应字段**

* Alias *string*： 微信号
* NickName *string*： 昵称
* EncryptUserName *string*：v3
* Remark *string*： 备注
* RemarkPYInitial *string*： 备注拼音首字母大写
* RemarkQuanPin *string*： 备注拼音全
* PYInitial *string*： 昵称拼音首字母大写
* QuanPin *string*： 昵称拼音全
* profilePicture *string*：头像
* profilePictureSmall *string*：小头像
* type *string*
* UserName *string*：wxid
* 其余字段请自测

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
    * msgId *uint64|string*：消息id（通常可以用消息回调或者`websocket`回调获取到，当前是消息回调中的`MsgSvrID`字段）


#### 2.2.1.12、获取数据库句柄
**协议信息**

GET /api/dbs

**响应字段**

*map<string, uint64>*
* *string*： 数据库名
* *uint64*： 句柄

#### 2.2.1.13、执行`SQL`语句
**协议信息**

> 同时支持GET和POST

GET /api/execsql?dbName=PublicMsg.db&sql='select * from PublicMsg ORDER BY localId DESC limit 1' *（这是一个实际可用的例子，前提是你收到过订阅号消息）*
POST /api/execsql

**别名**

/api/execSql

/api/exec-sql

/api/exec_sql

**请求字段**

* **JSON：**
    * dbName *string*：需要执行`SQL`的数据库名
    * sql *string*：需要执行的`SQL`语句

**可执行的SQL例子：**

**PublicMsg.db**

查询指定公众号的所有文章（本地已接受的）
`SELECT * FROM PublicMsg WHERE StrTalker=='gh_13508120ed24'`

查询指定时间范围的所有订阅号文章（20230115全天的）
`SELECT * FROM PublicMsg WHERE CreateTime>1705248000 AND CreateTime<1705334399;`

分页查询，从第21行开始，累加20条数据进行检索
`SELECT * FROM PublicMsg ORDER BY localId DESC limit 20,20;`

**MicroMsg.db**

查询通讯录所有成员（包括好友、群聊组、订阅号等）
`SELECT * FROM Contact`

#### 2.2.1.14、lz4数据解码
> 提交的数据需要经过`Base64 Encode`

**协议信息**

> 同时支持GET和POST

GET /api/lz4decode?compressContent=xxxxxxxxxxx
POST /api/lz4decode

**别名**

/api/lz4Decode

/api/lz4-decode

/api/lz4_decode

**请求字段**

* **JSON：**
    * compressContent *string*：经过`Base64 Encode`的lz4压缩数据

**响应字段**

* content *string*：lz4解压缩后的数据

#### 2.2.1.15、停止程序
**协议信息**

GET /close

停止 `wxbot-sidecar`（此命令用来停止`http server`，并中止程序运行），当以静默方式运行时此命令可以比较方便的用来停止程序

#### 2.2.2、回调注册类
> 目前已知BUG是部分环境/微信号有登陆后微信崩溃的问题，因为我本地环境均未复现出该问题，所以修复进度较慢，但修复中...
> 如果您愿意提供程序日志我不胜感激

##### 2.2.2.1、登陆二维码回调（qrcode）
**响应字段**
* url *string*：登陆二维码URL（需在微信中渲染为二维码后扫码）

##### 2.2.2.2、订阅号消息回调（public-msg）
* wxid *string*：当前实例登陆用户的wxid
* total *uint32*：每次回调的消息数量
* data：
  * BytesExtra *string*：扩展字段BASE64后的二进制数据
  * BytesTrans *string*
  * Content *string*：订阅号号XML消息
  * CreateTime *string*：秒级时间戳
  * IsSender *string*：是否是自己发出的消息（0：非自己发送、1：自己发送）
  * StrTalker *string*：订阅号发送者微信ID（wxid）
  * SubType *string*：消息类型子类，例如视频消息大类下可能存在小程序等小类的区分
  * Type *string*：消息类型
  * localId *string*：本地数据库ID，目前来看是一个自增ID
  * MsgSvrID *string*：消息id
  * StatusEx、FlagEx、Status、MsgServerSeq、MsgSequence、Reserved0-6、TalkerId 未知

##### 2.2.2.3、普通消息回调（general-msg）
**响应字段**
* wxid *string*：当前实例登陆用户的wxid
* total *uint32*：每次回调的消息数量
* data：
  * BytesExtra *string*：扩展字段BASE64后的二进制数据
  * BytesTrans *string*
  * StrContent *string*：字符串数据，除文本消息以为大部分均为XML数据
  * Content *string*：引用消息、用户转发的订阅号消息等
  * CreateTime *string*：秒级时间戳
    * 从PC登陆微信上发出的消息：标记代表的是每个消息点下发送按钮的那一刻
    * 从其它设备上发出的/收到的来自其它用户的消息：标记的是本地从服 务器接收到这一消息的时间
  * DisplayContent *string*：拍一拍，邀请入群等消息
  * IsSender *string*：是否是自己发出的消息（0：非自己发送、1：自己发送）
  * StrTalker *string*：消息发送者微信ID（wxid）
  * SubType *string*：消息类型子类，例如视频消息大类下可能存在小程序等小类的区分
  * Type *string*：消息类型
  * localId *string*：本地数据库ID，目前来看是一个自增ID
  * MsgSvrID *string*：消息id
  * Sender *string*：群聊消息发送人的wxid（仅在消息为chatroom群聊消息时存在该字段）
  * StatusEx、FlagEx、Status、MsgServerSeq、MsgSequence、Reserved0-6、TalkerId 未知

##### 2.2.2.1、websocket协议消息
**协议信息**

订阅号消息
GET ws://xxxxx/ws/publicMsg

普通消息
GET ws://xxxxx/ws/generalMsg

消息体参考回调响应字段

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
* *type： string*
  * `qrcode`：二维码消息回调（仅存在登陆二维码时触发）
  * `public-msg`：订阅号消息回调
  * `general-msg`：普通消息回调

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
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "content": "测试内容\nhello world!"}' http://127.0.0.1:8080/api/sendtxtmsg

# 发送艾特消息
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "content": "测试内容\nhello world!", "atlist": ["被艾特人的wxid"]}' http://127.0.0.1:8080/api/sendtxtmsg

# 发送图片消息
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "image": "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="}' http://127.0.0.1:8080/api/sendimgmsg

# 发送文件消息
curl -Method POST -ContentType "application/json" -Body '{"wxid":"47331170911@chatroom", "file": "aGVsbG8gd29ybGQ=", "filename": "1.txt"}' http://127.0.0.1:8080/api/sendfilemsg

# 注册普通消息回调
curl -Method POST -ContentType "application/json" -Body '{"url":"http://127.0.0.1:8011", "timeout": 3000, "type": "general-msg"}' http://127.0.0.1:8080/api/syncurl
```

**Linux**
```bash
# 获取登陆用户信息
curl 127.0.0.1:8080/api/userinfo

# 获取通讯录信息
curl 127.0.0.1:8080/api/contacts

# 发送文本消息
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "content": "测试内容\nHello World"}' 127.0.0.1:8080/api/sendtxtmsg

# 发送图片消息1（使用form-data表单方式提交）
curl -XPOST -F "wxid=47331170911@chatroom" -F "image=@/home/jwping/1.jpg" 127.0.0.1:8080/api/sendimgmsg
# 发送图片消息2（使用json方式提交）
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "image": "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="}' 127.0.0.1:8080/api/sendimgmsg

# 发送文件消息1（使用form-data表单方式提交）
curl -XPOST -F "wxid=47331170911@chatroom" -F "file=@/home/jwping/1.txt" 127.0.0.1:8080/api/sendfilemsg
# 发送文件消息2（使用json方式提交）
curl -XPOST -H "Content-Type: application/json" -d'{"wxid": "47331170911@chatroom", "filename": "1.txt", "file": "aGVsbG8gd29ybGQh"}' 127.0.0.1:8080/api/sendfilemsg

# 注册ws回调
# 使用任意程序websocket客户端连接127.0.0.1:8080/ws

# 注册http 普通消息回调（http协议头不能少！）
curl -XPOST -d'{"url": "http://127.0.0.1:8081/callback", "timeout": 6000, "type": "general-msg"}' 127.0.0.1:8080/api/sync-url

# 获取当前已注册的http回调
curl 127.0.0.1:8080/api/sync-url

# 删除一个已注册的http回调
curl -XDELETE -d'{"url": "http://127.0.0.1:8081/callback"}' 127.0.0.1:8080/api/sync-url
```

## 3、赞助码
**如果觉得本项目对你有帮助，可以打赏一下作者，毕竟开源不易**

<img src=https://raw.githubusercontent.com/jwping/wxbot/main/public/wechat_collection.jpg width=40% />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<img src=https://raw.githubusercontent.com/jwping/wxbot/main/public/alipay_collection.jpg width=40% />

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


## 6、wxbox、可用版本微信安装包等获取
* **阿里网盘：**
https://www.aliyundrive.com/s/4eiNnE4hp4n
提取码: rt25

* **百度网盘：**
https://pan.baidu.com/s/1cmzXe8AxYvzXWW2WTVCdxQ?pwd=l671 
提取码：l671

## 7、交流
### 7.1、微信
请添加微信：**Anshan_PL**，备注 **wxbot** 拉微信交流群

**Tips：此群仅限学习和交流，无其他用处**

### 7.2、TG
Android端下载地址： https://telegram.org/android 其他客户端从这个网址找过去

安装后复制下面的链接到tg中打开： https://t.me/+DVigUtfAIOthNmNl

**Tips：此TG群同样仅限学习和交流，无其他用处**
