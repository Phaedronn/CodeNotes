package main

import "fmt"

func balls(title string, writer string, artist string, publisher string, year int, pageNumber int, grade float64) {
	fmt.Println(
		"\n\n-8===D------------(l)",
		"\n"+title,
		"\n\nWritten by:",
		writer,
		"\n\nDrawn by:",
		artist,
		"\n\nPublished by:",
		publisher,
		"\n\nYear printed:",
		year,
		"\n\nNumber of Pages:",
		pageNumber,
		"\n\nCondition:",
		grade,
	)
}

func main() {

	title := "Mr. GoToSleep"
	writer := "Tracey Hatchet"
	artist := "Jewel Tampson"
	publisher := "DizzyBooks Publishing Inc."
	year := 1997
	pageNumber := 14
	grade := 6.5

	balls(title, writer, artist, publisher, year, pageNumber, grade)

	//These are place holder names.

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	balls(title, writer, artist, publisher, year, pageNumber, grade)
}
