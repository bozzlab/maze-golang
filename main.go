package main

import (
	"fmt"
	"strconv"
	"container/list"
)

type Point struct {
	x,y int
}

type QueueNode struct {
	p Point
	distance int
}

type Maze struct {
	matrix      [][]int
	visited    [][]bool
	start, end Point
	queue      *list.List
}

func PrepareMaze(mazeMap [][]int, start, goal Point) Maze {
	m := Maze{
		matrix:   mazeMap,
		visited: make([][]bool, 10),
		start:   start,
		end:     goal,
		queue:   list.New(),
	}

	for idx := 0; idx < 10; idx++ { m.visited[idx] = make([]bool, 10) }
	m.visited[start.x][start.y] = true

	m.queue.PushBack(QueueNode{p: start, distance: 0})
	
	return m
}

func (m *Maze) FindShortestPath() int {
	for m.queue.Len() != 0 {
		mazeQ := m.queue.Front()
		var qn QueueNode
		if mazeQ != nil {
			qn = mazeQ.Value.(QueueNode)
			m.queue.Remove(mazeQ)
		}
		if qn.p.x == m.end.x && qn.p.y == m.end.y {
			return qn.distance
		}
		m.findNextPath(qn)
	}
	return 9999
}

func (m *Maze) findNextPath(qn QueueNode) {
	totalDistance := qn.distance + 1

	current := Point{qn.p.x, qn.p.y + 1} // up
	if current.y < 10 && m.matrix[current.x][current.y] < 1 && !m.visited[current.x][current.y] {
		m.queue.PushBack(QueueNode{current, totalDistance})
		m.visited[current.x][current.y] = true
	}

	current = Point{qn.p.x, qn.p.y - 1} // down
	if current.y >= 0 && m.matrix[current.x][current.y] < 1 && !m.visited[current.x][current.y] {
		m.queue.PushBack(QueueNode{current, totalDistance})
		m.visited[current.x][current.y] = true
	}	

	current = Point{qn.p.x - 1, qn.p.y} // left
	if current.x >= 0 && m.matrix[current.x][current.y] < 1 && !m.visited[current.x][current.y] {
		m.queue.PushBack(QueueNode{current, totalDistance})
		m.visited[current.x][current.y] = true
	}

	current = Point{qn.p.x + 1, qn.p.y} // right
	if current.x < 10 && m.matrix[current.x][current.y] < 1 && !m.visited[current.x][current.y] {
		m.queue.PushBack(QueueNode{current, totalDistance})
		m.visited[current.x][current.y] = true
	}
}

func main() {
	var input string
	var start Point
	var goal Point
	matrix := make([][]int, 10)

	for row := 0; row < 10; row++ {
		fmt.Scan(&input)
		for idx, val := range input { 
			if (string(val) == "S") {
				start.x = row
				start.y = idx
			}
			if (string(val) == "G") {
				goal.x = row
				goal.y = idx
			}

			num, _ := strconv.Atoi(string(val))
			matrix[row] = append(matrix[row], num)
		}
	}

	MazeObject := PrepareMaze(matrix, start, goal)
	shortestPath := MazeObject.FindShortestPath()
	fmt.Println(shortestPath)

}