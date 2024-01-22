import WebSocket from 'ws';

// 替换为你的 WebSocket 服务端地址
const wsUrl = 'wss://your-websocket-server.com/path';

// 创建 WebSocket 客户端实例
const ws = new WebSocket(wsUrl);

// 监听 WebSocket 连接打开事件
ws.on('open', function open() {
  console.log('Connected to the server.');
});

// 监听接收消息事件
ws.on('message', function message(data) {
  console.log('Received message from server:', data.toString());
});

// 监听错误事件
ws.on('error', function error(err) {
  console.error('WebSocket encountered error:', err);
});

// 监听连接关闭事件
ws.on('close', function close(code, reason) {
  console.log(`Connection closed. Code: ${code}, Reason: ${reason}`);
});
