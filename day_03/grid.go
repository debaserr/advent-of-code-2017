package main

import "fmt"

type Grid interface {
	Add() int
	GetCurrentValue() int
}

type grid struct {
	layout [][]int
	currX int
	currY int
	direction Direction
}

func (g *grid) PrintGrid() {
	fmt.Printf("currY: %d, currX: %d, direction: %s\n", g.currY, g.currX, g.direction.GetName())
	for y := 0; y < len(g.layout); y++ {
		for x := 0; x < len(g.layout[y]); x++ {
			fmt.Printf(" %d ", g.layout[y][x])
		}
		fmt.Printf("\n")
	}
}

// Increments the current position and sets its value
func (g *grid) incrementPosition() int {
	if g.atEdge() {
		if g.direction.GetName() == DIR_RIGHT {
			g.expandGrid()
			g.move()
			val := g.calculateCurrentPositionValue()
			g.direction.TurnLeft()
			return val
		}
		g.direction.TurnLeft()
	}
	g.move()
	return g.calculateCurrentPositionValue()
}

func (g *grid) expandGrid() {
	// Create new grid
	newSize := len(g.layout) + 2
	newLayout := make([][]int, newSize)
	for i := range newLayout {
		newLayout[i] = make([]int, newSize)
	}

	// Insert old grid in new grid
	for y := range g.layout {
		for x := range g.layout[y] {
			newLayout[y + 1][x + 1] = g.layout[y][x]
		}
	}

	g.layout = newLayout

	g.currY += 1
	g.currX += 1
}

func (g *grid) atEdge() bool {
	return (g.direction.GetName() == DIR_RIGHT && g.currY == len(g.layout) - 2 && g.currX == len(g.layout[g.currY]) - 2) ||
		(g.direction.GetName() == DIR_UP && g.currY == 1 && g.currX == len(g.layout[g.currY]) - 2) ||
		(g.direction.GetName() == DIR_LEFT && g.currY == 1 && g.currX == 1) ||
		(g.direction.GetName() == DIR_DOWN && g.currY == len(g.layout) - 2 && g.currX == 1)
}

func (g *grid) move() {
	g.currY += g.direction.GetY()
	g.currX += g.direction.GetX()
}

// Calculates the value at the current position,
// inserts it to layout and returns it
func (g *grid) calculateCurrentPositionValue() int {
	newVal := 0
	g.layout[g.currY][g.currX] = 0
	for y := g.currY - 1; y <= g.currY + 1; y++ {
		for x := g.currX - 1; x <= g.currX + 1; x++ {
			newVal += g.layout[y][x]
		}
	}
	g.layout[g.currY][g.currX] = newVal
	return newVal
}

func (g *grid) Add() int {
	g.incrementPosition()
	return g.layout[g.currY][g.currX]
}

func (g *grid) GetCurrentValue() int {
	return g.layout[g.currY][g.currX]
}

func GetGrid() Grid {
	return &grid{
		[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		1,
		1,
		DirectionRight(),
		}
}
