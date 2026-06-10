package main

import (
	"fmt"
	"math"
)

func naLinha(matriz [][]rune, lin int, ch rune) bool {

	for i := 0; i < len(matriz); i++ {
		if matriz[lin][i] == ch {
			return true
		}
	}
	return false
}

func naColuna(matriz [][]rune, col int, ch rune) bool {

	for i := 0; i < len(matriz); i++ {
		if matriz[i][col] == ch {
			return true
		}
	}
	return false
}

func Quadrante(matriz [][]rune, lin int, col int, ch rune) bool {
	dim := len(matriz)
	tamSubgrade := int(math.Sqrt(float64(dim)))

	inicioLinha := (lin / tamSubgrade) * tamSubgrade
	inicioColuna := (col / tamSubgrade) * tamSubgrade

	for i := 0; i < tamSubgrade; i++ {
		for j := 0; j < tamSubgrade; j++ {
			if matriz[inicioLinha+i][inicioColuna+j] == ch {
				return true
			}
		}
	}
	return false
}

func resolver(matriz [][]rune, index int) bool {

	n := len(matriz)
	if index == n*n {
		return true
	}

	lin := index / n
	col := index % n

	if matriz[lin][col] != '.' {
		return resolver(matriz, index+1)
	}

	for num := '1'; num <= rune('0'+n); num++ {
		if !naLinha(matriz, lin, num) && !naColuna(matriz, col, num) && !Quadrante(matriz, lin, col, num) {
			matriz[lin][col] = num

			if resolver(matriz, index+1) {
				return true
			}
			matriz[lin][col] = '.'
		}
	}

	return false
}

func main() {

	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	matriz := make([][]rune, n)
	for i := 0; i < n; i++ {
		var linha string
		fmt.Scan(&linha)
		matriz[i] = []rune(linha)
	}

	if resolver(matriz, 0) {
		for i := 0; i < n; i++ {
			fmt.Println(string(matriz[i]))
		}
	}
}
