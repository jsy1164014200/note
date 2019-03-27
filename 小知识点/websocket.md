# websocket

> [详解websocket](http://www.52im.net/thread-331-1-1.html)

1. 通讯机制建立在http之上,首先通过 http发送带有 header updatewebsocket 的头部,返回再使用http返回一个 是否支持websocket 的信息
2. 这样就建立起来了 websocket连接.
3. 这个连接只要没有显示的调用close方法,那么他就不会关闭

