package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	s  = newStack()
)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	line := nextLine()
	tokens := strings.Split(line, " ")
	for _, t := range tokens {
		processToken(t)
	}
	fmt.Println(s.pop())
}

func processToken(t string) {

	a, b := -100, -100

	switch t {
	case "+":
		a = s.pop()
		b = s.pop()
		s.push(a + b)
	case "-":
		a = s.pop()
		b = s.pop()
		s.push(b - a)
	case "*":
		a = s.pop()
		b = s.pop()
		s.push(a * b)
	default:
		n, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		s.push(n)
	}
}

type stack struct {
	buf []int
	top int
}

func newStack() *stack {
	return &stack{
		buf: make([]int, 200),
	}
}

func (s *stack) push(x int) {
	s.buf[s.top] = x
	s.top++
}

func (s *stack) pop() int {
	s.top--
	return s.buf[s.top]
}
