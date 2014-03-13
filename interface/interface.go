package main

import (
	"fmt"
	"sort"
)

type Sequence []int

//Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	// sort.Sort(s)
	/*
		str := "["
		for i, elem := range s {
			if i > 0 {
				str += " "
			}
			str += fmt.Sprint(elem)
		}
	*/
	return fmt.Sprint([]int(s))
}

type Interfacer interface {
	String() string
}

type Animal interface {
	Speak() string
	String() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) String() string {
	return d.Speak()
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func (c *Cat) String() string {
	return c.Speak()
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

func (l Llama) String() string {
	return l.String()
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func (j JavaProgrammer) String() string {
	return j.String()
}

func main() {
	// seq := Sequence{5, 1, 2, 7, 3, 4, 6, 9, 8}
	// fmt.Println(seq)

	/*
		stringer := Interfacer{}
		var facer Interfacer
		str, ok := facer.(string)
		if ok {
			fmt.Printf("string value is: %q\n", str)
		} else {
			fmt.Printf("value is not a string\n")
		}
	*/

	dog := &Dog{}
	cat := &Cat{}
	llama := &Llama{}
	jp := &JavaProgrammer{}
	animals := []Animal{dog, cat, llama, jp}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
