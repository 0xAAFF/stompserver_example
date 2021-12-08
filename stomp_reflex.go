package main

import (
	"github.com/0xAAFF/stompserver"
)

// 客户端针对不同的地址,做出对应的反应 reflex_stomp.go
func Reflex(sourceStompMessage *stompserver.Frame, unit *stompserver.StompUnit) {
	/*
	* 如下 是示例代码 供测试(需要配合相应的Web项目)
	 */

	// 此行代码勿删,防止连接未建立,客户端非法提交Send数据包
	if !unit.IsConnected {
		return
	}
	switch sourceStompMessage.Header.Get(stompserver.StompHeaders.Destination) {
	case AcceptInterface["AskWhatamidoing"]:
		{
			go Whatamidoing(sourceStompMessage, unit)
			break
		}
	case AcceptInterface["AskWhereami"]:
		{
			// go Whereami(sourceStompMessage, unit)
			break
		}
	case AcceptInterface["AskWhoami"]:
		{
			//go Whoami(sourceStompMessage, unit)
			break
		}
	case "/application/control":
		{

			messageFrame, errtxt := stompserver.NewMessageFrame("/application/control", StompServerInstance.IStompManager.NewMessageId(), "")
			if errtxt != "" {
				messageFrame = stompserver.NewErrorFrame("AskWhatamidoing Error", errtxt, sourceStompMessage)
			} else {
				var jsonText = `This Message is only back to you`
				messageFrame.SetBody(jsonText)
			}
			unit.SendStompMessage(messageFrame) // to one
			break
		}
	case "/application/cmd":
		{
			messageFrame, errtxt := stompserver.NewMessageFrame("/topic/cmd", StompServerInstance.IStompManager.NewMessageId(), "")
			if errtxt != "" {
				messageFrame = stompserver.NewErrorFrame("AskWhatamidoing Error", errtxt, sourceStompMessage)
			} else {
				var jsonText = `This Message is to All . orgin message.body is :` + string(sourceStompMessage.Body())
				messageFrame.SetBody(jsonText)
			}
			StompServerInstance.IStompManager.Publish(messageFrame) // to all
		}
		// ... or more
	}

}

// reflex_stomp.go
func Whatamidoing(sourceStompMessage *stompserver.Frame, unit *stompserver.StompUnit) {

	messageFrame, errtxt := stompserver.NewMessageFrame(SubscribeMap["Whatamidoing"], StompServerInstance.IStompManager.NewMessageId(), "")
	if errtxt != "" {
		messageFrame = stompserver.NewErrorFrame("AskWhatamidoing Error", errtxt, sourceStompMessage)
	} else {
		var jsonText = `{"name:"wstomp"}` // Whatamidoing.Serialize()
		messageFrame.SetBody(jsonText)
	}
	unit.SendStompMessage(messageFrame)
}
