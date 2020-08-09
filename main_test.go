package main

import (
	"testing"
)

func TestFun(t *testing.T) {

	push := []int{1, 2, 3, 4, 5}
	pop := []int{4, 5, 3, 2, 1}

	out := validateStackSequences(push, pop)

	if out != true {
		t.Errorf("got %t, want %t", out, true)
	}

}
