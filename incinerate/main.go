package main

/* Book */
type Book struct {
}

type Named interface {
	Name() string
}

func main() {
	books1 := []Book{}
	books2 := []Book{}
	for i := 0; i < 5; i++ {
		books1 = append(books1, Book{})
		books2 = append(books2, Book{})
	}

	p1 := NewPile("P-1", books1)
	p2 := NewPile("P-2", books2)
	i1 := NewIncinerator("I-1")
	i2 := NewIncinerator("I-2")

	g1 := &Gopher{
		name:        "Al",
		Pile:        p1,
		Incinerator: i1,
	}

	g2 := &Gopher{
		name:        "Bob",
		Pile:        p2,
		Incinerator: i2,
	}

	g3 := &Gopher{
		name:        "Chad",
		Pile:        p1,
		Incinerator: i1,
	}

	g4 := &Gopher{
		name:        "Dan",
		Pile:        p2,
		Incinerator: i2,
	}

	g5 := &Gopher{
		name:        "Ed",
		Pile:        p1,
		Incinerator: i1,
	}

	g6 := &Gopher{
		name:        "Frank",
		Pile:        p2,
		Incinerator: i2,
	}

	go g1.Work()
	go g2.Work()
	go g3.Work()
	go g4.Work()
	go g5.Work()
	go g6.Work()

	select {
	// default:
	// 	fmt.Println("No Comms")
	}
}
