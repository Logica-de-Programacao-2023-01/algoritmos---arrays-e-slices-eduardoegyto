package main

import "fmt"

func main() {
    var tamanho int
    fmt.Print("Digite o tamanho da lista: ")
    fmt.Scanln(&tamanho)

    lista := make([]int, tamanho)
    soma := 0

    for x := 0; x < tamanho; x++ {
        fmt.Printf("Digite o valor %d: ", x+1)
        fmt.Scanln(&lista[x])
        soma += lista[x]
    }

    fmt.Println("O Slice informado é:", lista)
    fmt.Println("A soma dos valores do Slice é:", soma)
}
