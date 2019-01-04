package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

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
	insertSort(a, n)
	printFormattedArr(a)
}

func insertSort(a []int, n int) {
	for i := 1; i < n; i++ {
		printFormattedArr(a)
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
	}
}

func printFormattedArr(a []int) {
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}
