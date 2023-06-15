package main

import "fmt"

func main() {
    lista := make([]int, 0, 5)

    var x int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scanln(&x)

    presente := false
    for _, num := range lista {
        if num == x {
            presente = true
            break
        }
    }

    if !presente {
        lista = append(lista, x)
    }

    fmt.Println("A lista fica assim:", lista)
}
