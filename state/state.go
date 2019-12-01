package state

import "fmt"

// State is where the robot is
type State struct {
	x         int
	y         int
	direction string
	grid      grid
	robotLost bool
}

// NewState creates a state for a robot to run within
func NewState(gridHeight int, gridWidth int, initialXPossition int, initialYPossition int, initialDirection string) State {
	g := grid{gridWidth, gridHeight}
	state := State{
		initialXPossition,
		initialYPossition,
		initialDirection,
		g,
		false,
	}
	state.robotLost = state.grid.hasRobotFallenOff(initialXPossition, initialYPossition)
	return state
}

// MoveForward move robot forward
func (s *State) MoveForward() {
	if s.robotLost {
		return
	}
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

// ReportPosition returns string of current possition
func (s *State) ReportPosition() string {
	if s.robotLost {
		return fmt.Sprintf("%d %d %s LOST", s.x, s.y, s.direction)
	}
	return fmt.Sprintf("%d %d %s", s.x, s.y, s.direction)
}
