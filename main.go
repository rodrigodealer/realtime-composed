package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rodrigodealer/realtime/server"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", server.Server())
	if err != nil {
		log.Panic("Something is wrong : " + err.Error())
	}
}
