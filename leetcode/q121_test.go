package leetcode

import (
	"testing"
)

func TestLeetcode121Example1(t *testing.T) {
	r := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	e := 5

	if r != e {
		t.Errorf("LC121 - Did not pass base example #1")
	}
}

func TestLeetcode121Example2(t *testing.T) {
	r := MaxProfit([]int{7, 6, 4, 3, 1})
	e := 0

	if r != e {
		t.Errorf("LC121 - did not pass base example #2")
	}
}
