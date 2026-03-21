package main
import ("fmt"
        "math")

func main() {
   
    qtd := 0
    fmt.Scan(&qtd)

    jogadores := make([] int, qtd)
    vencedor := 0 
    
    for  i := 0 ; i < qtd ; i++ {

        var jogA, jogB int
        fmt.Scan(&jogA, &jogB)

        
        if jogA < 10 || jogB < 10 { 
            jogadores[i] = -1
            continue
        }

        jogadores[i] = int (math.Abs(float64( jogA - jogB ) ) )
    }
    comparable := 0
    maiorvalor := 101
    for i:=0 ; i<qtd ; i++ {

        if jogadores[i] < 0 {
            vencedor = i - 1
       
        } else {
            if maiorvalor > jogadores[i]{
                maiorvalor = jogadores[i] 
                vencedor = i 
            }


        }


        comparable += jogadores[i]
    }

    if comparable < 0 {
        fmt.Printf("sem ganhador\n")
    
    } else {

        fmt.Println(vencedor)
    }
}
