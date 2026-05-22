package main

import "fmt"

func diagonal(pal string, qtd int) {

	if len(pal) == 0 {
		return
	}

	for i := 0; i < qtd; i++ {

		fmt.Print(" ")
	}

	fmt.Printf("%c\n", pal[0])
	diagonal(pal[1:], qtd+1)
}

func main() {

	var palavra string
	_, err := fmt.Scan(&palavra)
	if err != nil {
		return
	}

	diagonal(palavra, 0)
}
