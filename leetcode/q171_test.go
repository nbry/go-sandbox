package leetcode

import (
	"testing"
)

func TestLeetcode171Example1(t *testing.T) {
	r := TitleToNumber("A")
	e := 1

	if r != e {
		t.Errorf("LC171 - Did not pass base example #1")
	}
}

func TestLeetcode171Example2(t *testing.T) {
	r := TitleToNumber("Z")
	e := 26

	if r != e {
		t.Errorf("LC171 - Did not pass base example #2")
	}
}

func TestLeetcode171Example3(t *testing.T) {
	r := TitleToNumber("ZY")
	e := 701

	if r != e {
		t.Errorf("LC171 - Did not pass base example #3")
	}
}
