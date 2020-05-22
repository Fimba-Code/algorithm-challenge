package main

import (
	"fmt"
)

// 1 - Receber o input do utilizador
// 2 - Armazenar o que foi digitado numa variavel numa variável do tipo String
// 3 - Criamos uma função que recebe uma string e retorna um boleano
// 4 - Iteraramos até a divisão do tamanho do valor que foi recebido dividido por dois
// 5 - dentro do for fizemos a condição para verificar se cada posição do nosso contador(i) do valor recebido for diferente do
//		do inverso que é o valor na posisão da quantidade do valor que recebemos menos o nosso contador menos um

func IsPalindromic(value string) bool {
    for i := 0; i < len(value)/2; i++ {
        if value[i] != value[len(value)-i-1] {
            return false
        }
    }
    return true
}


func main() {

    var input string

	fmt.Print("Enter a number or word to see if is palindrome or not: ")
	fmt.Scanln(&input)	// 1

    fmt.Printf("IsPalindromic ??\nA: %v",IsPalindromic(input))
}
