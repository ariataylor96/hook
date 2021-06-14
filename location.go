package main

type Location struct {
	X, Y int
	Grid *[][]string
}

func (loc Location) Bounds() (x, y int) {
	x = len(*loc.Grid) - 1
	y = len((*loc.Grid)[0]) - 1

	return
}

func (loc *Location) Move(vel Velocity) {
	x_max, y_max := loc.Bounds()

	loc.X += vel.X
	loc.Y += vel.Y

	if loc.X > x_max {
		loc.X = 0
	}

	if loc.Y > y_max {
		loc.Y = 0
	}

	if loc.X < 0 {
		loc.X = x_max
	}

	if loc.Y < 0 {
		loc.Y = y_max
	}
}
