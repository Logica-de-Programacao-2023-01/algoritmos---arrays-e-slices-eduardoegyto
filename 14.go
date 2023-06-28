package main

import "fmt"

func main() {
	
  x := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var num1, num2 int 

	
	fmt.Print("Digite o índice do primeiro elemento: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o índice do segundo elemento: ")
	fmt.Scanln(&inum2)

	x[num1], x[num2] = x[num2], x[num1]

	fmt.Println(x)
}
