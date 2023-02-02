package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escriba el primer operador")

	// Lee el imput del usuario
	scanner.Scan()

	// Convierte el imput a un string y lo guarda en la variable "operacion"
	operador1string := scanner.Text()

	fmt.Println("Ingrese el tipo de operacion")

	scanner.Scan()

	operador := scanner.Text()

	fmt.Println("Ingrese el segundo operador")

	scanner.Scan()

	operador2string := scanner.Text()

	//fmt.Println(operador1string, operador, operador2string)

	operador1, err1 := strconv.Atoi(operador1string)

	operador2, err2 := strconv.Atoi(operador2string)

	if err1 == nil && err2 == nil {
		switch operador {
		case "+":
			fmt.Println(operador1, operador, operador2, "=", operador1+operador2)
		case "-":
			fmt.Println(operador1, operador, operador2, "=", operador1-operador2)
		case "*":
			fmt.Println(operador1, operador, operador2, "=", operador1*operador2)
		case "/":
			fmt.Println(operador1, operador, operador2, "=", operador1/operador2)
		default:
			fmt.Println("Operador invalido")
		}
	} else {
		fmt.Println("Se esperaban numeros enteros")
	}
}
