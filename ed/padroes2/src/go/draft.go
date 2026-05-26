package main

import (
	"fmt"
)

func padroes(n int) int {

	if n == 1 {
		return 3
	}

	return padroes(n-1) + (2*n + 1)
}
func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	totpecas := padroes(n)
	fmt.Println(totpecas)
}
