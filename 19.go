package main

import "fmt"

func main() {
	x := 5
	a1 := make([]int, x)
	a2 := make([]int, x)

	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < x; i++ {
		fmt.Scanln(&a1[i])
	}

	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < x; i++ {
		fmt.Scanln(&ay2[i])
	}

	a3 := make([]int, x)
	for i := 0; i < n; i++ {
		array3[i] = ay1[i] + ay2[i]
	}

	fmt.Println("Resultado da soma:")
	fmt.Println(a3)
}
