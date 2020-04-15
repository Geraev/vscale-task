package main

import (
	"flag"
	"fmt"
	"log"

	"vscale-task/cmd/httpserver"
	"vscale-task/cmd/manager"
	"vscale-task/cmd/providers/vscale"
	"vscale-task/cmd/storage"
)

var (
	httpAddr string
	token string
)

func init() {
	flag.StringVar(&httpAddr, "port", "8081", "The HTTP port to bind to")
	flag.StringVar(&token, "token", "", "Token for VScale API")
	flag.Parse()
}

func main() {

	flag.PrintDefaults()
	fmt.Println(httpAddr, "===", token, "<<<")

	APIManager := manager.NewAPIManager(
		vscale.NewClient(token),
		storage.NewStorage(),
		)

	http := httpserver.NewServer(httpAddr, APIManager)
	if err := http.Run(); err != nil {
		log.Fatalln(err)
	}

}
