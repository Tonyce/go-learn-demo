package main

import "fmt"

// Base ...
type Base struct {
	b   int
	tag string
}

// Container ...
type Container struct {
	Base
	c   string
	tag string
}

// Describe ...
func (base Base) Describe() string {
	return fmt.Sprintf("base %d belongs to us", base.b)
}

// DescribeTag ...
func (base Base) DescribeTag() string {
	return fmt.Sprintf("tag is %s", base.tag)
}

// DescribeTag ...
// func (co Container) DescribeTag() string {
// 	return fmt.Sprintf("tag is %s", co.tag)
// }

// Fooer ...
type Fooer interface {
	Foo() string
}

// Containerfoo ...
type Containerfoo struct {
	Fooer
}

// Foo ...
func (cont Containerfoo) Foo() string {
	return cont.Fooer.Foo()
}

func sink(f Fooer) {
	fmt.Println("sink:", f.Foo())
}

// TheRealFoo is a type that implements the Fooer interface.
type TheRealFoo struct {
}

// Foo ...
func (trf TheRealFoo) Foo() string {
	return "TheRealFoo Foo"
}

func main() {

	// cc := Container{}
	// cc.b = 1
	// cc.c = "string"
	// fmt.Printf("co -> { b: %v, c: %v }\n", cc.b, cc.c)
	// fmt.Println(cc.Describe())
	// fmt.Println(cc.Base.Describe())

	// co := Container{Base: Base{b: 10}, c: "foo"}
	// fmt.Printf("co -> {b: %v, c: %v}\n", co.b, co.c)
	// fmt.Println(co.Describe())
	// fmt.Println(co.Base.Describe())

	b := Base{b: 10, tag: "b's tag"}
	co := Container{Base: b, c: "foo", tag: "co's tag"}

	fmt.Println(b.DescribeTag())
	fmt.Println(co.DescribeTag())

	fmt.Println("---------------------------------")

	coc := Containerfoo{Fooer: TheRealFoo{}}
	sink(coc)
}
