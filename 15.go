package main

import "fmt"

func main() {
	
x := [10]float64{1.2, 3.4, 5.6, 7.8, 9.0, 10.1, 12.3, 14.5, 16.7, 18.9}

	var y []float64

	for _, num := range x {
		if num > 5 {
			slice = append(y, num)
		}
	}

	fmt.Println(y)
}
