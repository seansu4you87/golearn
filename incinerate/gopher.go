package main

import "fmt"

/* Gopher */
type Gopher struct {
	name        string
	Book        Book
	Pile        *Pile
	Incinerator *Incinerator
	BookCount   int
	BurnCount   int
}

func (g *Gopher) ReportStatus(status string) {
	fmt.Printf("Gopher %s:\t %s\n", g.name, status)
}

//Implement Named
func (g *Gopher) Name() string {
	return g.name
}

func (g *Gopher) Work() {
	b, err := g.Pile.RequestBook(g)
	if err != nil {
		status := fmt.Sprintf("Done with work! Book count: %d", g.BookCount)
		g.ReportStatus(status)
	} else {
		g.Book = b
		g.BookCount++
		g.ReportStatus("Taking Book")

		err = g.Incinerator.RequestBurn(g)
		if err != nil {
			g.ReportStatus("fucked up burning =[")
		} else {
			g.ReportStatus("Book Burned!")
			g.BurnCount++
			g.Work()
		}
	}
}
