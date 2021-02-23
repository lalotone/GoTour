package main

// Package name is the same as the last element of the import path
// fmt -> the package as itself
// path math/ package rand
// packages could be imported on this way (good one). This import way is the factored one:
import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

// or
// import "fmt"

// Variables defined at package level.
// Keyword val + list of variables + type at the end, like functions
var foo, bar, baz bool
var foobar int

// Variables could be initialized with values. the type is not required, it will be taken from
// the value that initializes the variable
var foobax = "Foobax"

/* Same factored-way to create variables:
The valid var types are:
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

The int, uint are by default 32 bits wide on 32 bits system, and 64 bits on 64 bits OS
*/
var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Variables without initial value are started with the Zero value to each type:
// false for bool
// 0 for numeric vars
// "" for strings
var (
	aZeroNumber int
	aZeroString string
	aZeroBool   bool
	aZeroFloat  int
)

// Constant variable declared at package level
const Pi = 3.14

// Constants declared in a factored way
// Value MUST be provided
const (
	apples   = 10
	cherries = 40
)

/*
Function, composition of:
sum -> name
(x int, y int) -> Parameters
int -> The type (return value will be an int). Should be at the end.
*/
func sum(x int, y int) int {
	return x + y
}

// Same as above, but using only one type declaration on parameters. shortened way of function.
func sum_v2(x, y int) int {
	return x + y
}

// Function that returns multiple values
func multiple_returns(foo, bar string) (string, string) {
	return foo, bar
}

// Function with named return values
// Return values must be different from param names
func return_named(a, b int) (c, d int) {
	// Assigning values to variables.
	c = 30
	d = 10
	// Naked return statement, bcoz the return params are defined on the function itself
	return
}

func main() {
	// Variables at function level instead of package:
	var foobaz int
	var a, b, c = 1, 2, 3
	// CAUTION, the short assignment could be used only INSIDE FUNCIONS!
	// Outside from functions, the var keyword must be used, short mode not available.
	yz := 30
	// Constant variable declared on function, CANNOT be declared with the := syntax
	const theTruth = true
	// the Println from package fmt is an exported name. When we import packages, only the exported names are
	// accesible from the "outside" of the package.
	// Same for the Intn name, if we use intn insted, it will fail.
	fmt.Println("Hello world")
	// Always same number returned. Need to seed before giving another number
	fmt.Println(rand.Intn(10))
	fmt.Println(sum(4, 5))
	fmt.Println(sum_v2(5, 4))
	fmt.Println("Foo", "Bar")
	fmt.Println(10, 20)
	fmt.Println(foo, bar, baz)
	fmt.Println(foobar)
	fmt.Println(foobaz)
	fmt.Println(foobax)
	fmt.Println(a, b, c)
	fmt.Println(yz)
	fmt.Println(return_named(1, 1))
	// Print Type and value of variables:
	fmt.Printf("type: %T Value: %v\n", ToBe, ToBe)
	// %q	a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%v %q %v %v \n", aZeroNumber, aZeroString, aZeroBool, aZeroFloat)
	// TYPE CONVERTIONS
	/*
		In Go the way to "cast" types is:
		T(v) -> Type(value)
	*/
	var i int = 23
	var f float64 = float64(i)
	// And this will fail if uncomment
	// var f float64 = i
	fmt.Println(i, f)

	// Variables whose Type is not declared, Type will be infered from the value of the right:
	bartolo := 12
	fmt.Printf("%T with value %v \n", bartolo, bartolo)
	barbaz := "Foobar"
	fmt.Printf("%T with value %v \n", barbaz, barbaz)
	fmt.Printf("%v %v", apples, cherries)
}
