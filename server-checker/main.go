package main

import (
	"fmt"
	"net/http"
)

func serverChecker(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "no envio ninguna respuesta")
	} else {
		fmt.Println(server, "esta funcionando correctamente")
	}
}

func main() {
	servers := []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://google.com",
	}
	for _, server := range servers {
		serverChecker(server)
	}
}
