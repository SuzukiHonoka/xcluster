package main

import (
	"flag"
	"xcluster/internal/launch"
)

func main() {
	server := flag.Bool("s", false, "server mode")
	client := flag.Bool("c", false, "client mode")
	flag.Parse()
	switch {
	case *server:
		launch.Server()
	case *client:
		launch.Client()
	}
}
