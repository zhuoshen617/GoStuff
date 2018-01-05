package main 

import "fmt"

// this is a struct, a collection of fields
type Vertex struct {
    X int
    Y string
}


func main() {
	
	var p *int // p is a pointer to int, init to nil

	a := 5
	p = &a // p points to a
	q := &a // q points to a using short assignment

	fmt.Println(a)
	fmt.Println(*p) //deference
	fmt.Println(*q) 

	//Go has not pointer arithmetic

	fmt.Println(Vertex{1, "xyz"})

	v := Vertex{1, "xyz"}
	v.X = 2
	fmt.Println(v)

	v_ptr := &v //pointer
	v_ptr.X = 10 //this is same as (*v_ptr).X
				 //Go thinks this is cumbersome and allow implicit dereference
	fmt.Println(v)


	var (
		v1 = Vertex{1, "xyz"}
		v2 = Vertex{X: 1}
		v3 = Vertex{Y: "xyz"}
		v4 = Vertex{X: 1, Y: "xyz"}
		v5 = Vertex{}
		ptr = &Vertex{1, "xyz"}
	)

	fmt.Println()
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(*ptr)
}