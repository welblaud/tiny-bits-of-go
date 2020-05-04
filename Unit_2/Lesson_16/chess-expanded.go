// chess.go
package main

import "fmt"

func printBoard(board [8][8]rune) {
	fmt.Println("----------------")
	// don't forget, _, column instead of board[column]
	for _, column := range board {
		fmt.Print("|")
		for _, cell := range column {
			if cell == 0 {
				fmt.Print(" |")
			} else {
				fmt.Printf("%c|", cell)
			}
		}
		fmt.Println("\n-----------------")
	}
}

func main() {
	var board [8][8]rune
	board[0][1] = 'k'
	board[0][2] = 'q'
	board[0][3] = 'n'
	board[0][4] = 'b'
	board[0][5] = 'r'
	board[0][6] = 'p'

	board[7][1] = 'K'
	board[7][2] = 'Q'
	board[7][3] = 'N'
	board[7][4] = 'B'
	board[7][5] = 'R'
	board[7][6] = 'P'
	printBoard(board)
	// fmt.Printf("%c", board[0][2])
}
