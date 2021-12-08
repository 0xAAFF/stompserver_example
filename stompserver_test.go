package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println(StompServerInstance.IStompManager)
	fmt.Println(StompServerInstance.IStompManager.NewMessageId())

}
