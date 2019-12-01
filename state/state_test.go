package state

import "testing"

func TestRobotMove(t *testing.T) {
	state := NewState(3, 3, 2, 2, East)
	state.MoveForward()
	if !(state.x == 3 && state.y == 2) {
		t.Error("robot has not moved to the correct posstion")
	}
}

func TestRobotMoveTwice(t *testing.T) {
	state := NewState(3, 5, 2, 2, East)
	state.MoveForward()
	state.MoveForward()
	if !(state.x == 4 && state.y == 2) {
		t.Error("robot has not moved to the correct posstion")
	}
}

func TestReportPossition(t *testing.T) {
	state := NewState(3, 5, 3, 1, West)
	possition := state.ReportPosition()
	if possition != "3 1 W" {
		t.Error("reported location was not correct")
	}
}

func TestReportPossitionLost(t *testing.T) {
	state := NewState(2, 2, 3, 3, West)
	possition := state.ReportPosition()
	if possition != "3 3 W LOST" {
		t.Error("location was not reported as lost")
	}
}

func TestReportRotateLeft(t *testing.T) {
	state := NewState(3, 3, 2, 2, North)
	state.RotateLeft()
	if state.ReportPosition() != "2 2 W" {
		t.Error("direction was not updated correctly")
	}
}

func TestReportRotateRight(t *testing.T) {
	state := NewState(3, 3, 2, 2, North)
	state.RotateRight()
	if state.ReportPosition() != "2 2 E" {
		t.Error("direction was not updated correctly")
	}
}
