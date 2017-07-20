package main

import (
	"fmt"
	"net/http"

	"github.com/rodrigodealer/realtime/server"
)

func main() {
	err := http.ListenAndServe(":8080", server.Server())
	if err != nil {
		fmt.Println("Something is wrong : " + err.Error())
	}
}
