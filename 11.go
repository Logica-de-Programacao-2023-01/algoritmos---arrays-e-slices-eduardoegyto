package main

import "fmt"

func main() {
	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var x, y int

	fmt.Print("Digite o índice da linha: ")
	fmt.Scanln(&x)
	fmt.Print("Digite o índice da coluna: ")
	fmt.Scanln(&y)

	fmt.Printf("Valor armazenado na posição [%d][%d]: %d\n", x, y, array[x][y])
}
