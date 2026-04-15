package main
import (

    "fmt"
)
func main() {

    var qtd_pessoas_inicial int
    fmt.Scan(&qtd_pessoas_inicial)

    map_pessoas_inicial := make(map[int] int, qtd_pessoas_inicial)
    slice_chaves := make([]int, 0, len(map_pessoas_inicial))
    
    var pessoas int
    for i := 0 ; i < qtd_pessoas_inicial ; i++ {
        fmt.Scan(&pessoas)
        
        map_pessoas_inicial[pessoas] = pessoas
        slice_chaves = append(slice_chaves, pessoas)

        
        
    }
    
    var qtd_pessoas_sairam int 
    fmt.Scan(&qtd_pessoas_sairam)
    
    for i := 0 ; i < qtd_pessoas_sairam ; i++ {
        var pessoas_sairam int
        fmt.Scan(&pessoas_sairam)
        
        if _, ok := map_pessoas_inicial[pessoas_sairam]; ok{ 
            
            map_pessoas_inicial[pessoas_sairam] = 0 
            
        }
    }
    
    for _, pessoas := range slice_chaves {

        if map_pessoas_inicial[pessoas] != 0 {

            fmt.Printf("%d ", pessoas)
        }
    }
    fmt.Println()
}