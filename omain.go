package main

import (
	"fmt"
	"math"
)

func main() {
	// fn_fors()
	// fn_array_slice()
	// fn_if()
	// fn_maps()
	fn_switch()
}

func fn_switch() {
	// rating := 3
	switch rating := 3; rating {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two or Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Default case")
	}

	temp := -5
	switch {
	case temp < 0:
		fmt.Println("It's cold")
	case temp == 0:
		fmt.Println("It's right at zero")
	case temp > 20:
		fmt.Println("It's moderate tmp")
	}
}

func fn_maps() {
	// var score map[string]int
	score := make(map[string]int)
	score["rodo"] = 89
	score["sofia"] = 100
	score["juan"] = 53
	score["vicky"] = 45
	fmt.Println(score)

	delete(score, "vicky")
	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}

	_, found := score["k2"]
	fmt.Println("found:", found)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func fn_if() {
	age := 19

	if age >= 18 {
		fmt.Println("You can ride!")
	} else if age >= 14 {
		fmt.Println("You can ride with a parent!")
	} else {
		fmt.Println("You can not ride")
	}

	var x, n, lim float64
	x, n, lim = 3, 2, 10
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
}

func fn_array_slice() {
	// Arrays
	var x [5]int
	fmt.Printf("%T\n", x)

	x[4] = 4

	fmt.Println(x)
	fmt.Println(x[2:])

	// Slices

	s := make([]int, 5)
	fmt.Printf("%T\n", s)

	s[4] = 4
	s = append(s, 5)

	fmt.Println(s)
	fmt.Println(s[4:])
}

func fn_fors() {
	things := []string{"one", "two", "three", "four"}
	fmt.Println(things)

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i, thing := range things {
		if i == 1 {
			continue
		}

		if i == 3 {
			break
		}

		fmt.Println(i, thing)
	}

	start := 1
	for start < 100 {
		start += start
		fmt.Println(start)
	}
}
