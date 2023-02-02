package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada, operador string) (int, error) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1, err1 := parsear(entradaLimpia[0])
	operador2, err2 := parsear(entradaLimpia[1])
	if err1 == nil && err2 == nil {
		switch operador {
		case "+":
			return operador1 + operador2, nil
		case "-":
			return operador1 - operador2, nil
		case "*":
			return operador1 * operador2, nil
		case "/":
			return operador1 / operador2, nil
		default:
			return 0, fmt.Errorf("operador invalido")
		}
	} else {
		return 0, fmt.Errorf("se esperaba un entero")
	}
}

func parsear(entrada string) (int, error) {
	operador, err := strconv.Atoi(entrada)
	return operador, err
}

func escanearEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Ingrese los valores de la forma: e.g. 2+2 , 2*3 , etc.")
	input := escanearEntrada()
	fmt.Println("Ingrese el tipo de operacion de la forma: e.g. + , - , * , /")
	operador := escanearEntrada()
	c := calc{}
	result, err := c.operate(input, operador)
	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
