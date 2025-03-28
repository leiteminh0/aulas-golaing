package main

import "fmt"

func main() {
	var num1, num2  float64
	
	fmt.Print("Digite o primeiro numero:")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo numero:")
	fmt.Scan(&num2)

	sum:= num1 + num2
	diff:= num1 - num2
	prod:= num1 * num2

	fmt.Println("\nResultados:")
	fmt.Printf("Soma: %.2f\n", sum)
	fmt.Printf("Subtração: %.2f\n", diff)
	fmt.Printf("Multiplicação: %.2f\n", prod)

	f num2 != 0 {
		quot := num1 / num2
		fmt.Printf("Divisão: %.2f\n", quot)
	} else {
		fmt.Println("Divisão por zero não é permitida.")
	}
	if num2 != 0 {
		mod := int(num1) % int(num2)
		fmt.Printf("Módulo: %d\n", mod)
	} else {
		fmt.Println("Módulo por zero não é permitido.")
	}
}