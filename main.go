package main

import (
	"fmt"
)

func main() {
	fmt.Println("Web/Stomp server")

	go Server()

	// just for wait enter
	waitKey := ""
	fmt.Scanln(&waitKey)

}
