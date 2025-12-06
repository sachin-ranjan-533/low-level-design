package main

import (
	"fmt"
	"proxy-pattern/server"
)

func main() {
	proxy := server.NewProxyServer()
	fmt.Println(proxy.ServeRequest("admin"))
	fmt.Println(proxy.ServeRequest("guest"))
}
