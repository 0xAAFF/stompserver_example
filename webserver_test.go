package main

import (
	"fmt"
	"testing"
)

func TestWebServer(t *testing.T) {
	// fmt.Println(IsAllowAccess("chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/css/chunk-vendors.377fcb5a.js.css.map"))
	// fmt.Println(IsAllowAccess("/stataic/css/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/js/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/Static/js/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/jS/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/jss/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/jss/chunk-vendors.377fcb5a.js.map/../static/jss/chunk-vendors.377fcb5a.js.map/../static/jss/chunk-vendors.377fcb5a.js.map"))
	// fmt.Println(IsAllowAccess("/static/jss/ch5a/../static/jss/chuna.js.map"))

	fmt.Println(IsAllowAccess("/img/jschunk-vendors.377fcb5a.js.map"))
	fmt.Println(IsAllowAccess("/img/3.png"))
	fmt.Println(IsAllowAccess("/static/img/static/jschunk-vendors.377fcb5a.js.map"))
	fmt.Println(IsAllowAccess("/static/js/chunk-vendors.377fcb5a.js.map"))
	fmt.Println(IsAllowAccess("/images/../static/jschunk-vendors.377fcb5a.js.map"))
	fmt.Println(IsAllowAccess("/fonts/jschunk-vendors.377fcb5a.js.map"))
	fmt.Println(IsAllowAccess("/static/fonts//static/.377fcb5a.js."))
	fmt.Println(IsAllowAccess("/static/jschunk-vendors.377fcb5a.js.map"))

}
