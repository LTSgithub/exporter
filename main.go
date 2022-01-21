package main

import (
	"github.com/lits01/xiaozhan/server"
	"log"
)

func main() {
	server, err := server.InitServer()
	if err != nil {
		log.Fatalln(err)
	}

	server.Run()
}
