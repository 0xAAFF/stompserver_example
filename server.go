package main

import (
	"fmt"
	"net/http"

	"github.com/0xAAFF/stompserver/tools"
	"golang.org/x/net/websocket"
)

/*
 * "/*"  		Web服务
 *
 * "/stomp"		Stomp服务
 */

func Server() {

	serverPort := 3459

	http.HandleFunc("/", HttpVueServer)                                        // Web服务
	http.Handle("/stomp", websocket.Handler(StompServerInstance.NewStompUnit)) // Stomp服务

	fmt.Println(fmt.Sprint("Web Server  : http://127.0.0.1:", serverPort, "/"))
	filepath := tools.GetCurrentPathNoError() + ROOT_PATH
	fmt.Println("Web source :" + filepath)
	fmt.Println("Vue Project:(if u have,then)")
	fmt.Println("1 npm run build \n2 open  ./dist\n3 copy all files to " + filepath + "\n4 refresh Web Page.")

	fmt.Println(fmt.Sprint("Stomp Server: ws://localhost:", serverPort, "/stomp"))

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", serverPort), nil) //only local
	if err != nil {
		fmt.Println("init Serve: " + err.Error())
		panic("initServer: " + err.Error())
	}
}
