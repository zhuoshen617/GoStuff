package main

import ("fmt"
		"math"
		"time"
)
		


func sqrt(x float64) string {
	// no () is need
	// {} is required
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		//Sprint return a string
		return fmt.Sprint(math.Sqrt(x))
	}


	// there is also a way to do a short statement
	// kind of like for
	// only SHORT assignment/statement is allowed
	// only ONE short assignment/statement is allowed

	if v := x * x; v < 10 {
		//do something
		return ""
	}

	return ""
}

func main() {
	
	fmt.Println("Hello World!")

	sum := 0

	//{} is always required
	//i is at the scope of for
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum from 0 to 9", sum)

	// init statement can be empty
	// post statement can be empty
	// conditiion statement can NOT be em

	// in case both init and post statement is empty
	// where for is just like while
	j := 1
	for j < 10 {
		sum += j
		j++
	}

	//this is an infinite loop
	//for {
	//}

	fmt.Println(sqrt(2), sqrt(-4))

	//switch with no condition can be used as if-elseif-else
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
			// break will be added at the end automatically
			break // this is optional
		case t.Hour() < 17:
			fmt.Println("Good afternoon!")

		default:
			fmt.Println("Good evening!")
	}
	
	// we can do a short assignment here just like for/if
	//today := time.Now().Weekday()
	//switch time.Saturday{
	switch today := time.Now().Weekday(); time.Saturday{
		case today + 0: // today + 0 == "Saturday"
			fmt.Println("Today!")
		case today + 1:
			fmt.Println("Tomorrow!")
		case today + 2:
			fmt.Println("In two days.")		
		default:
			fmt.Println("Too far away.")
	}

	//defer, panic, recover


	
	
}