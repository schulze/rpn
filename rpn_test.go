package rpn

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	c := New()
	c.Push(big.NewInt(4))
	c.Push(big.NewInt(2))
	c.Push(big.NewInt(3))
	c.Add()
	got := c.Top()
	if got.Int64() != 5 {
		t.Fatal("Expected 5, but got", got)
	}

	c.Add()
	got = c.Pop()
	if got.Int64() != 9 {
		t.Fatal("Expected 9, but got", got)
	}
}
