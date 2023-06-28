package main

import "fmt"

func main() {
	x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var y []int

	for _, num := range x {
		if num%2 == 0 {
			y = append(y, num)
		}
	}

	fmt.Println(y)
}
