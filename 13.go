package main

import "fmt"
 
func main() {
	x := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	x[0] += num
	x[len(array)-1] += num
  
	fmt.Println(x)
}
