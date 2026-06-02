package main

import "fmt"

func subsetSum(arr []int, n int, target int, index int, currentSum int) bool {
	if currentSum == target {
		return true
	}

	if currentSum > target || index >= n {
		return false
	}

	if subsetSum(arr, n, target, index+1, currentSum+arr[index]) {
		return true
	}
	if subsetSum(arr, n, target, index+1, currentSum) {
		return true
	}

	return false
}

func main() {

	var n, k int

	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	if subsetSum(arr, n, k, 0, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
