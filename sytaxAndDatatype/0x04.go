// Primaries

// bool

// string

// int, int8, int16, int32, int64

// uint, uint8, uint16, uint32, uint64

// uintptr

// float32, float64

// complex64, complex128

// byte (alias for uint8)

// rune (alias for int32)

// compsitive data type 
// arr , arrslice and map
package main 

import (
	"fmt"
)

func main(){
	b := [5]int{1,2,3,4,5}
	c := []int {}
	c = append(c,1)

	type Kilometers int     // Type definition, distinct from int
	type ID = int           // Type alias, identical to int

	var distance Kilometers = 99
	var identity ID = 123
}