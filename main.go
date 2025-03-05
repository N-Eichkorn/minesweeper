// +---------------------------------------------------+
// | Author: Niklas Eichkorn
// | Date: 28.02.25
// | Version: 1.0
// |---------------------------------------------------+
// | Notes: merke https://gist.github.com/jordansissel/1e08b1c65157bde0f30a87c4fb569237
// +---------------------------------------------------+

package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/bit101/go-ansi"
)

func main() {

	ansi.ClearScreen()

	// get some random Numbers
	var r [4][2]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			r[i][j] = rand.IntN(10)
		}
	}

	/*
		0 = empty
		1 = mine
	*/
	a := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// fill the field with random mines
	for i := 0; i < 4; i++ {
		x := r[i][0]
		y := r[i][1]
		a[x][y] = 1
	}

	/*
		calculate Fields
		X = is a mine
	*/
	var b [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if a[i][j] == 1 {
				b[i][j] = 'X' // X stands for mine
			} else if i == 0 && j == 9 {
				b[i][j] = a[i+1][j] + a[i][j-1] + a[i+1][j-1]
			} else if i == 9 && j == 0 {
				b[i][j] = a[i-1][j] + a[i][j+1] + a[i-1][j+1]
			} else if i == 0 && j == 0 {
				b[i][j] = a[i+1][j] + a[i][j+1] + a[i+1][j+1]
			} else if i == 9 && j == 9 {
				b[i][j] = a[i-1][j-1] + a[i-1][j] + a[i][j-1]
			} else if i == 0 {
				b[i][j] = a[i+1][j] + a[i][j+1] + a[i+1][j+1] + a[i][j-1] + a[i+1][j-1]
			} else if j == 0 {
				b[i][j] = a[i+1][j] + a[i][j+1] + a[i+1][j+1] + a[i-1][j] + a[i-1][j+1]
			} else if i == 9 {
				b[i][j] = a[i][j+1] + a[i-1][j-1] + a[i-1][j] + a[i][j-1] + a[i-1][j+1]
			} else if j == 9 {
				b[i][j] = a[i+1][j] + a[i-1][j-1] + a[i-1][j] + a[i][j-1] + a[i+1][j-1]
			} else {
				b[i][j] = a[i+1][j] + a[i][j+1] + a[i+1][j+1] + a[i-1][j-1] + a[i-1][j] + a[i][j-1] + a[i+1][j-1] + a[i-1][j+1]
			}
		}
	}

	// generate player screen
	c := [10][10]int{
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
		{'█', '█', '█', '█', '█', '█', '█', '█', '█', '█'},
	}
	/*
		// Methods to Print fields

		// print the A Field clear
		ansi.Print(ansi.BoldGreen, "  0 1 2 3 4 5 6 7 8 9\n +------------------+\n")
		for i := 0; i < 10; i++ {
			ansi.Print(ansi.BoldGreen, i, "|")
			for j := 0; j < 10; j++ {
				if a[i][j] == 1 {
					ansi.Print(ansi.BoldRed, a[i][j], " ")
				} else {
					ansi.Print(ansi.BoldBlue, a[i][j], " ")
				}

			}
			fmt.Println()
		}

		// print the B Field clear
		ansi.Print(ansi.BoldGreen, "  0 1 2 3 4 5 6 7 8 9\n +------------------+\n")
		for i := 0; i < 10; i++ {
			ansi.Print(ansi.BoldGreen, i, "|")
			for j := 0; j < 10; j++ {
				if b[i][j] == 'X' {
					ansi.Print(ansi.BoldRed, b[i][j], " ")
				} else {
					ansi.Print(ansi.BoldCyan, b[i][j], " ")
				}

			}
			fmt.Println()
		}
	*/
	// print the player field
	ansi.Print(ansi.BoldGreen, "  0 1 2 3 4 5 6 7 8 9\n +------------------+\n")
	for i := 0; i < 10; i++ {
		ansi.Print(ansi.BoldGreen, i, "|")
		for j := 0; j < 10; j++ {

			i := '█'
			fmt.Printf("%c%c", i, i)

		}
		fmt.Println()
	}

	// Play
	for true {

		lost := false

		var r, t int
		fmt.Println("choose fields x y")
		fmt.Scan(&t, &r)

		// clear Screen
		ansi.ClearScreen()

		fmt.Printf("Your numbers are %d and %d\n", r, t)
		if b[r][t] == 'X' { //Mine
			ansi.Println(ansi.BoldRed, "You lost the Game!")
			c[r][t] = b[r][t]
			lost = true
		} else {
			c[r][t] = b[r][t]
		}

		/*
				Debug

			ansi.Print(ansi.BoldGreen, "  0 1 2 3 4 5 6 7 8 9\n +------------------+\n")
			for i := 0; i < 10; i++ {
				ansi.Print(ansi.BoldGreen, i, "|")
				for j := 0; j < 10; j++ {
					if c[i][j] == 'X' {
						ansi.Print(ansi.BoldRed, b[i][j], " ")
					} else {
						ansi.Print(ansi.BoldCyan, b[i][j], " ")
					}

				}
				fmt.Println()
			}
		*/

		// print the player field
		ansi.Print(ansi.BoldGreen, "  0 1 2 3 4 5 6 7 8 9\n +------------------+\n")
		for i := 0; i < 10; i++ {
			ansi.Print(ansi.BoldGreen, i, "|")
			for j := 0; j < 10; j++ {

				if c[i][j] == 0 {
					fmt.Printf("  ")
				} else if c[i][j] == 'X' {
					fmt.Printf("X ")
				} else if c[i][j] == '█' {
					fmt.Printf("██")
				} else {
					fmt.Printf("%d ", c[i][j])
				}

			}
			fmt.Println()
		}
		if lost {
			break
		}
	}
}
