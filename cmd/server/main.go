package main

import (
	"backend_todoapp/internal/server"
	"log"
)

func main() {
	e := server.NewServer()
	log.Fatal(e.Start(":8080"))
}
