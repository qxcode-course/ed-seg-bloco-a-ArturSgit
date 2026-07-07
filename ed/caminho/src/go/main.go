package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{l: p.l - 1, c: p.c}, // cima
		{l: p.l + 1, c: p.c}, // baixo
		{l: p.l, c: p.c - 1}, // esquerda
		{l: p.l, c: p.c + 1}, // direita
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := NewQueue[Pos]()
	queue.Enqueue(startPos)

	visited := make(map[Pos]bool)
	visited[startPos] = true

	caminho := make(map[Pos]Pos)
	found := false

	for !queue.IsEmpty() {
		atual, _ := queue.Dequeue()

		if atual == endPos {
			found = true
			break
		}

		for _, vizinho := range atual.getNeig() {

			if !visited[vizinho] && match(grid, vizinho, ' ') {
				visited[vizinho] = true
				caminho[vizinho] = atual
				queue.Enqueue(vizinho)
			}
		}
	}

	if found {
		voltar(grid, startPos, endPos, caminho)
	}
}

func voltar(grid [][]rune, startPos Pos, endPos Pos, caminho map[Pos]Pos) {

	atual := caminho[endPos]

	for atual != startPos {
		grid[atual.l][atual.c] = '.'
		atual = caminho[atual]
	}

	grid[startPos.l][startPos.c] = '.'
	grid[endPos.l][endPos.c] = '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
