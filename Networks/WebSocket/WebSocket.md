# WebSocket
***
> 原理：
- WebSocket 同 HTTP 一样也是应用层的协议，但是它是一种双向通信协议，是建立在 TCP 之上的。

> WebSocket和Http的关系和异同点
- 每个WebSocket连接都始于一个HTTP请求。 具体来说，WebSocket协议在第一次握手连接时，通过HTTP协议在传送WebSocket支持的版本号，协议的字版本号，原始地址，主机地址等等一些列字段给服务器端
- Upgrade首部，用来把当前的HTTP请求升级到WebSocket协议，这是HTTP协议本身的内容，是为了扩展支持其他的通讯协议。 如果服务器支持新的协议，则必须返回101
- 一个WebSocket连接是在客户端与服务器之间HTTP协议的初始握手阶段将其升级到Web Socket协议来建立的，其底层仍是TCP/IP连接
---
- 相同点：
  - （1）都是建立在TCP之上，通过TCP协议来传输数据。
  - （2）都是可靠性传输协议。
  - （3）都是应用层协议。
- 不同点：
  - （1）WebSocket支持持久连接，HTTP不支持持久连接。
  - （2）WebSocket是双向通信协议，HTTP是单向协议，只能由客户端发起，做不到服务器主动向客户端推送信息。

> WebSocket和Socket
- Socket 其实并不是一个协议，而是为了方便使用 TCP 或 UDP 而抽象出来的一层，是位于应用层和传输控制层之间的一组接口。 Socket本身并不是一个协议，它工作在OSI模型会话层，是一个套接字，TCP/IP网络的API，是为了方便大家直接使用。
- 更底层协议而存在的一个抽象层。Socket其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在Socket接口后面，对用户来说，一组简单的接口就是全部，让Socket去组织数据，以符合指定的协议。
而WebSocket则是一个典型的应用层协议。

> WebSocket  HTTP和TCP/IP
- WebSocket和HTTP一样，都是建立在TCP之上，通过TCP来传输数据。

> Socket和TCP/IP
- Socket是对TCP/IP协议的封装，像创建Socket连接时，可以指定使用的传输层协议，Socket可以支持不同的传输层协议(TCP或UDP),当使用TCP协议进行连接时，该Socket连接就是一个TCP连接。
