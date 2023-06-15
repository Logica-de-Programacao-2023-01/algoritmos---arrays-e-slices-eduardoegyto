package main

import "fmt"

func main() {
    lista := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

    var x int
    fmt.Print("Digite um valor: ")
    fmt.Scanln(&x)

    encontrado := false
    for _, y := range lista {
        if y == x {
            encontrado = true
            break
        }
    }

    if encontrado {
        fmt.Println("O valor", x, "foi encontrado na lista.")
    } else {
        fmt.Println("O valor", x, "não foi encontrado na lista.")
    }
}
