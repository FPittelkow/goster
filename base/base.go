package main

import (
	//formating
	"fmt"
)

//	TO-DO:	Eingabe prüfen
//	TO-DO:	Fehler Abfangen

type rectangel struct {
	x, y int
}

func main() {

	rect := new(rectangel)

	fmt.Println("Bitte geben Sie Breite und Länge ein:")
	fmt.Scanf("%d %d", &rect.x, &rect.y)

	drawRectangel(rect.x, rect.y)
	fmt.Println("Der Flächeninhalt beträgt ", calculateRectangleVolume(rect.x, rect.y))
}

func drawRectangel(x, y int) {
	for i := 1; i < x; i++ {

		for i := 1; i < y; i++ {
			fmt.Print("-")
		}
		fmt.Println("-")
	}
}

func calculateRectangleVolume(x, y int) int {
	return x * y
}
