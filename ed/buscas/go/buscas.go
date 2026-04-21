package main
import "fmt"

func avaliarStrings(strings []string, stringConsultas []string) []int {
    
    counts := make(map[string]int)

    for _, v := range strings {
        counts[v]++
    }

    resultados := make([]int, len(stringConsultas))

    for i, q := range stringConsultas {
        resultados[i] = counts[q]
    }

    return resultados
}

func main() {
    
    var qtd_string_principal int
    var qtd_string_buscas int
    
    fmt.Scan(&qtd_string_principal)
    stringEntrada := make([]string, qtd_string_principal)
    
    for i := 0 ; i < qtd_string_principal ; i++ {
        
        fmt.Scan(&stringEntrada[i])
    }

    fmt.Scan(&qtd_string_buscas)
    stringConsultas := make([]string, qtd_string_buscas)
    
    for i := 0 ; i < qtd_string_buscas ; i++ {
        
        fmt.Scan(&stringConsultas[i])
    }

    resultado := avaliarStrings(stringEntrada, stringConsultas)

    for i, v := range resultado {
        
        fmt.Print(v)
        
        if i < len(resultado) - 1 {
         
            fmt.Print(" ")
        }
    }
    fmt.Println()

    
    

    

}
