package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// tostr: Converte o vetor para string "[1, 2, 3]" de forma recursiva.
func tostr(vet []int) string {
	// Caso base: se o vetor está vazio, retorna colchetes vazios.
	if len(vet) == 0 {
		return "[]"
	}

	// Função auxiliar para lidar com a lógica dos elementos.
	var rec func(v []int) string
	rec = func(v []int) string {
		// Caso base: se só restou um elemento, retorna ele sem vírgula.
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		// Passo recursivo: elemento atual + vírgula e ESPAÇO + resto.
		// Note o ", " com espaço para bater com o esperado nos testes.
		return strconv.Itoa(v[0]) + ", " + rec(v[1:])
	}

	return "[" + rec(vet) + "]"
}

// tostrrev: Converte para string de trás para frente "[3, 2, 1]".
func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	var rec func(v []int) string
	rec = func(v []int) string {
		if len(v) == 1 {
			return strconv.Itoa(v[0])
		}
		// Inverte a ordem: primeiro chama o resto, depois coloca o atual.
		return rec(v[1:]) + ", " + strconv.Itoa(v[0])
	}

	return "[" + rec(vet) + "]"
}

// reverse: Inverte o vetor original (in-place).
func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}
	// Troca as extremidades.
	vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
	// Recursão no miolo do vetor.
	reverse(vet[1 : len(vet)-1])
}

// sum: Soma recursiva.
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: Produto recursivo.
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: Retorna o índice do menor valor.
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(v []int, i int) int
	rec = func(v []int, i int) int {
		if i == len(v)-1 {
			return i
		}
		menorDoResto := rec(v, i+1)
		if v[i] <= v[menorDoResto] {
			return i
		}
		return menorDoResto
	}
	return rec(vet, 0)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println("$" + line)
		args := strings.Fields(line)

		switch args[0] {
		case "read":
			vet = []int{}
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

// Join: Função utilitária (certifique-se de que está fora da main).
func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}