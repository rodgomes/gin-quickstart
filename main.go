package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rodgomes/go-quickstart/config"
	"github.com/rodgomes/go-quickstart/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
