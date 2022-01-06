package main

import (
	"flag"
	"fmt"
	"os"

	"sdtp.io/pkg/common"
	"sdtp.io/scripts"
)

var protoType string
var client bool

func main() {
	if client {
		fmt.Println("start client..")
		scripts.Client()
	} else {
		fmt.Println("start server..")
		server := common.NewNonBlockingServer()
		server.Start("tcp://127.0.0.1:9999", protoType)
	}
	os.Exit(0)
}

func init() {
	flag.StringVar(&protoType, "protocol", "v1", "select protocol")
	flag.BoolVar(&client, "client", false, "start client")
	flag.Parse()
}
