package main

import (
	"fmt"
	"os"
	"strings"
)

type boardGame struct {
	complete     bool
	subBoardSize int
	subBoards    [][]*subBoard
}

func createBoardGame(size int) boardGame {
	board := boardGame{
		subBoardSize: size,
		subBoards:    make([][]*subBoard, size),
	}

	for i := range board.subBoards {
		board.subBoards[i] = make([]*subBoard, size)
		for j := range board.subBoards[i] {
			board.subBoards[i][j] = createSubBoard(size)
		}
	}

	return board
}

func createSubBoard(size int) *subBoard {
	subBoard := &subBoard{
		complete: false,
		tiles:    make([][]*tile, size),
	}

	var id int8 = 1
	for i := range subBoard.tiles {
		subBoard.tiles[i] = make([]*tile, size)
		for j := range subBoard.tiles[i] {
			subBoard.tiles[i][j] = createTile(id, size*size)
			id++
		}
	}

	return subBoard
}

func createTile(id int8, possibleValueSize int) *tile {
	tile := &tile{
		id:               id,
		fixedValue:       0,
		tempValue:        0,
		valuePossibility: make([]bool, possibleValueSize),
	}
	return tile
}

func (board *boardGame) isComplete() bool {
	if !board.complete {
		complete := true
		for _, subi := range board.subBoards {
			for _, sub := range subi {
				if !sub.isComplete() {
					complete = false
					break
				}
			}

			if !complete {
				break
			}
		}
		board.complete = complete
	}

	return board.complete
}

func (board *boardGame) getSubBoard(i int, j int) *subBoard {
	size := board.subBoardSize * board.subBoardSize
	if i < 0 || i > size-1 || j < 0 || j > size-1 {
		fmt.Printf("subBoard[%d][%d] not exist\n", i, j)
		return nil
	}

	subi := i / board.subBoardSize
	subj := j / board.subBoardSize
	return board.subBoards[subi][subj]
}

func (board *boardGame) getTile(i int, j int) *tile {
	subBoard := board.getSubBoard(i, j)

	if subBoard == nil {
		return nil
	}

	tile := subBoard.getTile(i, j)
	return tile
}

func (board *boardGame) addFixedNumber(number int8, i int, j int) {
	size := board.subBoardSize * board.subBoardSize
	if number < 1 || number > int8(size) {
		fmt.Printf("value %d is not expected\n", number)
		return
	}

	tile := board.getTile(i, j)

	if tile.fixedValue > 0 {
		fmt.Printf("Can't update. Value for tile[%d][%d] already fixed\n", i, j)
		return
	}

	tile.fixedValue = number
}

func (board *boardGame) addTempNumber(number int8, i int, j int) {
	size := board.subBoardSize * board.subBoardSize
	if number < 1 || number > int8(size) {
		fmt.Printf("value %d is not expected\n", number)
		return
	}

	tile := board.getTile(i, j)

	(*tile).tempValue = number
}

func (board *boardGame) removeTempNumber(i int, j int) {
	tile := board.getTile(i, j)

	(*tile).tempValue = 0
}

func (board *boardGame) toString() string {
	sb := strings.Builder{}
	size := board.subBoardSize * board.subBoardSize
	length := size*2 + board.subBoardSize
	for i := 1; i <= 2; i++ {
		for j := 0; j < length; j++ {
			sb.WriteString("-")
		}
		sb.WriteString("\n")
	}

	for i := 0; i < size; i++ {
		sb.WriteString("|")
		for j := 0; j < size; j++ {
			tile := board.getTile(i, j)
			value := (*tile).fixedValue
			if value == 0 {
				value = (*tile).tempValue
			}
			sb.WriteString(fmt.Sprintf("%d|", value))
			if j%board.subBoardSize == board.subBoardSize-1 && j < size-1 {
				sb.WriteString("|")
			}
		}

		sb.WriteString("\n")
		for j := 0; j < length; j++ {
			sb.WriteString("-")
		}
		sb.WriteString("\n")

		if i%board.subBoardSize == board.subBoardSize-1 && i < size {
			for j := 0; j < length; j++ {
				sb.WriteString("-")
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func (board *boardGame) printOutput() {
	file, createErr := os.Create("output.txt")
	if createErr != nil {
		fmt.Println(createErr)
	}

	file.WriteString(board.toString())
}

func (board *boardGame) printConsole() {
	fmt.Println(board.toString())
}
