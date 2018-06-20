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
#### nginx 代理转发
* [参考 链接](https://wiki.swoole.com/wiki/page/326.html)
