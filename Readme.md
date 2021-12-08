# 示例
this is https://github.com/0xAAFF/stompserver example

in this case , we have WebServer(for Vue Project)  and Stomp Server

## stompserver.go How to use Stomp Server 如何使用Stomp Server
follw stompserver.go.init() :
    you can follw steps to use it .
    1 regist root path(message direction)
    2 regist interface path(client can subscribe)
    3 create a server
    4 link message reflex handle(implements your interface)

具体查阅 stompserver.go文件中的 init()这个函数,这里已经包含了相关步骤
步骤如下:
    1 注册根地址(根地址决定了数据的流向,是否群发,单发,组发)
    2 注册接口地址(提供给客户端订阅接口)
    3 创建一个server对象(for "net/http") 
        //http.Handle("/stomp", websocket.Handler(StompServerInstance.NewStompUnit)) // Stomp服务
    4 为对象关联处理消息的函数(所有接口都在这里实现)
        消息接收和消息返回都在Reflex函数中处理

    // 其实这个处理方式与Java的broker的类似
    // 我这里并没有完全实现stomp的协议,后面又时间加吧



## webserver_vue.go
webserver_vue is a server,to work on Vue Project build files
是用来解析Vue build的项目的,是一个Vue Server

in webserver_vue.go:157 func onReflex(responseW http.ResponseWriter, request *http.Request) {
    you can add some http Get interface.

这个模块是给Vue准备的 Vue项目生成后,dist目录下的文件放在./data/目录下,使用web访问应该就可以解析到了


其实我们的项目是本地用的,前后端分离.所以一个端口干好几件事情.
http总是使用定时器拉取数据也存在弊端,消息不能快速同步



### steps

1 
go mod tidy

2
go build -o ./out/ser.exe

3 
cd clientweb/

4 
npm i

5
// (just check if may build)
npm run serve 

6
npm run build

7

copy ./clientweb/dist/ all files to ./out/data

8
cd ..
./out/ser.exe

9 
open brower : 
http://127.0.0.1:3459/
http://127.0.0.1:3459/?i=send

10 
F12-open develop tools-console