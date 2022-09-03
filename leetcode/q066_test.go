package leetcode

import (
	"reflect"
	"testing"
)

func TestLeetcode066Example1(t *testing.T) {
	r := PlusOne([]int{1, 2, 3})
	e := []int{1, 2, 4}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC066 - Did not pass base example #1")
	}
}

func TestLeetcode066Example2(t *testing.T) {
	r := PlusOne([]int{4, 3, 2, 1})
	e := []int{4, 3, 2, 2}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC066 - did not pass base example #2")
	}
}

func TestLeetcode066Example3(t *testing.T) {
	r := PlusOne([]int{9})
	e := []int{1, 0}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC066 - did not pass base example #3")
	}
}

func TestLeetcode066EdgeCase(t *testing.T) {
	r := PlusOne([]int{9, 9, 9})
	e := []int{1, 0, 0, 0}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("LC066 - did not pass base example #3")
	}
}
