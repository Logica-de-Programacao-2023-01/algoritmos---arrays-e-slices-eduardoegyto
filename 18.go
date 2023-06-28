package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&x)

	y := make([]int, 0, x)

	for i := 2; len(y) < x; i++ {
		if primo(i) {
			y = append(y, i)
		}
	}

	fmt.Println(y)
}

func primo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
