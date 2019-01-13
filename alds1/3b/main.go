package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc          = bufio.NewScanner(os.Stdin)
	elapsedTime int
	opts        *options
	q           = newQueue()
	processed   = []process{}
)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func parseProcess(s string) (*process, error) {
	strs := strings.Split(s, " ")
	t, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}

	return &process{
		name:      strs[0],
		remaining: t,
	}, nil
}

type options struct {
	processNum, quantumMS int
}

func parseOpts(s string) (*options, error) {
	strs := strings.Split(s, " ")
	processNum, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil, err
	}
	quantumMS, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}

	return &options{
		processNum: processNum,
		quantumMS:  quantumMS,
	}, nil
}

func main() {
	var err error
	line := nextLine()
	opts, err = parseOpts(line)
	if err != nil {
		panic(err)
	}
	for i := 0; i < opts.processNum; i++ {
		str := nextLine()
		p, err := parseProcess(str)
		if err != nil {
			panic(err)
		}
		q.enqueue(*p)
	}
	for q.size() >= 1 {
		p := q.dequeue()
		p.exec()
		if !p.done() {
			q.enqueue(p)
		}
	}
	printRes()
}

type process struct {
	name      string
	remaining int
	endAt     int
}

func (p *process) exec() {
	if p.remaining <= opts.quantumMS {
		p.endAt = elapsedTime + p.remaining
		elapsedTime += p.remaining
		p.remaining = 0
		processed = append(processed, *p)
		return
	}
	p.remaining -= opts.quantumMS
	elapsedTime += opts.quantumMS
}

func (p *process) done() bool {
	return p.remaining <= 0
}

type queue struct {
	buf        [100000]process
	l, head, c int
}

func newQueue() *queue {
	return &queue{
		c: 100000,
	}
}

func (q *queue) size() int {
	return q.l
}

func (q queue) get(i int) *process {
	return &q.buf[(i+q.head)%q.c]
}

func (q *queue) set(i int, p process) {
	q.buf[(i+q.head)%q.c] = p
}

func (q *queue) enqueue(p process) {
	q.buf[(q.l+q.head)%q.c] = p
	q.l++
}

func (q *queue) dequeue() process {
	p := q.buf[(q.head)%q.c]
	q.l--
	q.head++
	return p
}

func printRes() {
	for _, p := range processed {
		fmt.Printf("%v %v\n", p.name, p.endAt)
	}
}
