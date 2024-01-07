import express from 'express'
import axios from 'axios'
import WebSocket from 'ws'

class HttpServer {
  constructor({ host, port, callback }) {
    this.app = express()
    this.host = host || '127.0.0.1'
    this.port = port || 8081
    this.callback = callback || (() => {})

    this.initializeMiddlewares()
    this.initializeRoutes()

    this.listen()
  }

  initializeMiddlewares() {
    this.app.use(express.json())
    this.app.use(express.urlencoded({ extended: true }))
  }

  initializeRoutes() {
    this.app.post('/callback', (req, res) => {
      this.callback(req.body)

      res.json({
        code: 200,
        data: null,
        msg: 'success'
      })
    })
  }

  listen() {
    this.app.listen(this.port, () => {
      console.log(`=================================`)
      console.log(`🚀 App listening on the port ${this.port}`)
      console.log(`=================================`)
    })
  }

  async register(domain) {
    await axios.post(`${domain}/api/sync-url`, {
      url: `http://${this.host}:${this.port}/callback`,
      timeout: 3000
    })
  }

  getServer() {
    return this.app
  }

  getConfig() {
    return {
      port: this.port
    }
  }
}

class SocketClient {
  constructor(domain) {
    this.ws = new WebSocket(`${domain}/ws/msg`)
  }

  register(callback) {
    return new Promise((resolve, reject) => {
      this.ws.on('open', () => {
        console.log('Socket Connected to server')
        resolve()
      })
      this.ws.on('message', message => {
        message && callback(JSON.parse(message))
      })
      this.ws.on('error', msg => {
        console.log(`Socket Error: ${msg}`)
        reject(msg)
      })
      this.ws.on('close', () => {
        console.log('Socket Connection closed')
      })
    })
  }
}

class WxBot {
  constructor(domain) {
    this.domain = domain
  }

  async sendtextmsg(wxid, content) {
    return axios.post(`${this.domain}/api/sendtxtmsg`, {
      wxid,
      content
    })
  }

  httpRegister = callback => new HttpServer({ callback }).register(this.domain)

  socketRegister = async callback => {
    await new SocketClient(this.domain).register(callback)
  }

  async register(callback, type = 'socket') {
    switch (type) {
      case 'http':
        await this.httpRegister(callback)
        break
      case 'socket':
        await this.socketRegister(callback)
        break
    }
    return this
  }
}

const delay = () => new Promise(resolve => setTimeout(resolve, 3000))

const main = async () => {
  const wxid = '47331170911@chatroom'
  const domain = 'http://127.0.0.1:8080'

  const bot = new WxBot(domain)

  // 1. http callback register
  await bot.register(data => {
    console.log(`--------- http recv start ---------`)
    console.log(data)
    console.log(`--------- http recv end ---------`)
  }, 'http')
  // 2. socket callback register（default）
  await bot.register(data => {
    console.log(`--------- scoket recv start ---------`)
    console.log(data)
    console.log(`--------- scoket recv end ---------`)
  })

  await delay()

  // 3. send msg
  const response = await bot.sendtextmsg(wxid, 'hello world!')
  console.log(response.data)
}

await main()
