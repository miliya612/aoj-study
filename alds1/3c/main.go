package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	result = newDLL()
)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() (int, error) {
	sc.Scan()
	return strconv.Atoi(sc.Text())
}

func main() {
	sc.Split(bufio.ScanWords)
	n, err := nextInt()
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		line := nextLine()
		exec(line)
	}
	printRes()
}

func printRes() {
	b := bytes.Buffer{}
	n := result.dummy.next
	if n == result.dummy {
		return
	}
	for {
		b.WriteString(strconv.Itoa(n.x))
		n = n.next
		if n == result.dummy {
			break
		}
		b.WriteString(" ")
	}
	fmt.Println(b.String())
}

func exec(c string) {
	if c[0] == 'i' {
		i, _ := nextInt()
		result.insert(i)
	} else if len(c) == 6 {
		i, _ := nextInt()
		result.deleteKey(i)
	} else if len(c) == 11 {
		result.deleteFirst()
	} else {
		result.deleteLast()
	}
}

type node struct {
	x          int
	prev, next *node
}

type dll struct {
	dummy *node
}

func newDLL() *dll {
	d := new(node)
	d.next = d
	d.prev = d
	return &dll{
		dummy: d,
	}
}

// insert inserts an element to the head of list
func (d *dll) insert(i int) {
	n := new(node)

	n.x = i
	n.prev = d.dummy
	n.next = d.dummy.next

	d.dummy.next.prev = n
	d.dummy.next = n

}

func (d *dll) deleteKey(k int) {
	n := d.find(k)
	if n.x != k {
		return
	}
	d.deleteNode(n)
}

func (d *dll) deleteNode(n *node) {
	if n == d.dummy {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (d *dll) deleteFirst() {
	d.deleteNode(d.dummy.next)
}

func (d *dll) deleteLast() {
	d.deleteNode(d.dummy.prev)
}

func (d *dll) find(k int) (n *node) {
	n = d.dummy.next
	for n != d.dummy && n.x != k {
		n = n.next
	}
	return
}
