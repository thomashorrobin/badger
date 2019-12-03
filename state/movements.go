package state

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
	s.robotLost = s.grid.hasRobotFallenOff(s.x, s.y)
}

// RotateLeft rotate the robot left
func (s *State) RotateLeft() {
	if s.robotLost {
		return
	}
	switch s.direction {
	case North:
		s.direction = West
		break
	case South:
		s.direction = East
		break
	case West:
		s.direction = South
		break
	case East:
		s.direction = North
		break
	}
}

// RotateRight rotate the robot right
func (s *State) RotateRight() {
	if s.robotLost {
		return
	}
	switch s.direction {
	case North:
		s.direction = East
		break
	case South:
		s.direction = West
		break
	case West:
		s.direction = North
		break
	case East:
		s.direction = South
		break
	}
}
