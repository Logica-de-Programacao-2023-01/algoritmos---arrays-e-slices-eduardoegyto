package main

import (
    "fmt"
    "math"
)

func main() {
    lista := []int{8, 2, 5, 10, 1, 6, 9, 3, 7, 4}

    min := math.MaxInt64
    max := math.MinInt64

    for _, num := range lista {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    fmt.Println("O valor mínimo na lista é:", min)
    fmt.Println("O valor máximo na lista é:", max)
}
