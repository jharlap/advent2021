package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Accumulator struct {
	zeros, ones int
}

func (a *Accumulator) Add(bit rune) {
	switch bit {
	case '0':
		a.zeros++
	case '1':
		a.ones++
	}
}

func (a Accumulator) GammaBit() rune {
	if a.ones > a.zeros {
		return '1'
	}
	return '0'
}

func (a Accumulator) EpsilonBit() rune {
	if a.ones > a.zeros {
		return '0'
	}
	return '1'
}

type Report []Accumulator

func NewReport(bits int) Report {
	var r Report
	for i := 0; i < bits; i++ {
		r = append(r, Accumulator{})
	}
	return r
}

func (r Report) Add(row string) {
	for i, c := range []rune(row) {
		r[i].Add(c)
	}
}

func (r Report) GammaString() string {
	var rr []rune
	for _, c := range r {
		rr = append(rr, c.GammaBit())
	}
	return string(rr)
}

func (r Report) EpsilonString() string {
	var rr []rune
	for _, c := range r {
		rr = append(rr, c.EpsilonBit())
	}
	return string(rr)
}

func BitStrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		errExit("unexpected error parsing bitstring", err)
	}
	return i
}

func main() {
	var r Report
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()
		if r == nil {
			r = NewReport(len(l))
		}

		r.Add(l)
	}
	if err := s.Err(); err != nil {
		errExit("unexpected scanner error", err)
	}

	fmt.Println("gamma:", r.GammaString())
	fmt.Println("epsilon:", r.EpsilonString())
	fmt.Println("product:", BitStrToInt64(r.GammaString())*BitStrToInt64(r.EpsilonString()))
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}
