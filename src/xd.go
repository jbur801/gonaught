package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gridSize = 20
var winAmount = 3

type grid = [][]int

func create2DArray(x int) [][]int {
	// Create a 2D slice with x rows and x columns
	array := make([][]int, x)
	for i := range array {
		array[i] = make([]int, x)
	}
	return array
}

func printGrid(grid grid) {
	for x := 0; x < gridSize; x++ {
		line := ""
		for y := 0; y < gridSize; y++ {
			gridVal := "[-]"
			switch kek := grid[x][y]; kek {
			case 0:
				gridVal = "[-]"

			case 1:
				gridVal = "[O]"

			case 2:
				gridVal = "[X]"
			}
			line += gridVal
		}
		fmt.Println(line)
	}
}

func print2DArray(array [][4]int) string {
	var builder strings.Builder
	for _, row := range array {
		builder.WriteString(fmt.Sprintf("%v\n", row))
	}
	return builder.String()
}

func getMove(grid grid, reader *bufio.Reader) (int, int) {
	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		coords := strings.Split(text, ",")
		if len(coords) != 2 {
			fmt.Println("invalid value, enter a comma separated list of coords")
			fmt.Println("e.g. '1,1' ")
			continue
		}
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			fmt.Println("invalid value, couldn't parse one of your coords to int: " + err.Error())
			continue
		}
		y, err2 := strconv.Atoi(coords[1])
		if err2 != nil {
			fmt.Println("invalid value, couldn't parse one of your coords to int: " + err2.Error())
			continue
		}
		if !(x < gridSize && x >= 0) || !(y < gridSize && y >= 0) {
			fmt.Println("coords out of range, try again ")
			continue
		}

		if grid[x][y] != 0 {
			fmt.Println("illegal move, space " + text + " already controlled by player " + strconv.Itoa(grid[x][y]))
			continue
		}

		return x, y
	}
}

func isWin(ownedCoords [][2]int, winAmount int, gridSize int) bool {
	count := len(ownedCoords)
	if count < winAmount {
		return false
	}
	var thingy = make([][4]int, gridSize*2)
	ownsMiddle := false
	//checks straight line wins
	for i := 0; i < count; i++ {
		ree := ownedCoords[i]
		if ree[0] == 1 && ree[1] == 1 {
			ownsMiddle = true
		}
		thingy[ree[0]][0] += 1
		thingy[ree[1]][1] += 1
		diagPlus := ree[0] + ree[1]
		//normalised diagMinus so as to not negative index the array
		diagMinus := ree[0] - ree[1] + 2
		thingy[diagPlus][2] += 1
		thingy[diagMinus][3] += 1
		if thingy[ree[0]][0] >= winAmount || thingy[ree[1]][1] >= winAmount {
			return true
		}
		if thingy[diagPlus][2] >= winAmount || thingy[diagMinus][3] >= winAmount {
			println("diag victory")
			return true
		}
	}
	//checks diagonal wins
	if ownsMiddle {
		for i := 0; i < count; i++ {

		}
	}
	return false
}

func checkDone(grid grid) (bool, string) {

	full := true
	winner := 0
	var player1 [][2]int
	var player2 [][2]int
	for x := 0; x < gridSize; x++ {

		for y := 0; y < gridSize; y++ {
			switch grid[x][y] {
			case 0:
				full = false
			case 1:
				player1 = append(player1, [2]int{x, y})
			case 2:
				player2 = append(player2, [2]int{x, y})
			}
		}
	}
	player1Won := isWin(player1, winAmount, gridSize)
	if player1Won {
		winner = 1
	}
	player2Won := isWin(player2, winAmount, gridSize)
	if player2Won {
		winner = 2
	}
	if winner != 0 {
		return true, "Player " + strconv.Itoa(winner) + " won"
	}
	if full {
		return true, "Draw, Game board full"
	}
	return false, "not done"
}

func doTurn(grid grid, player int, reader *bufio.Reader) (grid, bool) {
	x, y := getMove(grid, reader)
	grid[x][y] = player
	done, reason := checkDone(grid)
	if done {
		println(reason)
	}
	return grid, done
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Set Grid Size -> ")

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		size, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("invalid value, couldn't parse to int: " + err.Error())
			continue
		}
		if size < 0 {
			fmt.Println("size must be positive")
			continue
		}
		gridSize = size
		winAmount = size
		break
	}
	// for {
	// 	fmt.Print("Set Win Amount -> ")

	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to LF
	// 	text = strings.Replace(text, "\n", "", -1)
	// 	text = strings.Replace(text, "\r", "", -1)
	// 	size, err := strconv.Atoi(text)
	// 	if err != nil {
	// 		fmt.Println("invalid value, couldn't parse to int: " + err.Error())
	// 		continue
	// 	}
	// 	if size > gridSize || size < 0 {
	// 		fmt.Println("invalid value, game must be winnable")
	// 		continue
	// 	}
	// 	winAmount = size
	// 	break
	// }
	player := 1
	var grid = create2DArray(gridSize)
	done := false
	printGrid(grid)
	for !done {

		grid, done = doTurn(grid, player, reader)
		printGrid(grid)
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}
}
