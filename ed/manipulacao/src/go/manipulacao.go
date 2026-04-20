package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
	"slices"
)

func getMen(vet []int) []int {
	// slice sem alocação de memória
	var men []int
	for _, v := range vet {
		
		if v > 0 {
			men = append(men, v)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	
	var women []int
	for _, v := range vet {

		if v < 0 && v > -10 {
			women = append(women, v)

		}
	}
	return women
}

func sortVet(vet []int) []int {
	
	// função sort para auxiliar
	// os ... faz com q pegue um elemnto por vez de vet
	copia := append([]int{}, vet...)
	sort.Ints(copia)
	return copia
}

func abs(n int) int {

	if n < 0 {
		return -n
	}
	return n
}
func sortStress(vet []int) []int {
	
	copia := append([]int{}, vet...)
	sort.Slice(copia, func (i, j int) bool {
		
		return abs(copia[i]) < abs (copia[j])
	})

	return copia
}

func reverse(vet []int) []int {
	
	copia := slices.Clone(vet)
	slices.Reverse(copia)
	return copia
	
}

func unique(vet []int) []int {
	
	var copia []int 
	procurar := make(map[int]bool)

	for _, v := range vet {

		if !procurar[v] {

			procurar[v] = true
			copia = append(copia, v)
		}
	}
	return copia

}

func repeated(vet []int) []int {
	
	counts := make(map[int]int)
	var res []int
	
	for _, v := range vet {

		counts[v]++
	}

	for _, v := range vet {
		
		if counts[v] > 1 {
			res = append(res, v)
			counts[v]--
		}
	}
	slices.Sort(res)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

