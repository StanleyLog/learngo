package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {

	file, err := os.Open(filename)
	if nil != err {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

func (p point) at(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {

	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	q := []point{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)
			if !ok || val != 0 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			q = append(q, next)

		}
	}
	return steps
}

func main() {

	maze := readMaze("/Users/stanley/Documents/GitHub/learngo/maze/maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, i := range steps {
		for _, j := range i {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}

	// fmt.Println(maze)
}
