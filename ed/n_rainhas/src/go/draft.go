package main

import (
	"fmt"
	"io"
	"os"
)

func solveNQueens(n int) int {
	cols := make([]bool, n)
	diagPrincipal := make([]bool, 2*n-1)
	diagSecundaria := make([]bool, 2*n-1)

	solucoes := 0

	var backtrack func(linha int)
	backtrack = func(linha int) {
		if linha == n {
			solucoes++
			return
		}

		for col := 0; col < n; col++ {

			idPrincipal := linha - col + (n - 1)
			idSecundaria := linha + col

			if cols[col] || diagPrincipal[idPrincipal] || diagSecundaria[idSecundaria] {
				continue
			}

			cols[col] = true
			diagPrincipal[idPrincipal] = true
			diagSecundaria[idSecundaria] = true

			backtrack(linha + 1)

			cols[col] = false
			diagPrincipal[idPrincipal] = false
			diagSecundaria[idSecundaria] = false
		}
	}

	backtrack(0)
	return solucoes
}

func main() {

	var n int

	_, err := fmt.Scan(&n)
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
		return
	}
	resultado := solveNQueens(n)
	fmt.Println(resultado)

}
