package main

import "fmt"

type cat struct {
	name   string
	age    int
	weight float32
}

func (c cat) sleep() {
	fmt.Println("zzzzz")
}

func (c *cat) eat(food float32) float32 {
	c.weight += food // (*c).weight += food
	return c.weight  // return (*c).weight
}

// This func below does not modify the receiver, only the local cat
// func (c cat) eat(food float32) float32 {
// 	c.weight += food
// 	return c.weight
// }

func main() {
	c := cat{"Oscar", 14, 15.2}
	c.sleep()
	c.eat(0.3) // (&c).eat(0.3)

	fmt.Println(c)

	// p := &c
	// p.sleep()  // (*p).sleep()
	// p.eat(0.3) // p.eat(0.3)
}
