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

func enesimoPrimoAux(n int, atual int) int {
	if ÉPrimo(atual) {
		if n == 1 {
			return atual
		}
		return enesimoPrimoAux(n-1, atual+1)
	}
	return enesimoPrimoAux(n, atual+1)
}

func EnesimoPrimo(n int) int {
	return enesimoPrimoAux(n, 2)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(input)
		if err == nil {
			fmt.Println(EnesimoPrimo(n))
		}
	}
}
