# Golang 客户端例子
> 这个例子并不干净、也并不完整，仅仅是我需要时临时测试的小demo！

## 使用
```powershell
# 使用websocket客户端回调注册
$ ./go-client -addr 127.0.0.1:8080 -mode ws

# 启动http服务端
# 需要访问wxbot的sync-url接口进行注册
$ ./go-client -addr 0.0.0.0:8081 -mode http
# 进行http回调接口注册（不要忘记带上协议头）
$ curl -XPOST -d'{"url": "http://127.0.0.1:8081/callback", "timeout": 3000}' 127.0.0.1:8080/sync-url
```