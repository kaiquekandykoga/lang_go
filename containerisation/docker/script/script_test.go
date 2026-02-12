package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	got := calc()
	want := 2
	if got != want {
		t.Errorf("calc() = %d; want %d", got, want)
	}
}
