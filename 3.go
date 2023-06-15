package main

import "fmt"

func main() {
    lista := [4]float64{10, 5.7, 3, 8}
    produto := 1.0
    for _, x := range lista {
        produto *= x
    }
    fmt.Println("O produto dos valores é:", produto)
}
