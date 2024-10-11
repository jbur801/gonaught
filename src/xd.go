package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grid = [3][3]int

func printGrid(grid grid) {
	for x := 0; x < 3; x++ {
		line := ""
		for y := 0; y < 3; y++ {
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
		if !(x < 3 && x >= 0) || !(y < 3 && y >= 0) {
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

func isWin(ownedCoords [][2]int) bool {
	count := len(ownedCoords)
	if count < 3 {
		return false
	}
	var thingy [5][4]int
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
		if thingy[ree[0]][0] > 2 || thingy[ree[1]][1] > 2 {
			return true
		}
		if thingy[diagPlus][2] > 2 || thingy[diagMinus][3] > 2 {
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
	for x := 0; x < 3; x++ {

		for y := 0; y < 3; y++ {
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
	player1Won := isWin(player1)
	if player1Won {
		winner = 1
	}
	player2Won := isWin(player2)
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
	var grid [3][3]int
	reader := bufio.NewReader(os.Stdin)
	player := 1
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
