package main

import ("fmt"
		"math"
		"math/rand")


// package level variables
var g_var1, g_var2 int

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func addSub(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y

	return //this is called a naked return
	       //Naked return statements should be used only in short functions.
	       //They can harm readability in longer functions.
}


func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("The square root of 4 is", math.Sqrt(4))	

	fmt.Println("A number from 1-100 is", rand.Intn(100))	
	
	fmt.Println("Pi is", math.Pi)

	fmt.Println("3 + 4 =", add(3, 4))

	fmt.Println("4 + 5 =", add2(4, 5))

	sum, sub := addSub(5, 6)
	fmt.Println("5 + 6 =", sum)
	fmt.Println("5 - 6 =", sub)

	a, b := swap("Zhuo", "Shen")
	fmt.Println("Zhuo Shen swapped to", a, b)

	//function level variable
	var i, j, k int // now i j k are all inited to be 0
	var i1, j1, k1 int = 1, 2, 3
	fmt.Println("i, j, k", i, j, k)
	fmt.Println("i1, j1, k1", i1, j1, k1)

	var int_var, bool_var, str_var = 1, false, "hello"
	int_var1, bool_var1, str_var1 := 1, false, "hello"
	fmt.Println("int_var, bool_var, str_var", int_var, bool_var, str_var)
	fmt.Println("int_var1, bool_var1, str_var1", int_var1, bool_var1, str_var1)
	
	// := this is called short assignment, used in place 
	// of a var declaration with implicit type
	// note := cannot be used wiht a statement start with
	// a key word(var, func)

	//Basic data types
	//bool
	//string
	//int, int8, int16, int32, int64 
	//uint, uint8, uint16, uint32, uint64
	//int and uint are 32bit/64bit system specific
	//uintptr
	//byte = uint8
	//rune = int32 represents a Unicode code point
	//float32, float64
	//complex64, complex128

	fmt.Printf("Type: %T Value: %v\n", i1, i1)

	//by default uninited int and float is 0, string is "", bool is false

	//all type conversion are explicit
	var fromVar float32 = 3.5
	toVar := int(fromVar) //round down
	fmt.Println("Type conversion:", toVar) // toVar = 3

	var fromVar1 int = 4
	toVar1 := float32(fromVar1) //explicit
	fmt.Println("Type conversion:", toVar1) 

	//Type inference (untyped)
	i2 := 42        //int
	j2 := 3.14      //float
	k2 := 0.5 + 3i  //complex128
	fmt.Println("i2, j2, k2", i2, j2, k2)

	//Type inference (untyped)
	var i3 int = 5
	j3 := i3 // j3 is int
	fmt.Println("i3, j3", i3, j3)

	//Constant
	const PI = 3.14
	const NAME = "Joe Shen"
	//constant cannot be declared using :=

	const (
		// here we dont use := and it is untyped
		Big = 1 << 100
		//Big is not stored as int64, as int64 cannot hold that big number
		//my understanding is Big is holding the result of 1 << 100
		Small = Big >> 99
	)
	//btw, you dont have to use consant var
	//no reference to Small and it is OK
	fmt.Println("Big right shift 90", Big >> 90)


}	
