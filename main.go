package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const theSum = 45

type field [][]int

func joinLine(items []int, sep string) string {
	line := strings.Builder{}
	for _, item := range items {
		line.WriteString(sep)
		line.WriteString(fmt.Sprintf("%d", item))
	}

	line.WriteString(sep)
	return line.String()
}

func sumInts(ints ...int) (sum int) {
	for _, v := range ints {
		sum += v
	}
	return
}

func (f field) valid() bool {
	for i, line := range f {
		s := sumInts(line...)
		if s != theSum {
			fmt.Printf("line: %d: %v has sum: %d\n", i, line, s)
			return false
		}

		for range line {
			column := f[i][0:]
			s := sumInts(column...)
			if s != theSum {
				fmt.Printf("column: %d: %v has sum: %d\n", i, line, s)
				return false
			}
		}
	}

	return true
}

func (f field) String() string {
	table := strings.Builder{}
	for _, line := range f {

		fline := joinLine(line, " │ ")
		table.WriteString(fline + "\n")
	}

	return table.String()
}

func (f field) done() bool {
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			if f[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

// type constraints struct {
// 	x, y          int
// 	possibilities []int
// }

func fillOutSolvedField() field {
	return field{
		{7, 1, 6, 5, 8, 4, 3, 9, 2},
		{5, 4, 9, 3, 1, 2, 8, 6, 7},
		{8, 2, 3, 6, 9, 7, 1, 5, 4},
		{6, 9, 5, 8, 7, 3, 4, 2, 1},
		{1, 7, 2, 4, 6, 5, 9, 3, 8},
		{3, 8, 4, 9, 2, 1, 5, 7, 6},
		{9, 6, 7, 1, 5, 8, 2, 4, 3},
		{2, 3, 8, 7, 4, 9, 6, 1, 5},
		{4, 5, 1, 2, 3, 6, 7, 8, 9},
	}
}

func easyField() field {
	return field{
		{0, 0, 3, 0, 2, 0, 6, 0, 0},
		{9, 0, 0, 3, 0, 5, 0, 0, 1},
		{0, 0, 1, 8, 0, 6, 4, 0, 0},
		{0, 0, 8, 1, 0, 2, 9, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 6, 7, 0, 8, 2, 0, 0},
		{0, 0, 2, 6, 0, 9, 5, 0, 0},
		{8, 0, 0, 2, 0, 3, 0, 0, 9},
		{0, 0, 5, 0, 1, 0, 3, 0, 0},
	}
}

func fillOutNewField() field {
	return field{
		{0, 0, 0, 0, 0, 4, 0, 9, 0},
		{0, 4, 0, 0, 0, 0, 0, 6, 0},
		{8, 2, 0, 0, 0, 7, 1, 0, 0},
		{0, 9, 5, 0, 7, 3, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 9, 2, 0, 5, 7, 0},
		{0, 0, 7, 1, 0, 0, 0, 4, 3},
		{0, 3, 0, 0, 0, 0, 0, 1, 0},
		{0, 5, 0, 2, 0, 0, 0, 0, 0},
	}
}

var areas = map[string]string{
	"0,0": "a",
	"0,1": "a",
	"0,2": "a",
	"0,3": "d",
	"0,4": "d",
	"0,5": "d",
	"0,6": "g",
	"0,7": "g",
	"0,8": "g",
	"1,0": "a",
	"1,1": "a",
	"1,2": "a",
	"1,3": "d",
	"1,4": "d",
	"1,5": "d",
	"1,6": "g",
	"1,7": "g",
	"1,8": "g",
	"2,0": "a",
	"2,1": "a",
	"2,2": "a",
	"2,3": "d",
	"2,4": "d",
	"2,5": "d",
	"2,6": "g",
	"2,7": "g",
	"2,8": "g",
	"3,0": "b",
	"3,1": "b",
	"3,2": "b",
	"3,3": "e",
	"3,4": "e",
	"3,5": "e",
	"3,6": "h",
	"3,7": "h",
	"3,8": "h",
	"4,0": "b",
	"4,1": "b",
	"4,2": "b",
	"4,3": "e",
	"4,4": "e",
	"4,5": "e",
	"4,6": "h",
	"4,7": "h",
	"4,8": "h",
	"5,0": "b",
	"5,1": "b",
	"5,2": "b",
	"5,3": "e",
	"5,4": "e",
	"5,5": "e",
	"5,6": "h",
	"5,7": "h",
	"5,8": "h",
	"6,0": "c",
	"6,1": "c",
	"6,2": "c",
	"6,3": "f",
	"6,4": "f",
	"6,5": "f",
	"6,6": "i",
	"6,7": "i",
	"6,8": "i",
	"7,0": "c",
	"7,1": "c",
	"7,2": "c",
	"7,3": "f",
	"7,4": "f",
	"7,5": "f",
	"7,6": "i",
	"7,7": "i",
	"7,8": "i",
	"8,0": "c",
	"8,1": "c",
	"8,2": "c",
	"8,3": "f",
	"8,4": "f",
	"8,5": "f",
	"8,6": "i",
	"8,7": "i",
	"8,8": "i",
}

func convertLocationToInt(location string) (x, y int) {
	values := strings.Split(location, ",")
	x, _ = strconv.Atoi(values[0])
	y, _ = strconv.Atoi(values[1])

	return
}

func getKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func findNeighborValues(i, j int, field field) []int {
	key := getKey(i, j)
	a := areas[key]

	neighbors := make([]int, 0)
	for location, area := range areas {
		if area == a {
			k, l := convertLocationToInt(location)
			n := field[k][l]
			if n > 0 {
				neighbors = append(neighbors, n)
			}
		}
	}

	return neighbors
}

func findNeighborPossibilities(i, j int, field field) []int {
	neighbors := findNeighborValues(i, j, field)
	return getPossibilities(neighbors, "area")
}

func getPossibilities(slice []int, name string) []int {
	present := map[int]bool{}
	for _, v := range slice {
		if v != 0 {
			present[v] = true
		}
	}

	// fmt.Printf("%s: %v\n", name, present)
	possibilities := []int{}
	for v := 1; v <= 9; v++ {
		if present[v] {
			continue
		}
		possibilities = append(possibilities, v)
	}

	return possibilities
}

func findColumnPossibilities(c int, field field) []int {
	column := []int{}
	for _, line := range field {
		column = append(column, line[c])
	}

	return getPossibilities(column, "column")
}

func findLinePossibilities(l int, field field) []int {
	line := field[l]
	return getPossibilities(line, "line")
}

// func propagateToPeers(key string, value int, poss map[string][]int) map[string][]int {

// }

func intersection(a, b, c []int) (inter []int) {
	u := append(a, b...)
	u = append(u, c...)
	sort.Ints(u)

	unique := map[int]int{}
	for _, i := range u {
		c := unique[i]
		unique[i] = c + 1
	}

	for k, c := range unique {
		if c == 3 {
			inter = append(inter, k)
		}
	}

	sort.Ints(inter)
	return
}

func main() {
	// solved := easyField()
	// fmt.Printf("is field valid: %t\n", solved.valid())

	field := easyField()
	fmt.Println(field)

	// pl := findLinePossibilities(0, f)
	// fmt.Printf("line 2: %+v\n", pl)

	// pc := findColumnPossibilities(0, f)
	// fmt.Printf("column 1: %+v\n", pc)

	// pn := findNeighborPossibilities(0, 6, field)
	// fmt.Printf("neighborhood of (0,6): %+v\n", pn)

	// fmt.Printf("inter: %+v\n", intersection(pl, pc, pn))

	empties := 0
	possibilities := map[string][]int{}

	for loops := 0; loops < 100; loops++ {
		for i := 0; i < 9; i++ {
			l := findLinePossibilities(i, field)
			for j := 0; j < 9; j++ {
				if field[i][j] != 0 {
					continue
				}
				empties++

				c := findColumnPossibilities(j, field)
				n := findNeighborPossibilities(i, j, field)
				u := intersection(l, c, n)
				possibilities[getKey(i, j)] = u
			}
		}

		for key, values := range possibilities {
			if len(values) == 1 {
				x, y := convertLocationToInt(key)
				field[x][y] = values[0]
				// if field[x][y] != solved[x][y] {
				// fmt.Println("value does not match solution")
				// break
				// }
				delete(possibilities, key)
				// fmt.Println("removed", x, y, values)
			}
		}
		fmt.Println("empties", empties)
		empties = 0

		if field.done() {
			fmt.Println("solved puzzle")
			fmt.Println("solution is valid?", field.valid())
			break
		}

	}
	for key, values := range possibilities {
		fmt.Printf("%s: %+v\n", key, values)
	}
}
