package main
import "fmt"

func main() {

    var numb int
    fmt.Scan(&numb)

    //poderia ser feita por meio de verificaçao em boolean

    animals := make([]int, numb)
    partners := 0

    for i := 0 ; i < numb ; i++ {
        fmt.Scan(&animals[i])
    }

    for i := 0 ; i < numb ; i++ {
        
        for j := i + 1 ; j < numb ; j++ {
            
            if (animals[i] + animals[j]) == 0 {
                partners++
                animals[j] = 1000
                break 
                
            }
        }
    }
    fmt.Println(partners)

}
