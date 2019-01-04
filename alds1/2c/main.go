package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type card struct {
	suit, num byte
}

func (c card) String() string {
	return fmt.Sprint(string(c.suit), string(c.num))
}

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func parseCard(s string) *card {
	c := &card{
		suit: s[0],
		num:  s[1],
	}
	return c
}

func main() {
	var a []card
	sc.Split(bufio.ScanWords)
	n, err := strconv.Atoi(next())
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; {
		c := parseCard(next())
		a = append(a, *c)
		i++
	}

	a1 := make([]card, len(a))
	copy(a1, a)

	bubbleSortByNum(a, n)

	selectionSortByNum(a1, n)

	printFormattedArr(a)
	printStability(isBubbleSortStable(a, a1))

	printFormattedArr(a1)
	printStability(isSelectionSortStable(a, a1))
}

func bubbleSortByNum(a []card, n int) {
	swapped := true
	for swapped {
		swapped = false
		for i := n - 1; i >= 1; i-- {
			if a[i-1].num > a[i].num {
				a[i-1], a[i] = a[i], a[i-1]
				swapped = true
			}
		}
	}
}

func selectionSortByNum(a []card, n int) {
	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if a[j].num < a[min].num {
				min = j
			}
		}
		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}

func isBubbleSortStable(_, _ []card) bool {
	return true
}

func isSelectionSortStable(a1, a2 []card) bool {
	return reflect.DeepEqual(a1, a2)
}

func printStability(b bool) {
	if b {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func printFormattedArr(a []card) {
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}
