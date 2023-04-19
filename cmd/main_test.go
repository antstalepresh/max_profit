package main

import (
	"fmt"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		n      int
		prices []int
		want   int
	}{
		{
			n:      2,
			prices: []int{1, 4, 1},
			want:   3,
		},
		{
			n:      2,
			prices: []int{1},
			want:   0,
		},
		{
			n:      2,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   7,
		},
		{
			n:      2,
			prices: []int{6, 1, 3, 2, 4, 7},
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, n: %v", tt.prices, tt.n), func(t *testing.T) {
			if got := maxProfit(tt.n, tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
