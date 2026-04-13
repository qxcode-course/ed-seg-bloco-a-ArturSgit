package main

import "fmt"

// funcao para inicializar
func initial(n int, e int){

	// inicializacao
	vivos := make([]int, n)
	for i := 0 ; i < n ; i++ {
		vivos[i] = i + 1 
	}

	// serve para encontrar no vetor a espada
	index_espada := e - 1

	for len(vivos) > 1 {

		impressao(vivos, index_espada)

		index_alvo := (index_espada + 1) % len(vivos)

		vivos = append(vivos[:index_alvo], vivos[index_alvo + 1:]... )

		index_espada = index_alvo % len(vivos)
	}

	impressao(vivos, 0)


}

func impressao(vivos[] int, index_espada int){

	fmt.Print("[ ")

	for i, v := range vivos {
		
		if i == index_espada {

			fmt.Printf("%d> ", v)

		} else  {

			fmt.Printf("%d ", v)
		}
	}

	fmt.Print("]\n")
}


func main() {

	var n int
	var e int

	fmt.Scan(&n, &e)

	initial(n, e)
}
