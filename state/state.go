package state

// State is where the robot is
type State struct {
	x         int
	y         int
	direction string
	grid      grid
}

// NewState creates a state for a robot to run within
func NewState(gridHeight int, gridWidth int, initialXPossition int, initialYPossition int, initialDirection string) State {
	g := grid{gridWidth, gridHeight}
	return State{
		initialXPossition,
		initialYPossition,
		initialDirection,
		g,
	}
}

// MoveForward move robot forward
func (s *State) MoveForward() {
	switch s.direction {
	case North:
		s.y++
		break
	case South:
		s.y--
		break
	case West:
		s.x--
		break
	case East:
		s.x++
		break
	}
}
