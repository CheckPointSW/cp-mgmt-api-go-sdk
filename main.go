package main

import (
	"fmt"
	"github.com/CheckPointSW/cp-mgmt-api-go-sdk/Examples"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("cp-mgmt-api-go-sdk main: Operation not found. Optional operations are rule/discard/add_host/show_hosts/dup_ip/auto_publish")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "discard":
		Examples.DiscardSessions()
	case "rule":
		Examples.AddAccessRule()
	case "add_host":
		Examples.AddHost()
	case "show_hosts":
		Examples.ShowHosts()
	case "dup_ip":
		Examples.DupIp()
	case "auto_publish":
		Examples.AutoPublish()
	default:
		fmt.Println("cp-mgmt-api-go-sdk main: Operation not supported. Optional operations are rule/discard/add_host/show_hosts/dup_ip/auto_publish")
	}
}
