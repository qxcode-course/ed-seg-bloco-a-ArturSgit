package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(vet []int, primeiro int, ultimo int) {
	for primeiro < ultimo {
		vet[primeiro], vet[ultimo] = vet[ultimo], vet[primeiro]
		primeiro++
		ultimo--
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	firstLine := strings.Fields(scanner.Text())
	if len(firstLine) < 2 {
		return
	}

	t, _ := strconv.Atoi(firstLine[0])
	r, _ := strconv.Atoi(firstLine[1])

	if !scanner.Scan() {
		return
	}

	elementoStr := strings.Fields(scanner.Text())
	vet := make([]int, t)
	for i := 0; i < t; i++ {
		vet[i], _ = strconv.Atoi(elementoStr[i])
	}

	if t > 0 {

		r = r % t

		if r > 0 {
			reverse(vet, 0, t-1)
			reverse(vet, 0, r-1)
			reverse(vet, r, t-1)
		}
	}

	fmt.Print("[ ")
	for i := 0; i < len(vet); i++ {

		fmt.Print(vet[i])
		if i < len(vet)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}
