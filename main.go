package main

func main() {
	HEIGHT, WIDTH, horiz_coord := readWidthAndHeight()

	horiz_coord_count(WIDTH, &horiz_coord)

	// Call printMap function to print the map
	printMap(HEIGHT, WIDTH, readGridValues(WIDTH, HEIGHT), &horiz_coord)
}
