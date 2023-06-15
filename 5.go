package main

import "fmt"

func main() {
    var matriz [3][2]int

    for i := 0; i < 3; i++ {
        for x := 0; x < 2; x++ {
            fmt.Printf("Digite o valor para a posição [%d][%d]: ", i, x)
            fmt.Scanln(&matriz[i][x])
        }
    }

    fmt.Println("A matriz informada é:")
    for _, y := range matriz {
        fmt.Println(y)
    }
}
