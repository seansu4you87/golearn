package main

import (
	"errors"
	"fmt"
	"time"
)

/* Pile */
type Pile struct {
	name    string
	books   []Book
	gophers chan *Gopher
	errors  chan error
}

func NewPile(name string, books []Book) (p *Pile) {
	p = &Pile{
		name:    name,
		books:   books,
		gophers: make(chan *Gopher),
		errors:  make(chan error),
	}

	go p.run()
	return
}

func (p *Pile) run() {
	for {
		select {
		case g := <-p.gophers:
			var b Book
			p.ReportStatus(fmt.Sprintf("%s trying to take a book", g.name))
			books := p.books
			origBooksLeft := len(books)
			if origBooksLeft == 0 {
				p.ReportStatus(fmt.Sprintf("out of books, %s failed!", g.name))
				err := errors.New("Pile out of books")
				p.errors <- err
			} else {
				time.Sleep(1 * time.Second)
				last := len(books) - 1
				b, p.books = books[last], books[:last]

				status := fmt.Sprintf("I have %d books, %s taking 1, %d books left", origBooksLeft, g.name, len(p.books))
				p.ReportStatus(status)
				p.errors <- nil
				g.Book = b
			}
		}
	}
}

func (p *Pile) ReportStatus(status string) {
	fmt.Printf("Pile %s:\t %s\n", p.name, status)
}

func (p *Pile) RequestBook(g *Gopher) (b Book, err error) {
	p.gophers <- g
	err = <-p.errors
	return b, err
}
