package main

import "fmt"

func main() {
	n := 5
	array := make([]int, n)

	fmt.Println("Digite os elementos do array:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&array[i])
	}

	ordem := true
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			ordem = false
			break
		}
	}

	if ordem {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
