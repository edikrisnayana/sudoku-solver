package main

func proccess(board *boardGame) {
	size := board.subBoardSize * board.subBoardSize
	for !board.isComplete() {
		markTile(board)
		arr := make([][]bool, size)
		for i := range arr {
			arr[i] = make([]bool, size)
		}
		proccessFill(arr, 0, 0, board, 1)
		board.printConsole()
	}
	board.printOutput()
}

func proccessFill(alreadyProccess [][]bool, i int, j int, board *boardGame, tileProcessedCount int) {
	size := board.subBoardSize * board.subBoardSize
	alreadyProccess[i][j] = true
	tile := board.getTile(i, j)
	if tile.isFixed() {
		if tileProcessedCount == size*size {
			fillAll(board)
		}
		if i > 0 && !alreadyProccess[i-1][j] {
			proccessFill(alreadyProccess, i-1, j, board, tileProcessedCount+1)
		} else if i < size-1 && !alreadyProccess[i+1][j] {
			proccessFill(alreadyProccess, i+1, j, board, tileProcessedCount+1)
		} else if j > 0 && !alreadyProccess[i][j-1] {
			proccessFill(alreadyProccess, i, j-1, board, tileProcessedCount+1)
		} else if j < size-1 && !alreadyProccess[i][j+1] {
			proccessFill(alreadyProccess, i, j+1, board, tileProcessedCount+1)
		}
	} else {
		for number := 1; number <= size; number++ {
			if tile.valuePossibility[number-1] {
				exist := isNumberExist(int8(number), i, j, false, board)
				if !exist {
					board.addTempNumber(int8(number), i, j)
					if tileProcessedCount == size*size {
						fillAll(board)
					}
					if i > 0 && !alreadyProccess[i-1][j] {
						proccessFill(alreadyProccess, i-1, j, board, tileProcessedCount+1)
					} else if i < size-1 && !alreadyProccess[i+1][j] {
						proccessFill(alreadyProccess, i+1, j, board, tileProcessedCount+1)
					} else if j > 0 && !alreadyProccess[i][j-1] {
						proccessFill(alreadyProccess, i, j-1, board, tileProcessedCount+1)
					} else if j < size-1 && !alreadyProccess[i][j+1] {
						proccessFill(alreadyProccess, i, j+1, board, tileProcessedCount+1)
					}
					board.removeTempNumber(i, j)
				}
			}
		}
	}
	alreadyProccess[i][j] = false
}

func fillAll(board *boardGame) {
	size := board.subBoardSize * board.subBoardSize
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tile := board.getTile(i, j)
			if !tile.isFixed() {
				tile.fixedValue = tile.tempValue
				tile.tempValue = 0
			}
		}
	}
}

func markTile(board *boardGame) {
	size := board.subBoardSize * board.subBoardSize
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tile := board.getTile(i, j)
			if !tile.isFixed() {
				for number := 1; number <= size; number++ {
					exist := isNumberExist(int8(number), i, j, true, board)
					tile.valuePossibility[number-1] = !exist
				}
			}
		}
	}
}

func isNumberExist(number int8, i int, j int, forMark bool, board *boardGame) bool {
	sub := board.getSubBoard(i, j)
	tile := sub.getTile(i, j)
	exist := isNumberExistInSubBoard(number, tile.id, forMark, *sub)
	exist = exist || isNumberExistInRow(number, i, j, forMark, board)
	exist = exist || isNumberExistInColumn(number, i, j, forMark, board)
	return exist
}

func isNumberExistInSubBoard(number int8, currId int8, forMark bool, sub subBoard) bool {
	for _, tilei := range sub.tiles {
		for _, tile := range tilei {
			if currId == tile.id {
				continue
			}

			if isNumberExistInTile(number, *tile, forMark) {
				return true
			}
		}
	}

	return false
}

func isNumberExistInRow(number int8, i int, j int, forMark bool, board *boardGame) bool {
	indexj := j - 1
	for indexj >= 0 {
		tile := board.getTile(i, indexj)
		if number == tile.fixedValue || number == tile.tempValue {
			return true
		}
		indexj--
	}
	size := board.subBoardSize * board.subBoardSize
	indexj = j + 1
	for indexj < size {
		tile := board.getTile(i, indexj)
		if isNumberExistInTile(number, *tile, forMark) {
			return true
		}
		indexj++
	}

	return false
}

func isNumberExistInColumn(number int8, i int, j int, forMark bool, board *boardGame) bool {
	indexi := i - 1
	for indexi >= 0 {
		tile := board.getTile(indexi, j)
		if number == tile.fixedValue || number == tile.tempValue {
			return true
		}
		indexi--
	}
	size := board.subBoardSize * board.subBoardSize
	indexi = i + 1
	for indexi < size {
		tile := board.getTile(indexi, j)
		if isNumberExistInTile(number, *tile, forMark) {
			return true
		}
		indexi++
	}

	return false
}

func isNumberExistInTile(number int8, tile tile, forMark bool) bool {
	exist := number == tile.fixedValue
	if !forMark {
		exist = exist || number == tile.tempValue
	}
	return exist
}
