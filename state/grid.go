package state

type grid struct {
	x int
	y int
}

func (g *grid) hasRobotFallenOff(x int, y int) bool {
	return x < 0 || y < 0 || y >= g.y || x >= g.x
}
