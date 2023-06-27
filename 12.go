package main

import "fmt"

func main() {
	   array := [5]int{10, 15, 20, 25, 30}

	var x []int
  for _, num := range array {
		if num%3 == 0 {
			slice = append(x, num)
		}
	}
	fmt.Println(x)
}
