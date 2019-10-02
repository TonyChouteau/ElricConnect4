package ai

import (
	"fmt"
	"math/rand"

	"github.com/TonyChouteau/elricconnect4/engine"
)

func testMove(board engine.Matrix6x7, move int, color int) float64 {

	board, status := engine.Play(board, move, color)
	if len(engine.ListLegal(board)) == 0 {
		return 0.5
	}

	backupBoard := engine.CopyBoard(board)

	var nbOfWins float64
	n := 10000

	for i := 0; i < n; i++ {
		testBoard := engine.CopyBoard(backupBoard)
		currentColor := engine.NextColor(color)

		finished := false
		for !finished {
			legalMoves := engine.ListLegal(testBoard)
			testBoard, status = engine.Play(testBoard, legalMoves[rand.Intn(len(legalMoves))], currentColor)
			if status == color {
				nbOfWins++
				finished = true
			} else if status == engine.NextColor(color) {
				finished = true
			} else if status == 3 {
				nbOfWins += 0.5
				finished = true
			}
			currentColor = engine.NextColor(currentColor)
		}
	}

	//fmt.Println(nbOfWins, n)
	probWins := nbOfWins / float64(n)
	return probWins
}

/*
GetBestMove : return best move
*/
func GetBestMove(inlineBoard string) int {

	// Decode URI board into a 3x3 Matrix
	if len(inlineBoard) == 42 {
		board := engine.Matrix6x7{}
		for i := range inlineBoard {
			board[i/7][i%7] = int(inlineBoard[i]) - 48
		}
		fmt.Println(board)

		l := engine.ListLegal(board)
		if len(l) == 0 {
			return -1
		}
		results := []float64{}
		for i := range l {
			results = append(results, testMove(board, l[i], engine.RED))
		}
		maxIndex := 0
		for i := range results {
			if results[i] > results[maxIndex] {
				maxIndex = i
			}
		}
		fmt.Println(results)
		return l[maxIndex]
	}

	fmt.Println("ERROR : board URI must have 9 values.")
	return -1
}
