package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func abs( n int ) int {

	if n < 0 {
		return -n
	}
	return n
}
func occurr(vet []int) []Pair {
	counts := make(map[int]int)
	for _, v := range vet {
		counts[abs(v)]++
	}

	var keys []int
	for k := range counts {
		keys = append(keys,k)
	}
	sort.Ints(keys)

	var res []Pair
	for _, k := range keys {
		res = append(res, Pair{k, counts[k]})
	}
	return res
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	var res []Pair
	start := 0
	for i := 1 ; i <= len(vet) ; i++ {
		if i == len(vet) || vet[i] != vet[start] {
			res = append(res, Pair{vet[start], i - start})
			start = i
		}
	}
	return res
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))

	for i := range vet {

		if vet[i] > 0 {
			tem_ao_lado := false
			
			if i > 0 && vet[i-1] < 0 {
				tem_ao_lado = true
			}
			
			if i < len(vet) - 1 && vet[i+1] < 0 {
				tem_ao_lado = true
			} 
			
			if tem_ao_lado {
				res[i] = 1
			}
		}
	}
	return res

}

func alone(vet []int) []int {
	
	res := make([]int, len(vet))

	for i := range vet {

		if vet[i] > 0 {
			tem_ao_lado := true
			
			if i > 0 && vet[i-1] < 0 {
				tem_ao_lado = false
			}
			
			if i < len(vet) - 1 && vet[i+1] < 0 {
				tem_ao_lado = false
			} 
			
			if tem_ao_lado {
				res[i] = 1
			}
		}
	}
	return res
	

}

func couple(vet []int) int {
	count := 0
	res := make([]int, len(vet))

	for i := range vet {

		if vet[i] > 0 {
			tem_ao_lado := false
			
			if (i > 0 && vet[i-1] < 0) ||  (i < len(vet) - 1 && vet[i+1] < 0) {
				tem_ao_lado = true
				count++
				
				if (i > 0 && vet[i-1] > 0) {
					count--
				}

			}
			if tem_ao_lado {
				res[i] = 1
				
			}
		}
	}
	return count
}

// func hasSubseq(vet []int, seq []int, pos int) bool {
// 	_ = vet
// 	_ = seq
// 	_ = pos
// 	return false
// }

func subseq(vet []int, seq []int) int {
	if len(seq) > len(vet) {
		return -1
	}
	if len(seq) == 0 {
		return 0
	}

	for i := 0 ; i <= len(vet) - len(seq) ; i++ {
		match := true
		for j := 0 ; j < len(seq) ; j++ {
			if vet[i+j] != seq[j] {
				match = false
				break
			}
		} 
		if match {
			return i
		}
	} 
	return -1
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
    for _, p := range posList {
        toRemove[p] = true
    }
    var res []int
    for i, v := range vet {
        if !toRemove[i] {
            res = append(res, v)
        }
    }
    return res
}

func clear(vet []int, value int) []int {
	var res []int
    for _, v := range vet {
        if v != value {
            res = append(res, v)
        }
    }
    return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
