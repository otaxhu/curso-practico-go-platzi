package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escriba sus 2 operadores de la forma: 2+2")

	// Lee el imput del usuario
	scanner.Scan()

	// Convierte el imput a un string y lo guarda en la variable "operacion"
	operacion := scanner.Text()

	// Separa los numeros cada vez que haya un "+" y los guarda en el array "valores[]"
	valores := strings.Split(operacion, "+")

	// Convierte "valores[0]" a un numero entero y lo guarda en la variable "operador1"
	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println("El primer operador no es un entero")
	}

	// Convierte "valores[1]" a un numero entero y lo guarda en la variable "operador2"
	operador2, err2 := strconv.Atoi(valores[1])
	if err2 != nil {
		fmt.Println("El segundo operador no es un entero")
	}

	// Si no hay ningun error en las conversiones imprime la suma de los dos operadores
	if err1 == nil && err2 == nil {
		fmt.Println(operador1 + operador2)
	}
}
