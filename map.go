package main 

import ("fmt"
		"math"
)	

type Vertex struct {
	lat, long float64
}

func main() {
	//m1 is a nil map, it has no keys and cannot add keys
	var m map[string]int
	fmt.Println(m)

	var m1 = map[string]int{
		"Bell Labs": 3,
		"Google": 4,
	}
	fmt.Println(m1)

	//use make to dynamically create a map
 	m2 := make(map[string]Vertex)
 	m2["Joe"] = Vertex{1.2, 4.3}

 	fmt.Println(m2["Joe"])


 	var m3 = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    Vertex{37.42202, -122.08408},
	}
	fmt.Println(m3["Joe"])

	//the vertex in value can be omitted
	var m4 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m4["Joe"])


	m5 := make(map[string]int)
	//adding Tom to map
	m5["Tom"] = 3

	//changing "Tom" => 4
	m5["Tom"] = 4

	delete(m5, "Tom")
	fmt.Println("The value:", m5["Tom"])

	v, ok := m5["Tom"]
	//when reference should check first
	fmt.Println("The value:", v, "Present?", ok)


	//function can be passed around or used as value
	foo := func(x, y float64) float64 {
		return x + y
	}

	//use as value
	fmt.Println(foo(3.0,4.0))

	fmt.Println(bar(foo))

	fmt.Println(bar(math.Pow))


	//function as a value
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//function used as a passed in value
func bar(fn func(float64, float64) float64) float64 {
	return fn(10.0, 2.0)
}


//closure
func fibonacci() func() int {
	a := -1
	b := 1
	
	return func () int {
		temp := b
		b = a + b 
		a = temp
		
		return b
		
	}
}