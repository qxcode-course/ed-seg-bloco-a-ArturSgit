package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

func (n *Node) Next() *Node {

	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {

	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root, size: 0}
}

func (ll *LList) Size() int {

	return ll.size
}

func (ll *LList) Clear() {

	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *LList) insertNode(value int, at *Node) *Node {

	newNode := &Node{
		Value: value,
		next:  at,
		prev:  at.prev,
		root:  ll.root,
	}
	at.prev.next = newNode
	at.prev = newNode
	ll.size++
	return newNode
}

func (ll *LList) removeNode(node *Node) *Node {

	if node == ll.root {
		return nil
	}
	nextNode := node.next

	node.prev.next = node.next
	node.next.prev = node.prev

	ll.size--

	if nextNode == ll.root {
		return nil
	}
	return nextNode
}

func (ll *LList) PushFront(value int) {
	ll.insertNode(value, ll.root.next)
}

func (ll *LList) PushBack(value int) {
	ll.insertNode(value, ll.root)
}

func (ll *LList) PopFront() {
	if ll.size > 0 {
		ll.removeNode(ll.root.next)
	}
}

func (ll *LList) PopBack() {
	if ll.size > 0 {
		ll.removeNode(ll.root.prev)
	}
}

func (ll *LList) Front() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.prev
}

func (ll *LList) Search(value int) *Node {
	for curr := ll.root.next; curr != ll.root; curr = curr.next {
		if curr.Value == value {
			return curr
		}
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	if node != nil {
		ll.insertNode(value, node)
	}
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil {
		return nil
	}
	return ll.removeNode(node)
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	curr := ll.root.next

	for curr != ll.root {
		sb.WriteString(strconv.Itoa(curr.Value))
		if curr.next != ll.root {
			sb.WriteString(", ")
		}
		curr = curr.next
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "$") {
			fmt.Println(line)
			line = strings.TrimPrefix(line, "$")
		} else {
			fmt.Println("$" + line)
		}

		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {

		case "show":

			fmt.Println(ll.String())

		case "size":

			fmt.Println(ll.Size())

		case "push_back":

			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}

		case "push_front":

			for i := 1; i < len(args); i++ {
				num, _ := strconv.Atoi(args[i])
				ll.PushFront(num)
			}

		case "pop_back":

			ll.PopBack()

		case "pop_front":

			ll.PopFront()

		case "clear":

			ll.Clear()

		case "walk":

			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")

		case "replace":

			if len(args) < 3 {
				continue
			}
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}

		case "insert":

			if len(args) < 3 {
				continue
			}
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}

		case "remove":

			if len(args) < 2 {
				continue
			}
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}

		case "end":

			return

		default:

			fmt.Println("fail: comando invalido")

		}
	}
}
