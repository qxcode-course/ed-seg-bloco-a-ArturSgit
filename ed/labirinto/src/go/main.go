package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	// Ordem: Esquerda, Cima, Direita, Baixo
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func solve(grid [][]rune, curr, end Pos) bool {
	// Se chegamos no F, marcamos como ponto e retornamos sucesso
	if curr == end {
		grid[curr.l][curr.c] = '.'
		return true
	}

	// Salva para backtracking
	original := grid[curr.l][curr.c]
	// Marca a posição atual como parte do caminho
	grid[curr.l][curr.c] = '.'

	for _, next := range getNeig(curr) {
		if next.l >= 0 && next.l < len(grid) && next.c >= 0 && next.c < len(grid[0]) {
			// Aceita espaços vazios ou o destino final
			if grid[next.l][next.c] == ' ' || grid[next.l][next.c] == 'F' {
				if solve(grid, next, end) {
					return true
				}
			}
		}
	}

	// Backtracking: restaura o caractere original (espaço ou 'I')
	grid[curr.l][curr.c] = original
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])

	grid := make([][]rune, nl)
	var startPos, endPos Pos

	for i := 0; i < nl; i++ {
		if scanner.Scan() {
			grid[i] = []rune(scanner.Text())
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 'I' {
					startPos = Pos{i, j}
				} else if grid[i][j] == 'F' {
					endPos = Pos{i, j}
				}
			}
		}
	}

	solve(grid, startPos, endPos)

	// Imprimir resultado
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
