# goChat
golang + beego + vue 直播聊天互动

###安装要求
1. 安装beego
2. npm 版本大于3.0.0，node 版本大于4.0.0


###项目实现目标
1. 实现直播用户发言互动
2. 使用nginx 代理转发websocket 请求，主要使用nginx 负载均衡，部署到多台服务器
3. 采用前后端分离方式进行开发，加深对项目合作过程中，对项目理解
4. 程序安装及相关配置，请分别参考backend、frontend 下read.me

* backend 是后端beego 项目
* frontend 是前端vue 项目
* chat.2345.com.conf 是nginx 配置文件

#### 这次采用前后端分离方式单独实现
* 后端采用beego 框架 ，使用 websocket 长链通信
* 前端采用vue.js + vuex + vue router + ele UI + webpack 实现 

#### 参考链接
*  [golang websocket example](:https://github.com/gorilla/websocket/tree/master/examples/chat)
* [nginx 代理转发websocket](:https://wiki.swoole.com/wiki/page/326.html)

#### 后端设计结构

## Server

The server application defines two types, `Client` and `Hub`. The server
creates an instance of the `Client` type for each websocket connection. A
`Client` acts as an intermediary between the websocket connection and a single
instance of the `Hub` type. The `Hub` maintains a set of registered clients and
broadcasts messages to the clients.

The application runs one goroutine for the `Hub` and two goroutines for each
`Client`. The goroutines communicate with each other using channels. The `Hub`
has channels for registering clients, unregistering clients and broadcasting
messages. A `Client` has a buffered channel of outbound messages. One of the
client's goroutines reads messages from this channel and writes the messages to
the websocket. The other client goroutine reads messages from the websocket and
sends them to the hub.

### Hub 

The code for the `Hub` type is in
[hub.go](https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go). 
The application's `main` function starts the hub's `run` method as a goroutine.
Clients send requests to the hub using the `register`, `unregister` and
`broadcast` channels.

The hub registers clients by adding the client pointer as a key in the
`clients` map. The map value is always true.

The unregister code is a little more complicated. In addition to deleting the
client pointer from the `clients` map, the hub closes the clients's `send`
channel to signal the client that no more messages will be sent to the client.

The hub handles messages by looping over the registered clients and sending the
message to the client's `send` channel. If the client's `send` buffer is full,
then the hub assumes that the client is dead or stuck. In this case, the hub
unregisters the client and closes the websocket.

### Client

The code for the `Client` type is in [client.go](https://github.com/gorilla/websocket/blob/master/examples/chat/client.go).

The `serveWs` function is registered by the application's `main` function as
an HTTP handler. The handler upgrades the HTTP connection to the WebSocket
protocol, creates a client, registers the client with the hub and schedules the
client to be unregistered using a defer statement.

Next, the HTTP handler starts the client's `writePump` method as a goroutine.
This method transfers messages from the client's send channel to the websocket
connection. The writer method exits when the channel is closed by the hub or
there's an error writing to the websocket connection.

Finally, the HTTP handler calls the client's `readPump` method. This method
transfers inbound messages from the websocket to the hub.

WebSocket connections [support one concurrent reader and one concurrent
writer](https://godoc.org/github.com/gorilla/websocket#hdr-Concurrency). The
application ensures that these concurrency requirements are met by executing
all reads from the `readPump` goroutine and all writes from the `writePump`
goroutine.

To improve efficiency under high load, the `writePump` function coalesces
pending chat messages in the `send` channel to a single WebSocket message. This
reduces the number of system calls and the amount of data sent over the
network.