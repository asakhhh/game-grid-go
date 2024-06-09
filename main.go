package main

func main() {
	// function to explain format of input
	greetingsMsg()

	// reads dimension of map
	height, width := readWidthAndHeight()

	if height <= 0 || width <= 0 {
		colorizeFG(1)
		printString("Wrong input for Width or Height!\n")
		return
	}
	// reads grid values of map
	arr := readGridValues(width, height)

	if arr == nil {
		colorizeFG(1)
		printString("Invalid input for matrix\n")
		return
	}

	readFormat()
	// Call printMap function to print the map
	printMap(height, width, arr)
}
