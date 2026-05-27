package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var entrada string

	if _, err := fmt.Scan(&entrada); err != nil {
		return
	}

	caracteres := strings.Split(entrada, "")
	sort.Strings(caracteres)

	atual := make([]string, 0, len(caracteres))
	usado := make([]bool, len(caracteres))

	var backtrack func()
	backtrack = func() {
		if len(atual) == len(caracteres) {
			fmt.Println(strings.Join(atual, ""))
			return
		}
		for i := 0; i < len(caracteres); i++ {

			if usado[i] {
				continue
			}

			usado[i] = true
			atual = append(atual, caracteres[i])

			backtrack()
			atual = atual[:len(atual)-1]
			usado[i] = false
		}
	}

	backtrack()

}
