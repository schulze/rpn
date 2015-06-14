package rpn

import (
	"math/big"
)

type Calc struct {
	stack []*big.Int
}

func New() *Calc {
	c := Calc{}
	c.stack = make([]*big.Int, 0)
	return &c
}

func (c *Calc) Pop() *big.Int {
	n := len(c.stack)
	top := c.stack[n-1]
	c.stack = c.stack[:n-1]
	return top
}

func (c *Calc) Top() *big.Int {
	n := len(c.stack)
	return c.stack[n-1]
}

func (c *Calc) Push(x *big.Int) int {
	c.stack = append(c.stack, x)
	return len(c.stack)
}

type BinOp func(r, x, y *big.Int) *big.Int

func (c *Calc) Apply(f BinOp) *big.Int {
	y := c.Pop()
	x := c.Top()
	f(x, x, y)
	return x
}

func (c *Calc) Add() *big.Int {
	return c.Apply((*big.Int).Add)
}

func (c *Calc) Sub() *big.Int {
	return c.Apply((*big.Int).Sub)
}

func (c *Calc) Mul() *big.Int {
	return c.Apply((*big.Int).Mul)
}

func (c *Calc) Div() *big.Int {
	return c.Apply((*big.Int).Div)
}

func (c *Calc) Mod() *big.Int {
	return c.Apply((*big.Int).Mod)
}

func (c *Calc) ModInverse() *big.Int {
	return c.Apply((*big.Int).ModInverse)
}
