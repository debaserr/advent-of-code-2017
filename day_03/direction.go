package main

//import "fmt"

const DIR_RIGHT string = "right"
const DIR_LEFT string = "left"
const DIR_UP string = "up"
const DIR_DOWN string = "down"

type Direction interface {
	TurnLeft()
	GetName() string
	GetY() int
	GetX() int
}

type direction struct {
	name string
	x int
	y int
}

func (d *direction) GetName() string {
	return d.name
}

func (d *direction) GetY() int {
	return d.y
}

func (d *direction) GetX() int {
	return d.x
}

func (d *direction) TurnLeft() {
	//fmt.Println("in TurnLeft")
	if d.name == DIR_RIGHT {
		d.name = DIR_UP
		d.x = 0
		d.y = -1
	} else if d.name == DIR_UP {
		d.name = DIR_LEFT
		d.x = -1
		d.y = 0
	} else if d.name == DIR_LEFT {
		d.name = DIR_DOWN
		d.x = 0
		d.y = 1
	} else if d.name == DIR_DOWN {
		d.name = DIR_RIGHT
		d.x = 1
		d.y = 0
	}
}

func DirectionRight() Direction {
	return &direction{
		DIR_RIGHT,
		1,
		0,
	}
}
