package main

func main() {
	greetingsMsg()
	height, width := readWidthAndHeight()

	if height <= 0 || width <= 0 {
		colorize(0)
		printString("Wrong input for Width or Height!\n")
		return
	}

	arr := readGridValues(width, height)

	if arr == nil {
		colorize(0)
		printString("Invalid input for matrix\n")
		return
	}

	readFormat()
	// Call printMap function to print the map
	printMap(height, width, arr)
}
