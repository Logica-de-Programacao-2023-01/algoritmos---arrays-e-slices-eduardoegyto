package main

import "fmt"

func main() {
    lista := []string{"banana", "maçã", "abacaxi", "laranja", "pêra", "uva", "manga", "maracujá"}

    var x string
    fmt.Print("Digite um valor: ")
    fmt.Scanln(&x)

    nlista := make([]string, 0)
    for _, item := range lista {
        if item != x {
            nlista = append(nlista, item)
        }
    }

    fmt.Println("A lista fica assim::", nlista)
}
