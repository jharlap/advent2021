package main

import (
	"testing"
)

func TestReport_OxygenGeneratorRating(t *testing.T) {
	tests := map[string]struct {
		values []string
		want   int
	}{
		"ex": {
			values: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want:   23,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := NewReport(len(tt.values[0]))
			for _, v := range tt.values {
				r.Add(v)
			}

			if got := r.OxygenGeneratorRating(); got != tt.want {
				t.Errorf("Report.LifeSupportRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_CO2ScrubberRating(t *testing.T) {
	tests := map[string]struct {
		values []string
		want   int
	}{
		"ex": {
			values: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want:   10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := NewReport(len(tt.values[0]))
			for _, v := range tt.values {
				r.Add(v)
			}

			if got := r.CO2ScrubberRating(); got != tt.want {
				t.Errorf("Report.LifeSupportRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
