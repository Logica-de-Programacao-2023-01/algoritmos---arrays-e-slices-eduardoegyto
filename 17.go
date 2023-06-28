package main

import "fmt"

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for i, num := range x {
		if i%2 == 0 {
			soma += num
		}
	}

	fmt.Println("Soma dos elementos nas posições pares:", soma)
}
