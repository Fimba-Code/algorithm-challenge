package main

import (
	"fmt"
	"unicode"
	"strconv"
)

// 1 - Receber o input do utilizador
// 2 - Armazenar o número numa variavel
// 3 - Se o numero for diferente de inteiro vai retornar 'The value entered is not a number'
// 4 - Iterar até o número que foi recebido
// 5 - executar a função fibonacci dentro da iteração
// 6 - como a sequência de fibonacci consiste na soma dos seus dois números anteriores 
//		então declaramos duas variaveis: first e second para receber as primeiras sequencia de fibonnaci e utilizá-las no próximo for
// 7 - Iteramos até o número que recebemos
// 8 - dentro da iteração declaramos mais uma variável chamada tempo que vai servir para trocar os valores de first e second
// 9 - armazenamos o valor da variavel second dentro da variavel first
// 10 - armazemos a soma das variaveis tempo e first dentro da variavel second para seguir a sequencia de fibonnaci (soma dos seus dois números anteriores)
// 11 - depois da iteração imprimimos a variável first que é a que vai ter sempre o resultado da variavel second (soma dos seus dois números anteriores)


func fibonacci(n int) int {
    first := 0
    second := 1
    for i := 0; i < n; i++ {
        tempo := first
        first = second
        second = tempo + first
    }
    return first
}

func main() {

    var number string

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)	// 1

	if (isInt((number))) {
		i1, _ := strconv.Atoi(number)
        for n := 0; n < i1; n++ {
            result := fibonacci(n)
            fmt.Printf("%v\n", result)
        }
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