package main

func main() {
	HEIGHT, WIDTH := readWidthAndHeight()

	horiz_coord := make([][7]rune, WIDTH)
	horiz_coord_count(WIDTH, &horiz_coord)

	full_map := readGridValues(WIDTH, HEIGHT)

	// Call printMap function to print the map
	printMap(HEIGHT, WIDTH, full_map, &horiz_coord)
}
