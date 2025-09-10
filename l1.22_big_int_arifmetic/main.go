package main

import (
	"fmt"
	"math/big"
)

type Arithmetic interface {
	Add(a, b string) string
	Sub(a, b string) string
	Mul(a, b string) string
	Div(a, b string) string
}

// math/big
type BigIntArithmetic struct{}

func (BigIntArithmetic) Add(a, b string) string {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)
	return new(big.Int).Add(x, y).String()
}

func (BigIntArithmetic) Sub(a, b string) string {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)
	return new(big.Int).Sub(x, y).String()
}

func (BigIntArithmetic) Mul(a, b string) string {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)
	return new(big.Int).Mul(x, y).String()
}

func (BigIntArithmetic) Div(a, b string) string {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString(a, 10)
	y.SetString(b, 10)
	return new(big.Int).Div(x, y).String()
}

// bitwise arithmetic
type BitwiseArithmetic struct{}

func bitwiseAddInt(a, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func bitwiseSubInt(a, b int) int {
	return bitwiseAddInt(a, bitwiseAddInt(^b, 1))
}

func bitwiseMulInt(a, b int) int {
	res := 0
	for b > 0 {
		if (b & 1) == 1 {
			res = bitwiseAddInt(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

func bitwiseDivInt(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	quotient := 0
	temp := 0
	for i := 31; i >= 0; i-- {
		if (temp + (b << i)) <= a {
			temp = bitwiseAddInt(temp, b<<i)
			quotient |= (1 << i)
		}
	}
	return quotient
}

func (BitwiseArithmetic) Add(a, b string) string {
	var ai, bi int
	fmt.Sscan(a, &ai)
	fmt.Sscan(b, &bi)
	return fmt.Sprint(bitwiseAddInt(ai, bi))
}

func (BitwiseArithmetic) Sub(a, b string) string {
	var ai, bi int
	fmt.Sscan(a, &ai)
	fmt.Sscan(b, &bi)
	return fmt.Sprint(bitwiseSubInt(ai, bi))
}

func (BitwiseArithmetic) Mul(a, b string) string {
	var ai, bi int
	fmt.Sscan(a, &ai)
	fmt.Sscan(b, &bi)
	return fmt.Sprint(bitwiseMulInt(ai, bi))
}

func (BitwiseArithmetic) Div(a, b string) string {
	var ai, bi int
	fmt.Sscan(a, &ai)
	fmt.Sscan(b, &bi)
	return fmt.Sprint(bitwiseDivInt(ai, bi))
}

// adapter implementations
type ArithmeticSystem struct {
	impl Arithmetic
}

func (s *ArithmeticSystem) SetImplementation(impl Arithmetic) {
	s.impl = impl
}

func (s *ArithmeticSystem) Calculate(a, b string) {
	fmt.Println("a + b =", s.impl.Add(a, b))
	fmt.Println("a - b =", s.impl.Sub(a, b))
	fmt.Println("a * b =", s.impl.Mul(a, b))
	fmt.Println("a / b =", s.impl.Div(a, b))
}

func main() {
	a := "1048577" // > 2^20
	b := "2097153" // > 2^21

	system := &ArithmeticSystem{}

	system.SetImplementation(BigIntArithmetic{})
	fmt.Println("--- math/big ---")
	system.Calculate(a, b)

	system.SetImplementation(BitwiseArithmetic{})
	fmt.Println("--- bitwise ---")
	system.Calculate(a, b)
}
