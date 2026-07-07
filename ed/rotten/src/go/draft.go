package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	l, c int
}

func funcRottingOranges(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	nrows := len(grid)
	ncols := len(grid[0])
	queue := NewQueue[Pos]()
	freshOranges := 0

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			if grid[i][j] == 2 {
				queue.Enqueue(Pos{l: i, c: j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	minutes := 0
	dirs := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for !queue.IsEmpty() {

		size := queue.queue.Len()
		infectedThisMinute := false

		for k := 0; k < size; k++ {

			atual, ok := queue.Dequeue()

			if !ok {
				break
			}
			for _, d := range dirs {
				nl, nc := atual.l+d.l, atual.c+d.c

				if nl >= 0 && nl < nrows && nc >= 0 && nc < ncols && grid[nl][nc] == 1 {
					grid[nl][nc] = 2
					freshOranges--
					infectedThisMinute = true
					queue.Enqueue(Pos{l: nl, c: nc})
				}
			}
		}

		if infectedThisMinute {
			minutes++
		}
	}

	if freshOranges > 0 {
		return -1
	}
	return minutes
}

type Queue[T any] struct {
	queue *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{queue: list.New()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.queue.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	element := q.queue.Front()
	if element == nil {
		var zero T
		return zero, false
	}
	q.queue.Remove(element)
	value := element.Value.(T)
	return value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.queue.Len() == 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	grid := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowParts := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowParts[j])
		}

		grid[i] = row
	}
	fmt.Println(funcRottingOranges(grid))
}
