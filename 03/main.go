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

func (a Accumulator) CommonBit() rune {
	if a.ones >= a.zeros {
		return '1'
	}
	return '0'
}

func (a Accumulator) UncommonBit() rune {
	if a.ones < a.zeros {
		return '1'
	}
	return '0'
}

type Report struct {
	acc    []Accumulator
	values []string
}

func NewReport(bits int) Report {
	var r Report
	for i := 0; i < bits; i++ {
		r.acc = append(r.acc, Accumulator{})
	}
	return r
}

func (r *Report) Add(row string) {
	for i, c := range []rune(row) {
		r.acc[i].Add(c)
	}
	r.values = append(r.values, row)
}

func (r Report) GammaString() string {
	var rr []rune
	for _, c := range r.acc {
		rr = append(rr, c.CommonBit())
	}
	return string(rr)
}

func (r Report) EpsilonString() string {
	var rr []rune
	for _, c := range r.acc {
		rr = append(rr, c.UncommonBit())
	}
	return string(rr)
}

func (r Report) Copy() Report {
	res := NewReport(len(r.acc))
	for _, v := range r.values {
		res.Add(v)
	}
	return res
}

func (r Report) Filter(f func(string) bool) Report {
	res := NewReport(len(r.acc))
	for _, v := range r.values {
		if f(v) {
			res.Add(v)
		}
	}
	return res
}

func (r Report) Len() int {
	return len(r.values)
}

func (r Report) OxygenGeneratorRating() int {
	cur := r.Copy()
	for i := range r.acc {
		b := cur.acc[i].CommonBit()
		cur = cur.Filter(func(v string) bool {
			return []rune(v)[i] == b
		})
		if cur.Len() == 1 {
			break
		}
	}
	s := cur.values[0]
	return int(BitStrToInt64(s))
}

func (r Report) CO2ScrubberRating() int {
	cur := r.Copy()
	for i := range r.acc {
		b := cur.acc[i].UncommonBit()
		cur = cur.Filter(func(v string) bool {
			return []rune(v)[i] == b
		})
		if cur.Len() == 1 {
			break
		}
	}
	s := cur.values[0]
	return int(BitStrToInt64(s))
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
		if r.acc == nil {
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

	fmt.Println("life support rating:", r.CO2ScrubberRating()*r.OxygenGeneratorRating())
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}
