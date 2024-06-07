package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// Adi

func main() {
	var HEIGHT, WIDTH int
	fmt.Scanf("%d %d", &HEIGHT, &WIDTH)

	full_map := make([][]int, HEIGHT)

	for i := 0; i < HEIGHT; i++ {
		full_map[i] = make([]int, WIDTH)
		for j := 0; j < WIDTH; j++ {
			fmt.Scanf("%d", &full_map[i][j])
		}
	}

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			fmt.Print(full_map[i][j])
			if j < WIDTH-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Azat
func printCell(value rune) {
	for i := 0; i < 7; i++ {
		ap.PutRune(value)
	}
}

// Additional tasks Yerkanat
