package main

import "testing"

func Test_countIncreases(t *testing.T) {
	tests := map[string]struct {
		ii   []int
		want int
	}{
		"empty":     {ii: nil, want: 0},
		"one entry": {ii: []int{1}, want: 0},
		"increased": {ii: []int{1, 100}, want: 1},
		"decreased": {ii: []int{1000, 100}, want: 0},
		"bouncy":    {ii: []int{1, 10, 2, 3, 100}, want: 3},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := countIncreases(tt.ii); got != tt.want {
				t.Errorf("countIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWindowedIncreases(t *testing.T) {
	tests := map[string]struct {
		ii   []int
		want int
	}{
		"empty":         {ii: nil, want: 0},
		"one entry":     {ii: []int{1}, want: 0},
		"two entries":   {ii: []int{1, 2}, want: 0},
		"three entries": {ii: []int{1, 2, 3}, want: 0},
		"increased":     {ii: []int{1, 2, 3, 2}, want: 1},
		"decreased":     {ii: []int{1, 2, 3, 0}, want: 0},
		"bouncy":        {ii: []int{1, 10, 2, 3, 100, 0}, want: 2},
		"example":       {ii: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, want: 5},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := countWindowedIncreases(tt.ii); got != tt.want {
				t.Errorf("countWindowedIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}
