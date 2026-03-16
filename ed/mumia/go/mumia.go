package main
import "fmt"

func main() {

    var name, namep string
    var age int

    fmt.Scan(&namep, &age)

    if age < 12 {
        
        name = "crianca"

    } else if age < 18 {
        
        name = "jovem"

    } else if age < 65 {

        name = "adulto"

    } else if age < 1000 {

        name = "idoso"

    } else {

        name = "mumia"
    }

    fmt.Printf ("%s eh %s\n", namep, name)





}
