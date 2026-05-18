package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func éPrimoAux(n int, divisor int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	if divisor*divisor > n {
		return true
	}
	return éPrimoAux(n, divisor+1)
}

func ÉPrimo(n int) bool {
	return éPrimoAux(n, 2)
}

func carregaPrimosAux(n int, atual int, primos []int) []int {
	if len(primos) == n {
		return primos
	}

	if ÉPrimo(atual) {
		primos = append(primos, atual)
	}

	return carregaPrimosAux(n, atual+1, primos)
}

func CarregaPrimos(n int) []int {
	return carregaPrimosAux(n, 2, make([]int, 0, n))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(input)
		if err == nil {
			resultado := CarregaPrimos(n)
			fmt.Println(strings.ReplaceAll(fmt.Sprint(resultado), " ", ", "))
		}
	}
}
