package main

import (
	"flag"
	"github.com/google/uuid"
	"log"
	"xcluster/internal/launcher"
)

func init() {
	// use rand pool by default for better performance
	uuid.EnableRandPool()
}

func main() {
	server := flag.Bool("s", false, "server mode")
	client := flag.Bool("c", false, "client mode")
	flag.Parse()
	switch {
	case *server:
		serverLauncher := launcher.NewServer()
		go serverLauncher.Launch()
		serverLauncher.ListenForShutdown()
	case *client:
		launcher.Client()
	default:
		log.Fatalln("You must need either specify server or client to run the xcluster.")
	}
}
