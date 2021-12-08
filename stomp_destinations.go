package main

// 提供给客户端访问的地址
var AcceptInterface map[string]string = map[string]string{
	"AskWhatamidoing": "/whatamidoing",
	"AskWhoami":       "/whoami",
	"AskWhereami":     "/whereami",
	// or more ...
}

// 客户端信息返回的订阅地址
var SubscribeMap map[string]string = map[string]string{
	"Whatamidoing": "/broadcast/whatamidoing",
	"Whoami":       "/broadcast/whoami",
	"Whereami":     "/broadcast/whereami",
	// or more
}
