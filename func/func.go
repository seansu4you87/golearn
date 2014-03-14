package main

import (
	"fmt"
	"math/rand"
	"time"
)

type binFunc func(int, int) int
type loudBinFunc func(int, int) int

func add(x, y int) int {
	return x + y
}

// any type that has an Error() string method is a valid error type
func (fn binFunc) Error() string {
	return "binFunc error"
}

func (fn loudBinFunc) Error() string {
	return "THIS ERROR IS A LOT LOUDER"
}

type op struct {
	name string
	fn   func(int, int) int
}

type walkFn func(*int) walkFn

func pickRandom(fns ...walkFn) walkFn {
	return fns[rand.Intn(len(fns))]
}

func walkEqual(i *int) walkFn {
	*i += rand.Intn(7) - 3
	return pickRandom(walkForward, walkBackward)
}

func walkForward(i *int) walkFn {
	*i += rand.Intn(6)
	return pickRandom(walkEqual, walkBackward)
}

func walkBackward(i *int) walkFn {
	*i += -rand.Intn(6)
	return pickRandom(walkEqual, walkForward)
}

func main() {
	fmt.Println("Seeding rand with time")
	rand.Seed(time.Now().Unix())

	// Go has closures!
	/*
		x := 5
		fn := func() {
			fmt.Println("x is:", x)
		}

		fn()
		x++
		fn()
	*/

	// Go can alias a function type
	/*
		fns := []binFunc{
			func(x, y int) int { return x + y },
			func(x, y int) int { return x - y },
			func(x, y int) int { return x * y },
			func(x, y int) int { return x / y },
			func(x, y int) int { return x % y },
		}

		fn := fns[rand.Intn(len(fns))]

		x, y := 12, 5
		fmt.Println(fn(x, y))
	*/

	// Go can use functions as fields
	/*
		ops := []op{
			{"add", func(x, y int) int { return x + y }},
			{"sub", func(x, y int) int { return x - y }},
			{"mul", func(x, y int) int { return x * y }},
			{"div", func(x, y int) int { return x / y }},
			{"mod", func(x, y int) int { return x % y }},
		}

		o := ops[rand.Intn(len(ops))]
		x, y := 12, 5
		fmt.Println(o.name, x, y)
		fmt.Println(o.fn(x, y))
	*/

	//Go has recursive function types
	/*
		fn, progress := walkEqual, 0
		for i := 0; i < 20; i++ {
			fn = fn(&progress)
			fmt.Println(progress)
		}
	*/

	//Go can use function types as interface values
	/*
		var err error
		err = binFunc(add)
		fmt.Println(err)
	*/
}
