package main

import "fmt"

func solve(n int) {

	if n == 0 {
		return
	}
	div := n / 2
	resto := n % 2
	solve(div)

	fmt.Printf("%d %d\n", div, resto)
}
func main() {

	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	solve(n)

}
