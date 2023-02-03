package main

import (
	"fmt"
	"net/http"
	"time"
)

func serverChecker(server string, c chan string) {
	_, err := http.Get(server)
	if err != nil {
		// Sin usar GoRoutines
		fmt.Println(server, "no envio ninguna respuesta")

		// Usando GoRoutines y Channels
		c <- server + " no envio ninguna respuesta con goroutine"
	} else {
		// Sin usar GoRoutines
		fmt.Println(server, "esta funcionando correctamente")
		// Usando GoRoutines y Channels
		c <- server + " esta funcionando correctamente con goroutine"
	}
}

func main() {
	tiempoInicio := time.Now()
	canal := make(chan string)

	servers := []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://google.com",
	}

	for _, server := range servers {
		go serverChecker(server, canal)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-canal)
	}

	tiempoFinal := time.Since(tiempoInicio)
	fmt.Printf("el programa tardo en ejecutarse\n%s", tiempoFinal)
}

/*
	Cuando ejecute el programa sin goroutines tardo en ejecutarse 2,92s
	mientras que usando goroutines tardo tan solo 1,38s
*/
