package main

import "fmt"

func main() {
    lista := [6]float64{2.5, 4.7, 6.2, 8.9, 10.1, 12.3}

    var x float64
    fmt.Print("Digite um número: ")
    fmt.Scanln(&x)

    for i := 0; i < len(lista); i++ {
        lista[i] += x
    }

    fmt.Println("A lista fica assim:", lista)
}
