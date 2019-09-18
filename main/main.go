package main

import (
	"flag"
	"log"
	server "top.guitoubing/gotest"
	"top.guitoubing/gotest/db"
)

func main() {
	var addr string
	var mode string
	flag.StringVar(&addr, "addr", ":1323", "server listens at this addr")
	flag.StringVar(&mode, "mode", "dev", "set server mode")
	flag.Parse()

	err := db.InitializeGlobalDB("127.0.0.1", "gotest")
	if err != nil {
		log.Panic(err)
	}

	s := server.NewServer(addr)
	err = s.Init()
	if err != nil {
		log.Panic(err)
	}
	s.Start()
}