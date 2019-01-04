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
	for i := 0; i < n; i++ {
		a = append(a, next())
	}
	selectionSort(a, n)
	printFormattedArr(a)
	fmt.Println(count)
}

func selectionSort(a []int, n int) {
	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			a[i], a[min] = a[min], a[i]
			count++
		}
	}
}

func printFormattedArr(a []int) {
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}
