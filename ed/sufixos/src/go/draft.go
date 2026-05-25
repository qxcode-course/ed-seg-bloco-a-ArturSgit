package main

import "fmt"

func sufixos(palavra string) {

	if len(palavra) == 0 {

		return
	}

	sufixos(palavra[1:])
	fmt.Println(palavra)
}

func main() {

	var palavra string
	_, err := fmt.Scan(&palavra)
	if err != nil {
		return
	}

	sufixos(palavra)

}
