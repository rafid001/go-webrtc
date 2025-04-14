package main

import (
	"log"

	"github.com/rafid001/videochat/internal/server"
)


func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}