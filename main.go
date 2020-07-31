package main

import (
	"fmt"
	"math"
)

type IFunctorInt interface {
	fmapI(fn func(int) int) IFunctorInt
	fmapS(fn func(int) string) IFunctorString
}

type IFunctorString interface {
	fmapI(fn func(string) int) IFunctorInt
	fmapS(fn func(string) string) IFunctorString
}

type OptionalInt struct {
	val int
}

type OptionalString struct {
	val string
}

func OptinalIntOf(value int) OptionalInt {
	return OptionalInt{value}
}

func OptionalStringOf(value string) OptionalString {
	return OptionalString{value}
}

func (f OptionalInt) fmapI(fn func(int) int) IFunctorInt {
	if f.val == 0 {
		return OptionalInt{0}
	}

	return OptionalInt{fn(f.val)}
}

func (f OptionalInt) fmapS(fn func(int) string) IFunctorString {
	if f.val == 0 {
		return OptionalString{""}
	}

	return OptionalString{fn(f.val)}
}

func (f OptionalString) fmapI(fn func(string) int) IFunctorInt {
	if f.val == "" {
		return OptionalInt{0}
	}

	return OptionalInt{fn(f.val)}
}

func (f OptionalString) fmapS(fn func(string) string) IFunctorString {
	if f.val == "" {
		return OptionalString{""}
	}

	return OptionalString{fn(f.val)}
}

type NonNegativeInt struct {
	val int
}

type NonNegativeString struct {
	val string
}

func NonNegativeIntOf(value int) NonNegativeInt {
	return NonNegativeInt{value}
}

func NonNegativeIntString(value string) NonNegativeString {
	return NonNegativeString{value}
}

func (f NonNegativeInt) fmapI(fn func(int) int) IFunctorInt {
	if f.val < 0 {
		return NonNegativeInt{0}
	}

	return NonNegativeInt{fn(f.val)}
}

func (f NonNegativeInt) fmapS(fn func(int) string) IFunctorString {
	if f.val < 0 {
		return NonNegativeString{""}
	}

	return NonNegativeString{fn(f.val)}
}

func (f NonNegativeString) fmapI(fn func(string) int) IFunctorInt {
	if len(f.val) == 0 {
		return NonNegativeInt{0}
	}

	return NonNegativeInt{fn(f.val)}
}

func (f NonNegativeString) fmapS(fn func(string) string) IFunctorString {
	if len(f.val) == 0 {
		return NonNegativeString{""}
	}

	return NonNegativeString{fn(f.val)}
}

func divide100By(dividend int) int {
	return 100 / dividend
}

func minusTen(v int) int {
	return v - 10
}

func recToA(v int, acc string) string {
	if v == 0 {
		return acc
	}

	return recToA(int(math.Abs(float64(v-1))), acc+"a")
}

func toA(v int) string {
	// look up to implement fixpoint so we can rec inside fn
	//rec := func(v int, acc string) string {
	//	if v == 0 {
	//		return acc
	//	}

	//	return rec(v-1, acc+"a")
	//}

	return recToA(v, "")
}

func lem(v string) int {
	return len(v)
}

func main() {
	v := OptinalIntOf(50).
		fmapI(divide100By).
		fmapI(minusTen).
		fmapS(toA).
		fmapI(lem)

	fmt.Printf("%v", v)
}
