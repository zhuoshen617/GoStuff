package main 

import ("fmt"
		"strings"
)

func main() {

	//array is always fixed-sized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//this is an array
	primes := [6]int{2,3,5,7,11,13}

	//this is a slice
	//primes := []int{2,3,5,7,11,13}
	fmt.Println(primes)


	//this is called slice, [low:high)
	//primes is the underlying array
	var s1 []int = primes[1:4]
	s2 := primes[2:5]
	fmt.Println(s1)
	fmt.Println(s2)

	//slice is like a range reference
	s1[1] = 99 //change primes and affect s1, s2 also
	fmt.Println(s1) //[3 99 7]
	fmt.Println(s2) //[99 7 11]
	fmt.Println(primes) //[2 3 99 7 11 13]

	//s3 is the slice, but there is an underlying array underlying
	s3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
	}

	fmt.Println(s3)
	
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// Slice the slice to give it zero length.
	s4 = s4[:0]
	// cap counting from the first element from the slice onward
	printSlice(s4) // len(s4) = 0, cap(s4) = 6 

	// Extend its length.
	s4 = s4[:4]
	printSlice(s4) // len(s4) = 4, cap(s4) = 6 

	// Drop its first two values.
	s4 = s4[2:]
	//here cap(s4) = 4 and cannot get back the first two element
	printSlice(s4) // len(s4) = 2, cap(s4) = 4  

	//use make to create dynamically-sized arrays
	s6 := make([]int, 5)
	printSlice(s6)
	//len = 5, cap = 5, [0,0,0,0,0]
	
	s7 := make([]int, 0, 5)
	printSlice(s7)
	//len = 0, cap = 5, []

	s8 := s6[ : 2]
	printSlice(s8)
	//len = 2, cap = 5, [0,0]

	s9 := s6[2 : 5]
	printSlice(s9)
	//len = 3, cap = 3, [0,0,0]


	//slice of slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	//a nil slice
	var s5 []int
	printSlice(s5) // s = nil, len(s) = 0, cap(s) = ? 

	s5 = append(s5, 0)
	printSlice(s5) // s = nil, len(s) = 0, cap(s) = ? 

	s5 = append(s5, 1)
	printSlice(s5) // s = nil, len(s) = 0, cap(s) = ? 

	s5 = append(s5, 2, 3, 4)
	printSlice(s5) // s = nil, len(s) = 0, cap(s) = ? 

	//If the backing array of s is too small to fit,
	//a bigger array will be allocated. 
	//The returned slice will point to the newly allocated array.
	//In this case the cap(s) is different on each machine.

	for i, v := range s5{
		//i is index, v is a COPY of value
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range s5{
		//i is index
		fmt.Println(i)
	}

	for _, v := range s5{
		//v is a COPY of value
		fmt.Println(v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}