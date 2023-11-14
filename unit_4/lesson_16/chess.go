package main

import "fmt"

var chessBoardInitialPositions [8][8]rune

func getInitialPositions() {
	chessBoardInitialPositions[0][0] = 'R'
	chessBoardInitialPositions[0][1] = 'K'
	chessBoardInitialPositions[0][2] = 'B'
	chessBoardInitialPositions[0][3] = 'Q'
	chessBoardInitialPositions[0][4] = 'K'
	chessBoardInitialPositions[0][5] = 'B'
	chessBoardInitialPositions[0][6] = 'K'
	chessBoardInitialPositions[0][7] = 'R'

	chessBoardInitialPositions[7][0] = 'r'
	chessBoardInitialPositions[7][1] = 'k'
	chessBoardInitialPositions[7][2] = 'b'
	chessBoardInitialPositions[7][3] = 'q'
	chessBoardInitialPositions[7][4] = 'k'
	chessBoardInitialPositions[7][5] = 'b'
	chessBoardInitialPositions[7][6] = 'k'
	chessBoardInitialPositions[7][7] = 'r'

	for i := range chessBoardInitialPositions[1] {
		chessBoardInitialPositions[1][i] = 'P'
		// fmt.Println(v)
	}
	for i := range chessBoardInitialPositions[7] {
		chessBoardInitialPositions[6][i] = 'p'
	}

	fmt.Println(chessBoardInitialPositions)

}

func displayBoard() {
	fmt.Println("---------------------------------")
	for i := len(chessBoardInitialPositions) - 1; i >= 0; i-- {
		for _, v2 := range chessBoardInitialPositions[i] {
			if v2 == 0 {
				fmt.Printf("|%3c ", v2)
			} else {
				fmt.Printf("|%2c ", v2)

			}

		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("---------------------------------")
}
