package main
import "fmt"
func main() {

    var qtdgomos int
    direction := ""
    var px, py int 
    
    fmt.Scan(&qtdgomos, &direction)
    tamvetor := 2 * qtdgomos

    before := make([]int, tamvetor)
    after := make([]int, tamvetor)

    for i := 0 ; i < tamvetor ; i += 2 {

        fmt.Scan(&px, &py)
        before[i] = px
        before[i + 1] = py
        
    }

    if qtdgomos > 1 {

        for i := 0 ; i < tamvetor - 2; i += 2 {
               

            after[i + 2] = before[i]
            after[i + 3] = before[i + 1]
            
    
        }

        switch direction {
        case "L":
            after[0] = before[0] - 1
            after[1] = before[1]

        case "R":
            after[0] = before[0] + 1
            after[1] = before[1]

        case "U":
            after[0] = before[0]
            after[1] = before[1] - 1

        case "D":
            after[0] = before[0]
            after[1] = before[1] + 1
        }

        for i := 0 ; i < tamvetor ; i += 2 {

            fmt.Printf("%d %d\n", after[i], after[i + 1])
        }

    } else {

        switch direction {
        case "L":
            fmt.Printf("%d %d\n", px - 1, py)

        case "R":
            fmt.Printf("%d %d\n", px + 1, py)

        case "U":
            fmt.Printf("%d %d\n", px, py - 1)

        case "D":
            fmt.Printf("%d %d\n", px, py + 1)
        }
    }



    

    
}
