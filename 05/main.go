package main

import (
	"bufio"
	"fmt"
	"os"
)

type Map [][]int

func NewMap(w, h int) Map {
	m := make([][]int, w+1)
	for i := 0; i <= h; i++ {
		m[i] = make([]int, h+1)
	}
	return m
}

func (m Map) AddLine(x1, y1, x2, y2 int) {
	if y1 == y2 {
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			m[x][y1]++
		}
	} else if x1 == x2 {
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			m[x1][y]++
		}
	} else {

		if x2 < x1 {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}

		var yd int
		if y2 > y1 {
			yd = 1
		} else if y2 < y1 {
			yd = -1
		} else {
			yd = 0
		}

		y := y1
		for x := x1; x <= x2; x++ {
			m[x][y]++
			y += yd
		}
	}
}

func (m Map) Count(filter func(n int) bool) int {
	var n int
	for _, col := range m {
		for _, v := range col {
			if filter(v) {
				n++
			}
		}
	}
	return n
}

func parseLine(l string) []int {
	var x1, y1, x2, y2 int
	fmt.Sscanf(l, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
	return []int{x1, y1, x2, y2}
}

func main() {
	var coords [][]int
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for s.Scan() {
		l := s.Text()
		coords = append(coords, parseLine(l))
	}
	if err := s.Err(); err != nil {
		errExit("unexpected scanner error", err)
	}

	var w, h int
	for _, r := range coords {
		if r[0] > w {
			w = r[0]
		}
		if r[2] > w {
			w = r[2]
		}
		if r[1] > h {
			h = r[1]
		}
		if r[3] > h {
			h = r[3]
		}
	}

	m := NewMap(w, h)
	for _, r := range coords {
		m.AddLine(r[0], r[1], r[2], r[3])
	}

	n := m.Count(func(n int) bool { return n >= 2 })
	fmt.Println("n:", n)
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}
