package main

import (
	//formating
	"fmt"
)

//	TO-DO:	Eingabe prüfen
//	TO-DO:	Fehler Abfangen

func main() {
	char := "."
	xValue, yValue := getUserInputFor_X_Y_Values()
	drawRectangel(xValue, yValue, char)
	fmt.Println("Der Flächeninhalt beträgt ", calculateRectangleVolume(xValue, yValue))
}

func drawRectangel(x, y int, char string) {
	for i := 1; i < x; i++ {

		for i := 1; i < y; i++ {
			fmt.Print(char)
		}
		fmt.Println(char)
	}
}

func calculateRectangleVolume(x, y int) int {
	return x * y
}

func getUserInputFor_X_Y_Values() (int, int) {
	xInput := 0
	yInput := 0
	fmt.Println("Bitte geben Sie Länge und Breite ein:")
	fmt.Scanf("%d %d", &xInput, &yInput)
	return xInput, yInput
}
