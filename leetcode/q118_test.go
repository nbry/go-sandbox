package leetcode

import (
	"reflect"
	"testing"
)

func TestLeetcode118Example1(t *testing.T) {
	r := GeneratePascalTriangle(5)
	e := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC118 - Did not pass base example #1")
	}
}

func TestLeetcode118Example2(t *testing.T) {
	r := GeneratePascalTriangle(1)
	e := [][]int{{1}}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC118 - did not pass base example #2")
	}
}
