package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	horizontal int
	depth      int
	aim        int
}

func (s *Submarine) Move(cmd string) error {
	ww := strings.Split(cmd, " ")
	if len(ww) != 2 {
		return fmt.Errorf("incorrect number of words in a command, expected 2 and got %d", len(ww))
	}

	i, err := strconv.Atoi(ww[1])
	if err != nil {
		return fmt.Errorf("command word 2 expected number, got '%s'", ww[1])
	}

	switch ww[0] {
	case "forward":
		s.horizontal += i
		s.depth += s.aim * i

	case "down":
		s.aim += i

	case "up":
		s.aim -= i
	}
	return nil
}

func (s Submarine) Area() int {
	return s.horizontal * s.depth
}

func main() {
	var sub Submarine
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()
		err := sub.Move(l)
		if err != nil {
			errExit("unexpected command", err)
		}
	}
	if err := s.Err(); err != nil {
		errExit("unexpected scanner error", err)
	}

	fmt.Println("Area:", sub.Area())
}

func errExit(reason string, err error, args ...interface{}) {
	args = append(args, err)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s: %%w", reason), args...)
	os.Exit(1)
}
