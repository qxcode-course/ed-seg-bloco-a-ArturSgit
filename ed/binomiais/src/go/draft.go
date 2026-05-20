package main

import (
	"fmt"
	"io"
	"os"
)

func binomialcoeficient(n, k int) int {

	if k < 0 || k > n {

		return 0
	}
	if k == 0 || k == n {

		return 1
	}
	if k > n/2 {
		k = n - k
	}

	result := 1
	for i := 1; i <= k; i++ {
		result = result * (n - k + i) / i
	}

	return result
}

func main() {

	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada", err)
		return
	}
	fmt.Println(binomialcoeficient(n, k))
}
