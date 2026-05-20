package main

import (
	"fmt"
)

func calcular(atual int, alvo int, dp []int) int {
	if atual > alvo {
		return 0
	}
	if atual == alvo {
		return 1
	}

	if dp[atual] != 0 {
		return dp[atual]
	}

	dp[atual] = calcular(atual + 1, alvo, dp) + calcular(atual+3, alvo, dp)
	return dp[atual]
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	if n == 1 || n == 2 {
		fmt.Println(1)
		return
	}
	if n == 3 {
		fmt.Println(2)
		return
	}

	dp := make([]int, n+1)
	fmt.Println(calcular(0, n, dp))
}