package main
import "fmt"

func main() {
    
    var totalbum, totbaruel int
    fmt.Scanln(&totalbum)
    fmt.Scanln(&totbaruel)

    figalbum := make([]int, totalbum)
    figbaruel := make([]int, totbaruel)

    for i := 0 ; i < totbaruel ; i++ {

        fmt.Scan(&figbaruel[i])
    }
    for i := 0 ; i < totalbum ; i++ {

        figalbum[i] = i + 1
    }


    var count1, count2 int = 0, 0
    for i := 0 ; i < totalbum ; i++ {

        for j := 0 ; j < totbaruel ; j++ {


            if figalbum[i] == figbaruel[j] {

                figalbum[i] = 0
                figbaruel[j] = 0
                break
            }
        }
    }

    firstbaruel := true

    for i := 0 ; i < totbaruel ; i++ {

        if figbaruel[i] != 0 {
            count1++
        }
    }

    for i := 0 ; i < totalbum ; i++ {

        if figalbum[i] != 0 {
            count2++
        }
    }

    if count1 == 0 {
        fmt.Println("N")

    } else {

        for i := 0 ; i < totbaruel ; i++ {
    
            if figbaruel[i] != 0 {
    
                if !firstbaruel {
                
                    fmt.Print(" ")
                }
    
                fmt.Print(figbaruel[i])
                firstbaruel = false
            }
    
        }
        fmt.Println()
    }
    
    firstalbum := true
    if count2 == 0 {
        fmt.Println("N")

    } else {

        for i := 0 ; i < totalbum ; i++ {
    
            if count2 == 0 {
    
                fmt.Println("N")
    
            } else if figalbum[i] != 0 {
    
                
                if !firstalbum {
                    
                    fmt.Print(" ")
                }
                
                fmt.Print(figalbum[i])
                firstalbum = false
            }
        }
        fmt.Println()

    }


}
