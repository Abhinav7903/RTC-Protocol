package main

import (
	"flag"
	"rtc/server"
)

func main() {
	var envType string
	flag.StringVar(&envType, "env", "dev", "Environment type: production, staging, dev, prod")
	flag.Parse()
	server.Run(&envType)
}
