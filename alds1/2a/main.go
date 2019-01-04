package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	count = 0
)

func next() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	var a []int
	sc.Split(bufio.ScanWords)
	n := next()
	for i := 0; i < n; {
		a = append(a, next())
		i++
	}
	flagBubbleSort(a, n)
	printFormattedArr(a)
	fmt.Println(count)
}

func bubbleSort(a []int, n int) {
	for j := n - 1; j >= 1; j-- {
		for k := n - 1; k >= 1; k-- {
			if a[k-1] > a[k] {
				a[k-1], a[k] = a[k], a[k-1]
				count++
			}
		}
	}
}

// flagBubbleSort is an implementation of bubble sort
// it improves the time complexity in the best case of already sorted
// https://hackr.io/blog/bubble-sort-in-c
func flagBubbleSort(a []int, n int) {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for j := n - 1; j >= 1; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				count++
				isSwapped = true
			}
			printFormattedArr(a)
		}
	}
}

func printFormattedArr(a []int) {
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}
