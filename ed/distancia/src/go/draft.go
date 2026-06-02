package main

import "fmt"

func valido(seq []rune, index int, num int, L int) bool {
	charNum := rune('0' + num)

	inicioEsquerda := index - L
	if inicioEsquerda < 0 {
		inicioEsquerda = 0
	}

	for i := inicioEsquerda; i < index; i++ {
		if seq[i] == charNum {
			return false
		}
	}

	fimDireita := index + L
	if fimDireita >= len(seq) {
		fimDireita = len(seq) - 1
	}

	for i := index + 1; i <= fimDireita; i++ {
		if seq[i] == charNum {
			return false
		}
	}

	return true
}

func resolver(seq []rune, L int, index int) (string, bool) {

	if index == len(seq) {
		return string(seq), true
	}

	if seq[index] != '.' {
		return resolver(seq, L, index+1)
	}

	for num := 0; num <= L; num++ {
		if valido(seq, index, num, L) {
			seq[index] = rune('0' + num)

			if resultado, ok := resolver(seq, L, index+1); ok {
				return resultado, true
			}

			seq[index] = '.'
		}
	}

	return "", false

}

func main() {
	var entrada string
	var L int

	_, err := fmt.Scan(&entrada)
	if err != nil {
		return
	}

	_, err = fmt.Scan(&L)
	if err != nil {
		return
	}

	seqRunes := []rune(entrada)
	if solucao, ok := resolver(seqRunes, L, 0); ok {
		fmt.Println(solucao)
	}

}
