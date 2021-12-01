package main

import (
	"testing"
)

func SlidingWindowTest(t *testing.T) {
	sw := NewSlidingWindowInt(3)
	sw.Append(1)
	sw.Append(2)
	sw.Append(3)
	if sw.Sum() != 6 {
		t.Fatalf("expected 6 but got %d\n", sw.Sum())
	}
	sw.Append(3)
	if sw.Sum() != 9 {
		t.Fatalf("expected 9 but got %d\n", sw.Sum())
	}
}
