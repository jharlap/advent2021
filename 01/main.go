package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var ii []int
	for s.Scan() {
		l := s.Text()
		i, err := strconv.Atoi(l)
		if err != nil {
			errExit("could not parse '%s' as an int", err, l)
		}
		ii = append(ii, i)
	}
	if err := s.Err(); err != nil {
		errExit("unexpected scanner error", err)
	}

	n := countIncreases(ii)
	fmt.Println("increases:", n)

	n = countWindowedIncreases(ii)
	fmt.Println("windowed increases:", n)
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}

func countIncreases(ii []int) int {
	if len(ii) < 2 {
		return 0
	}

	var n int
	for i := 1; i < len(ii); i++ {
		if ii[i] > ii[i-1] {
			n++
		}
	}
	return n
}

func countWindowedIncreases(ii []int) int {
	if len(ii) < 4 {
		return 0
	}

	lastWindow := ii[0] + ii[1] + ii[2]
	var n int
	for i := 3; i < len(ii); i++ {
		w := lastWindow - ii[i-3] + ii[i]
		if w > lastWindow {
			n++
		}
		lastWindow = w
	}
	return n
}
