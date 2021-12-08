package main

import (
	"github.com/0xAAFF/stompserver"
	"golang.org/x/net/websocket"
)

// #region 区域折叠
type StompServer struct {
	IStompManager      stompserver.StompManager // 所有管理相关都放入此处
	serviceStopChannel chan int                 // 服务停止信道
	ReflexHandle       func(sourceStompMessage *stompserver.Frame, unit *stompserver.StompUnit)
}

var StompServerInstance = &StompServer{
	IStompManager:      *stompserver.InstancesStompManager,
	serviceStopChannel: make(chan int),
}

// 设置Stomp服务的反射函数
func (stompServer *StompServer) SetReflex(reflex func(sourceStompMessage *stompserver.Frame, unit *stompserver.StompUnit)) {
	stompServer.ReflexHandle = reflex
}

// 新建立的stomp连接
func (stompServer *StompServer) NewStompUnit(ws *websocket.Conn) {
	stompUnit := stompserver.NewStompUnit(ws, &StompServerInstance.IStompManager, stompServer.ReflexHandle)
	stompUnit.Run()
}

// #endregion

func init() {

	// step 1
	/* 注册根地址
	* 根地址分为三种:群发(同一个订阅地址的消息将同步到每个订阅的客户端),单发(谁发送,谁接受),组发(成功验证的客户将接收到订阅地址的消息)
	* 只要路径与根地址匹配,则将采用约定的发送方式同步信息
	 */
	// regist root path 注册根地址
	StompServerInstance.IStompManager.RegistDestinationToBroadcastAll("/topic", "/toAll/", "/broadcast")
	StompServerInstance.IStompManager.RegistDestinationToApplication("/application", "/applications/")
	StompServerInstance.IStompManager.RegistDestinationToUser("/user")

	// step 2
	// 注册可订阅地址
	// regist interface path

	StompServerInstance.IStompManager.AddSubscribeDestination("/topic/cmd")
	StompServerInstance.IStompManager.AddSubscribeDestination("/application/control")
	StompServerInstance.IStompManager.AddSubscribeDestination("/application/cmd")

	// or

	// Stomp的订阅地址
	vlist := make([]string, 0)
	for _, v := range SubscribeMap {
		vlist = append(vlist, v)
	}
	// 初始化Stomp订阅地址
	StompServerInstance.IStompManager.AddSubscribeDestination(vlist...)

	// step 3
	// When stomp message coming , ReflexHandle hanld it
	StompServerInstance.SetReflex(Reflex)

}
