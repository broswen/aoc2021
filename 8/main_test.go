package main

import (
	"fmt"
	"testing"
)

func TestSetSubtract(t *testing.T) {
	one := SetFromString("cg")
	four := SetFromString("cgdf")
	fourMinusOne := four.Subtract(one)
	fmt.Println(one)
	fmt.Println(four)
	fmt.Println(fourMinusOne)
	diff := SetFromString("df")
	if !diff.Equals(fourMinusOne) {
		t.Fatalf("diff should equal fourMinusOne")
	}
}

func TestContains(t *testing.T) {
	one := SetFromString("cg")
	four := SetFromString("cgdf")
	fourMinusOne := four.Subtract(one)
	test := SetFromString("fdcge")
	if !test.Contains(fourMinusOne) {
		t.Fatalf("test should contain fourMinusOne")
	}
}
