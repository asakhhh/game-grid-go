package main

func main() {
	height, width, horiz_coord := readWidthAndHeight()

	if horiz_coord == nil || height <= 0 || width <= 0 {
		printString("Wrong input for Width or Height!\n")
	} else {
		horiz_coord_count(width, &horiz_coord)
		arr := readGridValues(width, height)

		if arr == nil {
			printString("Invalid input for matrix\n")
		} else {
			// Call printMap function to print the map
			printMap(height, width, arr, &horiz_coord)
		}
	}
}
