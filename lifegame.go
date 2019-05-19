package main

import (
	"flag"
	"log"

	termbox "github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

func tbprint(table [][]bool) {
	for y := range table {
		for x := range table[0] {
			if table[y][x] {
				termbox.SetCell(x*2, y, '　', termbox.ColorWhite, termbox.ColorWhite)
			} else {
				termbox.SetCell(x*2, y, '　', coldef, coldef)
			}
		}
	}
	termbox.Flush()
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Usage: ./lifegame {path/to/csvfile}")
	}
	filePath := args[0]

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(coldef, coldef)
	table := FromFile(filePath)
	tbprint(table)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		}
		table = UpdateTable(table)
		tbprint(table)
	}
}
