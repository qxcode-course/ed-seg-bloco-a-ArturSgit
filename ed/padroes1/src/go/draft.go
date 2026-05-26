package main

import "fmt"

func numeroBlocos(n int) int {
	return 8*n + 12
}
func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Println(numeroBlocos(n))
}
