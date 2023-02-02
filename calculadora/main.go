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
	scanner.Scan()
	operacion := scanner.Text()
	valores := strings.Split(operacion, "+")
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	fmt.Println(operador1 + operador2)
}
