package main

import "fmt"

func main() {
    lista := [3]int{10, 15, 20}
    soma := 0
    for _, num := range lista {
        soma += num
    }
    fmt.Println("A soma dos valores é:", soma)
}
