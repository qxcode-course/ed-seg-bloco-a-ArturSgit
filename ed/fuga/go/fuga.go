package main
import "fmt"
func main() {

    var h, p, f, d int 
    var result string
    fmt.Scan(&h, &p, &f, &d)

    if d == 1 { // anti-horario
        
        for  {

            

            if f == p {
                
                result = "N"
                break
                
            } else if f == h {
                    
                result = "S"
                break            }
                    
            f++
            if f > 15 {
                      
                f = 0
        
            }
        }
    } else {

        for {

            if f == p {
                
                result = "N"
                break
                
            } else if f == h {
                    
                    result = "S"
                    break            }
                    
            f--
                
            if f < 0 {
              
                f = 15
    
            } 
        }
    }
            
    fmt.Printf("%s\n", result)
    
    
}
