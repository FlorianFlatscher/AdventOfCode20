package solution

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/constants"
	"github.com/FlorianFlatscher/AdventOfCode/src/input"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day20 struct{}

type tile struct {
	ID   int
	Data [][]bool
}

var (
	regexTileId = regexp.MustCompile("\\w+ (\\d+):")
)

func newTile(input string) *tile {
	split := strings.SplitN(input, constants.LineSeparator, 2)
	id, err := strconv.Atoi(regexTileId.FindStringSubmatch(split[0])[1])
	if err != nil {
		panic(err)
	}

	data := make([][]bool, 0, 10)
	for _, line := range strings.Split(split[1], constants.LineSeparator) {
		lineData := make([]bool, 0, 10)
		for _, char := range line {
			switch rune(char) {
			case '#':
				lineData = append(lineData, true)
			case '.':
				lineData = append(lineData, false)
			default:
				panic("Invalid input")
			}
		}
		data = append(data, lineData)
	}

	newTile := tile{
		ID:   id,
		Data: data,
	}

	return &newTile
}

func (t *tile) alignsHorizontally(other *tile) bool {
	if len(t.Data) != len(other.Data) {
		panic("invalid")
	}
	for r := range t.Data {
		if t.Data[r][len(t.Data[r])-1] !=
			other.Data[r][0] {
			return false
		}
	}
	return true
}

func (t *tile) alignsVertically(other *tile) bool {
	if len(t.Data) != len(other.Data) {
		panic("invalid")
	}
	for columnIndex := range t.Data[0] {
		if t.Data[len(t.Data)-1][columnIndex] !=
			other.Data[0][columnIndex] {
			return false
		}
	}
	return true
}

func (t *tile) String() string {
	sb := strings.Builder{}
	sb.WriteString("Tile ")
	sb.WriteString(strconv.Itoa(t.ID))
	sb.WriteString(":\n")
	for _, row := range t.Data {
		lineSb := strings.Builder{}
		for _, car := range row {
			if car {
				lineSb.WriteRune('#')
			} else {
				lineSb.WriteRune('.')
			}
		}
		lineSb.WriteRune('\n')
		sb.WriteString(lineSb.String())
	}
	return sb.String()
}

func (t *tile) Rotate() *tile {
	newData := make([][]bool, len(t.Data))
	for row := range t.Data {
		newData[row] = make([]bool, len(t.Data[row]))
	}
	for row := range t.Data {
		for col := range t.Data[row] {
			newR, newC := col, row
			newC = len(t.Data) - 1 - newC
			newData[newR][newC] = t.Data[row][col]
		}
	}
	newTile := tile{
		ID:   t.ID,
		Data: newData,
	}
	return &newTile
}

func (t *tile) FlipHorizontally() *tile {
	newData := make([][]bool, len(t.Data))
	for row := range t.Data {
		newData[row] = make([]bool, len(t.Data[row]))
	}
	for row := range t.Data {
		for col := range t.Data[row] {
			newC, newR := col, row
			newC = len(t.Data[row]) - 1 - newC
			newData[newR][newC] = t.Data[row][col]
		}
	}
	newTile := tile{
		ID:   t.ID,
		Data: newData,
	}
	return &newTile
}

func (t *tile) FlipVertically() *tile {
	newData := make([][]bool, len(t.Data))
	for row := range t.Data {
		newData[row] = make([]bool, len(t.Data[row]))
	}
	for row := range t.Data {
		for col := range t.Data[row] {
			newC, newR := col, row
			newR = len(t.Data) - 1 - newR
			newData[newR][newC] = t.Data[row][col]
		}
	}
	newTile := tile{
		ID:   t.ID,
		Data: newData,
	}
	return &newTile
}

func (d Day20) Calc() string {
	tileInputs := strings.Split(input.ReadInputFile(20), strings.Repeat(constants.LineSeparator, 2))
	var tiles []tile
	for _, tileInput := range tileInputs {
		tiles = append(tiles, *newTile(tileInput))
	}
	//fmt.Println("********************** Tile 1:")
	//fmt.Println(tiles[0].String())
	//fmt.Println("********************** Rotate:")
	//fmt.Println(tiles[0].Rotate().String())
	order := d.alignTiles(nil, tiles, int(math.Round(math.Sqrt(float64(len(tiles))))))

	return fmt.Sprintf("1: %v\n2: %v\n", d.multiplyCorners(order), d.findSeaMonster(order))
	//return fmt.Sprintf("1: %v\n2: %v\n", nil, nil)
}

func (d Day20) multiplyCorners(order [][]tile) int {
	return order[0][0].ID * order[0][len(order[0])-1].ID * order[len(order)-1][0].ID * order[len(order)-1][len(order[0])-1].ID
}

func (d Day20) findSeaMonster(order [][]tile) int {
	combinedImage := make([][]bool, len(order)*(len(order[0][0].Data)-2))
	for _ = range order {
		for r := 0; r < len(order)*(len(order[0][0].Data)-2); r++ {
			combinedImage[r] = make([]bool, len(order)*(len(order[0][0].Data)-2))
		}
	}

	for tileRow := range order {
		for tileColumn := range order {
			t := order[tileRow][tileColumn]
			for r := range t.Data {
				if r > 0 && r < len(t.Data)-1 {
					for c := range t.Data[r] {
						if c > 0 && c < len(t.Data[r])-1 {
							combinedImage[tileRow*(len(t.Data)-2)+r-1][tileColumn*(len(t.Data[r])-2)+c-1] = t.Data[r][c]
						}
					}
				}
			}
		}
	}
	image := tile{
		ID:   0,
		Data: combinedImage,
	}
	seaMonster := strings.Split("                  # \n#    ##    ##    ###\n #  #  #  #  #  #   ", "\n")
	//Try out with every rotation
	rotatedTile := image
	for r := 0; r < 4; r++ {
		//Try to flip it
		for f := 0; f < 2; f++ {
			flippedTile := rotatedTile
			if f == 1 && r%2 == 0 {
				flippedTile = *flippedTile.FlipVertically()
			} else if f == 1 && f%2 == 1 {
				flippedTile = *flippedTile.FlipHorizontally()
			}
			monsterCount := 0
			water := 0
			image := flippedTile.Data
			for row := range image {
				for col := range image[row] {
					if image[row][col] {
						water++
					}
					invalid := false
					//Check if is monster
					for monsterY, monsterLine := range seaMonster {
						for monsterX, car := range monsterLine {
							if car == '#' {
								if row+monsterY >= len(image) ||
									col+monsterX >= len(image) {
									invalid = true
									break
								}
								if !image[row+monsterY][col+monsterX] {
									invalid = true
									break
								}
							}
						}
						if invalid {
							break
						}
					}
					if !invalid {
						water -= 15
						monsterCount++
					}
				}
			}

			if monsterCount > 0 {
				return water
			}
		}
		rotatedTile = *rotatedTile.Rotate()
	}
	return -1
}

func (d Day20) alignTiles(orderSoFar [][]tile, tilesLeft []tile, sideLength int) [][]tile {
	if len(tilesLeft) == 0 {
		return orderSoFar
	}
	//for r := range orderSoFar {
	//	for c := range orderSoFar[r] {
	//		fmt.Print(orderSoFar[r][c].ID, ", ")
	//	}
	//}
	//fmt.Println("")
	for tileIndex, nextTile := range tilesLeft {
		copyOrder := make([][]tile, len(orderSoFar))
		for i := range orderSoFar {
			copyOrder[i] = make([]tile, len(orderSoFar[i]))
			copy(copyOrder[i], orderSoFar[i])
		}
		newTilesLeft := make([]tile, len(tilesLeft))
		copy(newTilesLeft, tilesLeft)
		newTilesLeft = append(newTilesLeft[:tileIndex], newTilesLeft[tileIndex+1:]...)

		if len(copyOrder) > 0 && len(copyOrder[len(copyOrder)-1]) < sideLength {
			//Try out with every rotation
			rotatedTile := nextTile
			for r := 0; r < 4; r++ {
				//Try to flip it
				for f := 0; f < 2; f++ {
					flippedTile := rotatedTile
					if f == 1 && r%2 == 0 {
						flippedTile = *flippedTile.FlipVertically()
					} else if f == 1 && f%2 == 1 {
						flippedTile = *flippedTile.FlipHorizontally()
					}
					//Try out appending it to this line
					copyOrder[len(copyOrder)-1] = append(copyOrder[len(copyOrder)-1], flippedTile)
					if d.validateOrder(copyOrder) {
						try := d.alignTiles(copyOrder, newTilesLeft, sideLength)
						if try != nil {
							return try
						}
					}
					//Otherwise remove it again
					copyOrder[len(copyOrder)-1] = copyOrder[len(copyOrder)-1][:len(copyOrder[len(copyOrder)-1])-1]
				}
				//rotate the tile
				rotatedTile = *rotatedTile.Rotate()
			}
		} else {
			//Create a new line
			rotatedTile := nextTile
			for r := 0; r < 4; r++ {
				//Try to flip it
				for f := 0; f < 3; f++ {
					flippedTile := rotatedTile
					if f > 0 && f%2 == 0 {
						flippedTile = *flippedTile.FlipHorizontally()
					} else if f > 0 && f%2 == 1 {
						flippedTile = *flippedTile.FlipVertically()
					}
					//Try out appending it to next line
					copyOrder = append(copyOrder, []tile{flippedTile})
					if d.validateOrder(copyOrder) {
						try := d.alignTiles(copyOrder, newTilesLeft, sideLength)
						if try != nil {
							return try
						}
					}
					//Otherwise remove it again
					copyOrder = copyOrder[:len(copyOrder)-1]
				}
				//rotate the tile
				rotatedTile = *rotatedTile.Rotate()
			}
		}
	}

	return nil
}

func (d Day20) printOrder(order [][]tile) {
	for r, row := range order {
		fmt.Println("------------- Row", r, " -------------")
		for _, col := range row {
			fmt.Println(col.String())
		}
	}
}

func (d Day20) validateOrder(order [][]tile) bool {
	//fmt.Println("***************************")
	//d.printOrder(order)
	//defer func() {
	//	fmt.Println("***************************")
	//	fmt.Println("")
	//}()

	//validate horizontally
	for _, row := range order {
		for col := 0; col < len(row)-1; col++ {
			if !row[col].alignsHorizontally(&row[col+1]) {
				return false
			}
		}
	}

	//Validate vertically
	for r := 0; r < len(order)-1; r++ {
		row := order[r]
		for c := 0; c < len(row); c++ {
			if len(order[r+1]) > c && !order[r][c].alignsVertically(&order[r+1][c]) {
				return false
			}
		}
	}

	return true
}
