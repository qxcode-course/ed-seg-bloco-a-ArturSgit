package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {

	if capacity <= 0 {
		capacity = 1
	}
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, newCapacity)
	copy(newData, s.data)
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) Search(value int) int {

	left := 0
	right := s.size - 1

	for left <= right {
		mid := left + (right-left)/2
		if s.data[mid] == value {
			return mid
		} else if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}

func (s *Set) InsertPosition(value int) int {

	left := 0
	right := s.size - 1
	pos := s.size

	for left <= right {

		mid := left + (right-left)/2
		if s.data[mid] >= value {

			pos = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return pos
}

func (s *Set) insert(value int, index int) {

	if s.size == s.capacity {
		s.reserve(s.capacity * 2)
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++

}

func (s *Set) Insert(value int) {

	if s.size == 0 {
		s.insert(value, 0)
		return
	}

	if s.Search(value) != -1 {
		return
	}
	pos := s.InsertPosition(value)
	s.insert(value, pos)

}

func (s *Set) erase(index int) {

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
}

func (s *Set) Erase(value int) bool {

	idx := s.Search(value)
	if idx == -1 {
		return false
	}
	s.erase(idx)
	return true
}

func (s *Set) Contains(value int) bool {

	return s.Search(value) != -1
}

func (s *Set) String() string {

	if s.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")

	for i := 0; i < s.size; i++ {

		sb.WriteString(strconv.Itoa(s.data[i]))

		if i < s.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var s *Set

	// v := NewSet(0)
	for scanner.Scan() {
		line = scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]
		fmt.Printf("$%s\n", line)

		switch cmd {

		case "end":

			return

		case "init":

			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)

		case "insert":

			if s != nil {
				for _, part := range parts[1:] {
					value, _ := strconv.Atoi(part)
					s.Insert(value)
				}
			}

		case "show":
			if s != nil {
				fmt.Println(s.String())
			} else {
				fmt.Println("[]")
			}

		case "erase":

			if s != nil {

				value, _ := strconv.Atoi(parts[1])

				if !s.Erase(value) {
					fmt.Printf("value not found\n")
				}
			}

		case "contains":

			if s != nil {

				value, _ := strconv.Atoi(parts[1])
				if s.Contains(value) {
					fmt.Printf("true\n")
				} else {
					fmt.Printf("false\n")
				}
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
