package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	// card is the numerical cells of the bingo card ordered as row1, row2, ..., row5
	card []int

	// matches is a sequence of 0 or 1 runes that match the ordering of the card
	matches []rune
}

func NewBingo() *Bingo {
	return &Bingo{
		matches: []rune(strings.Repeat("0", 25)),
	}
}

func (b *Bingo) Reset() {
	b.matches = []rune(strings.Repeat("0", 25))
}

func (b *Bingo) Add(n int) {
	b.card = append(b.card, n)
}

func (b *Bingo) Call(n int) {
	for i := 0; i < 25; i++ {
		if b.card[i] == n {
			b.matches[i] = '1'
		}
	}
}

const win = "11111"

func (b Bingo) IsWinner() bool {
	// check rows
	for i := 0; i < 25; i += 5 {
		if string(b.matches[i:i+5]) == win {
			return true
		}
	}

	// check columns
	if b.matchesForIndices([]int{0, 5, 10, 15, 20}) == win ||
		b.matchesForIndices([]int{1, 6, 11, 16, 21}) == win ||
		b.matchesForIndices([]int{2, 7, 12, 17, 22}) == win ||
		b.matchesForIndices([]int{3, 8, 13, 18, 23}) == win ||
		b.matchesForIndices([]int{4, 9, 14, 19, 24}) == win {
		return true
	}

	/*
		0  1  2  3  4
		5  6  7  8  9
		10 11 12 13 14
		15 16 17 18 19
		20 21 22 23 24
	*/

	// check diagonals
	/*
		if b.matchesForIndices([]int{0, 6, 12, 18, 24}) == win ||
			b.matchesForIndices([]int{4, 8, 12, 16, 20}) == win {
			return true
		}
	*/
	return false
}

func (b Bingo) matchesForIndices(ii []int) string {
	var rr []rune
	for _, i := range ii {
		rr = append(rr, b.matches[i])
	}
	return string(rr)
}

func (b Bingo) SumUnmarked() int {
	var s int
	for i, m := range b.matches {
		if m == '0' {
			s += b.card[i]
		}
	}
	return s
}

func (b Bingo) Copy() Bingo {
	return Bingo{
		card:    b.card[:],
		matches: b.matches[:],
	}
}

func (b Bingo) String() string {
	var sb strings.Builder
	for i := range b.matches {
		if i%5 == 0 {
			sb.WriteString("\n")
		}
		mm := " "
		if b.matches[i] == '1' {
			mm = "*"
		}
		sb.WriteString(fmt.Sprintf("%s% 2d  ", mm, b.card[i]))
	}

	return sb.String()
}

func score1(nums []int, boards []*Bingo) int {
	for _, i := range nums {
		for _, b := range boards {
			b.Call(i)
			if b.IsWinner() {
				return b.SumUnmarked() * i
			}
		}
	}
	return 0
}

func score2(nums []int, boards []*Bingo) int {
	var (
		lastScore int
	)

	for _, i := range nums {
		for _, b := range boards {
			if !b.IsWinner() {
				b.Call(i)
				if b.IsWinner() {
					lastScore = b.SumUnmarked() * i
				}
			}
		}
	}
	return lastScore
}

func parseInput(in io.Reader) ([]int, []*Bingo) {
	var (
		nums     []int
		boards   []*Bingo
		curBoard int = -1
	)

	s := bufio.NewScanner(bufio.NewReader(in))
	for s.Scan() {
		l := s.Text()
		if nums == nil {
			vv := strings.Split(l, ",")
			for _, v := range vv {
				i, err := strconv.Atoi(v)
				if err != nil {
					errExit("could not parse nums row", err)
				}
				nums = append(nums, i)
			}
			continue
		}

		if len(l) == 0 {
			curBoard++
			boards = append(boards, NewBingo())
			continue
		}

		// add a row to a board
		vv := strings.Split(l, " ")
		for _, v := range vv {
			v = strings.TrimSpace(v)
			if len(v) == 0 {
				continue
			}

			i, err := strconv.Atoi(v)
			if err != nil {
				errExit("could not parse bingo row '%s'", err, l)
			}
			boards[curBoard].Add(i)
		}
	}
	if err := s.Err(); err != nil {
		errExit("unexpected scanner error", err)
	}

	return nums, boards
}

func main() {
	nums, boards := parseInput(os.Stdin)
	fmt.Println("score 1:", score1(nums, boards))
	fmt.Println("score 2:", score2(nums, boards))
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}
