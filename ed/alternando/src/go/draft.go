package main

import (
	"fmt"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func imprimirEstado(lista []int, idxEspada int) {
	fmt.Print("[")
	for i, p := range lista {
		fmt.Print(" ")
		if i == idxEspada {
			if p > 0 {
				fmt.Printf("%d>", p)
			} else {
				fmt.Printf("<%d", p)
			}
		} else {
			fmt.Printf("%d", p)
		}
	}
	fmt.Println(" ]")
}

func main() {
	var n, e, f int
	if _, err := fmt.Scan(&n, &e, &f); err != nil {
		return
	}

	participantes := make([]int, n)
	sinal := f
	for i := 0; i < n; i++ {
		participantes[i] = (i + 1) * sinal
		sinal *= -1
	}

	pos := 0
	for i, v := range participantes {
		if v == e || v == -e {
			pos = i
			break
		}
	}

	for len(participantes) > 1 {
		imprimirEstado(participantes, pos)

		atual := participantes[pos]
		var alvo int

		if atual > 0 {
			alvo = mod(pos+1, len(participantes))
		} else {
			alvo = mod(pos-1, len(participantes))
		}

		quemEstaComAEspada := participantes[pos]

		participantes = append(participantes[:alvo], participantes[alvo+1:]...)

		for i, v := range participantes {
			if v == quemEstaComAEspada {
				pos = i
				break
			}
		}

		if atual > 0 {
			pos = mod(pos+1, len(participantes))
		} else {
			pos = mod(pos-1, len(participantes))
		}
	}

	imprimirEstado(participantes, 0)
}
