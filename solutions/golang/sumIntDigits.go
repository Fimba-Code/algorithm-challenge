package main

import (
	"fmt"
	"unicode"
	"strconv"
)

// 1 - Receber o input do utilizador
// 2 - Armazenar o número numa variavel do tipo int "number"
// 3 - Se o numero for 0 vai retornar 0
// 4 - calcular o resto da divisõa por 10 do número recebido e somar o numero recebido dividido por 10 chamando a mesma função (recursividade)  ---> result = number%10 + sum(number/10)
// 5 - imprimir o resultado

func sum(number int) int {
	result := 0	// 2

	if number == 0 {	//3
		return 0
	}
	
	result = number%10 + sum(number/10)
	return result
}


func main() {
	var number string

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)	// 1

	if (isInt((number))) {
		i1, _ := strconv.Atoi(number)
		fmt.Printf("Digits of %d is equal to %d\n", i1, sum(i1))
	} else {
		fmt.Printf("The value entered is not a number\n", number)
	}
	
}

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}