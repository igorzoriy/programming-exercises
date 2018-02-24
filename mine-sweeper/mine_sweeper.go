package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeArea(file *os.File) [][]string {
	sb := make([]byte, 1)

	x, y := 0, 0
	area := make([][]string, 0)

	for {
		b, err := file.Read(sb)
		if b == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		if x == 0 {
			area = append(area, []string{})
		}

		str := string(sb)
		if str == "O" || str == "X" {
			area[y] = append(area[y], str)
			x++
		} else if str == "\n" {
			y++
			x = 0
		}
	}

	return area
}

func calculateMines(area [][]string) {
	for y, line := range area {
		for x, field := range line {
			if field != "O" {
				continue
			}

			count := 0

			if y-1 >= 0 && x-1 >= 0 && area[y-1][x-1] == "X" {
				count++
			}
			if y-1 >= 0 && area[y-1][x] == "X" {
				count++
			}
			if y-1 >= 0 && x+1 < len(line) && area[y-1][x+1] == "X" {
				count++
			}

			if x-1 >= 0 && area[y][x-1] == "X" {
				count++
			}
			if x+1 < len(line) && area[y][x+1] == "X" {
				count++
			}

			if y+1 < len(area) && x-1 >= 0 && area[y+1][x-1] == "X" {
				count++
			}
			if y+1 < len(area) && area[y+1][x] == "X" {
				count++
			}
			if y+1 < len(area) && x+1 < len(line) && area[y+1][x+1] == "X" {
				count++
			}

			area[y][x] = strconv.Itoa(count)
		}
	}
}

func main() {
	file, err := os.Open("./mine-sweeper.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	area := makeArea(file)
	calculateMines(area)

	for _, line := range area {
		fmt.Println(line)
	}
}
