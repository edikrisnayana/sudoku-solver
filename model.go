package main

import "fmt"

type tile struct {
	id               int8
	fixedValue       int8
	tempValue        int8
	valuePossibility []bool
}

type subBoard struct {
	complete bool
	tiles    [][]*tile
}

func (t *tile) isFixed() bool {
	if t.fixedValue == 0 {
		number := -1
		possibilityCount := 0
		for i, possible := range t.valuePossibility {
			if possible {
				number = i + 1
				possibilityCount++
			}
		}

		if possibilityCount == 1 {
			t.fixedValue = int8(number)
		}
	}
	return t.fixedValue > 0
}

func (sub *subBoard) getTile(i int, j int) *tile {
	size := len(sub.tiles) * len(sub.tiles)
	if i < 0 || i > size-1 || j < 0 || j > size-1 {
		fmt.Printf("tile[%d][%d] not exist\n", i, j)
		return nil
	}

	tilei := i % len(sub.tiles)
	tilej := j % len(sub.tiles)
	tile := sub.tiles[tilei][tilej]
	return tile
}

func (sub *subBoard) isComplete() bool {
	if !sub.complete {
		complete := true
		for _, tilei := range sub.tiles {
			for _, tile := range tilei {
				if tile.fixedValue == 0 {
					complete = false
					break
				}
			}

			if !complete {
				break
			}
		}
		sub.complete = complete
	}

	return sub.complete
}
