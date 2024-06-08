package main

import "fmt"

// ANSI color codes for console text styling
var (
	RESET  = "\033[0m"  // Reset text styling
	RED    = "\033[31m" // Red text color
	WHITE  = "\033[97m" // White text color
	BLUE   = "\033[34m" // Blue text color
	YELLOW = "\033[33m" // Yellow text color
)

// Constants representing different elements on the map
var (
	wall   = '0' // Wall character
	player = '2' // Player character
	award  = '3' // Award character
)

// main function reads input and calls the printMap function
func main() {
	var HEIGHT, WIDTH int
	fmt.Scanf("%d %d", &HEIGHT, &WIDTH)

	full_map := make([][]int, HEIGHT)

	// Read input grid values
	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%d", &full_map[i][j])
		}
	}

	// Call printMap function to print the map
	printMap(HEIGHT, WIDTH, full_map)
}

// printCell function prints a single cell based on its value
func printCell(value int, x int, y int) {
	if value == 0 {
		fmt.Print(RED + "X" + RESET) // Wall
	} else if value == 1 {
		if x%3 == 0 {
			fmt.Print("_") // Horizontal separator
		} else {
			fmt.Print(" ") // Empty space
		}
	} else if value == 2 {
		fmt.Print(BLUE + ">" + RESET) // Player
	} else if value == 3 {
		fmt.Print(YELLOW + "@" + RESET) // Award
	}
}

// printMap function prints the map along with coordinates
func printMap(height int, width int, value [][]int) {
	// Loop to iterate over every row.
	// Multiplied by 3 because every row has 3 rows inside.
	// And additional 1 to print first line
	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 && h != 0 {
			fmt.Printf("%2d", (h+1)/3) // Print vertical coordinate every 3rd row
		} else {
			fmt.Print("  ") // Print space for alignment
		}

		// Loop to iterate over every column
		if h == 0 {
			for w := 0; w < (7*width)+width+1; w++ {
				if (w+3)%7 == 0 && w != 0 {
					fmt.Print(string(rune(64 + (w+3)/7))) // Print horizontal coordinates
				} else {
					fmt.Print(" ") // Print space for alignment
				}
			}
			fmt.Print("\n") // Move to the next line after printing coordinates
		}

		for w := 0; w < (8*width)+1; w++ {
			if w == 0 && h != 0 || w%8 == 0 && h != 0 {
				fmt.Print("|") // Vertical border
			} else if h == 0 {
				if w == 0 {
					fmt.Print("  ") // Print space for alignment before the first row
				}
				fmt.Print("_") // Horizontal border
			} else if h%3 != 0 {
				printCell(value[h/3][w/8], h, w) // Print cell content
			} else {
				printCell(value[h/3-1][w/8], h, w) // Print cell content
			}
		}
		fmt.Print("\n") // Move to the next line after printing a row
	}
}
