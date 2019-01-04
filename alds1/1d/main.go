package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() (i int) {
	sc.Scan()
	i, _ = strconv.Atoi(sc.Text())
	return
}

func main() {
	n := next()
	max, min := -2000000000, 2000000000
	var s int

	for i := 0; i < n; {
		s = next()
		if max < s-min {
			max = s - min
		}
		if min > s {
			min = s
		}
		i++
	}
	fmt.Println(max)
}
