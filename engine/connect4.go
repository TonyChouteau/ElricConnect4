package engine

/*
Enum color value
*/
const (
	NONE   int = 0
	YELLOW int = 1
	RED    int = 2
)

/*
Matrix6x7 declaration
*/
type Matrix6x7 [6][7]int

/*
CopyBoard : deep copy of a board
*/
func CopyBoard(board Matrix6x7) Matrix6x7 {
	newBoard := Matrix6x7{}
	for i := range board {
		for j := range board[i] {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}

/*
CreateM : create Board matrix 3x3
*/
func CreateM() Matrix6x7 {
	return Matrix6x7{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
}

/*
ListLegal : give the list of legal move
*/
func ListLegal(board Matrix6x7) []int {
	l := []int{}
	for i := 0; i < 7; i++ {
		if board[0][i] == 0 {
			j := 0
			for j < 5 && board[j+1][i] == 0 {
				j++
			}
			l = append(l, i+j*7)
		}
	}
	return l
}

/*
Contains : if the slice contains a element
*/
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
HasWon : if someone won
*/
func HasWon(board Matrix6x7, color int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			if board[j][i] == color && board[j][i+1] == color && board[j][i+2] == color && board[j][i+3] == color {
				return true
			}
		}
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			if board[j][i] == color && board[j+1][i] == color && board[j+2][i] == color && board[j+3][i] == color {
				return true
			}
		}
	}
	for j := 0; j < 3; j++ {
		for i := 0; i < 4; i++ {
			if board[j][i] == color && board[j+1][i+1] == color && board[j+2][i+2] == color && board[j+3][i+3] == color {
				return true
			}
			if board[j][6-i] == color && board[j+1][5-i] == color && board[j+2][4-i] == color && board[j+3][3-i] == color {
				return true
			}
		}
	}
	return false
}

/*
IsLegal : is the move legal
*/
func IsLegal(board Matrix6x7, m int) bool {
	l := ListLegal(board)
	for _, e := range l {
		if e == m {
			return true
		}
	}
	return false
}

/*
NextColor : change color
*/
func NextColor(color int) int {
	if color == YELLOW {
		return RED
	} else if color == RED {
		return YELLOW
	} else {
		return NONE
	}

}

/*
Play : Play a move
*/
func Play(board Matrix6x7, move, color int) (Matrix6x7, int) {
	if color == 0 || !Contains(ListLegal(board), move) {
		return board, 99 // Error
	}

	board[move/7][move%7] = color
	if HasWon(board, color) {
		return board, color // Someone wins
	} else if len(ListLegal(board)) == 0 {
		return board, 3 // Draw
	} else {
		return board, 0 // Continue playing
	}
}
