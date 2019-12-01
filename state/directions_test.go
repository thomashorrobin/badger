package state

import (
	"testing"
)

func TestNorth(t *testing.T) {
	value := North
	if value != "N" {
		t.Error("enum did not return expected char")
	}
}

func TestSouth(t *testing.T) {
	value := South
	if value != "S" {
		t.Error("enum did not return expected char")
	}
}

func TestEast(t *testing.T) {
	value := East
	if value != "E" {
		t.Error("enum did not return expected char")
	}
}

func TestWest(t *testing.T) {
	value := West
	if value != "W" {
		t.Error("enum did not return expected char")
	}
}
func TestLeft(t *testing.T) {
	value := Left
	if value != "L" {
		t.Error("enum did not return expected char")
	}
}

func TestRight(t *testing.T) {
	value := Right
	if value != "R" {
		t.Error("enum did not return expected char")
	}
}

func TestForward(t *testing.T) {
	value := Forward
	if value != "F" {
		t.Error("enum did not return expected char")
	}
}
