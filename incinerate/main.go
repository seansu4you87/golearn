package main

import (
	// "fmt"
	"time"
)

/* Book */
type Book struct {
}

type Named interface {
	Name() string
}

func main() {
	books := []Book{}
	for i := 0; i < 10; i++ {
		books = append(books, Book{})
	}

	p := NewPile("P-1", books)
	i := NewIncinerator("I-1")

	g1 := &Gopher{
		name:        "Al",
		Pile:        p,
		Incinerator: i,
	}

	g2 := &Gopher{
		name:        "Bob",
		Pile:        p,
		Incinerator: i,
	}

	go g1.Work()
	time.Sleep(9 * time.Second)
	go g2.Work()

	select {
	// default:
	// 	fmt.Println("No Comms")
	}
}
