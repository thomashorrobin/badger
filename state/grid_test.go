package state

import (
	"testing"
)

func TestRobotOnTheGrid(t *testing.T) {
	grid := grid{5, 3}
	testCases := []struct {
		desc string
		x    int
		y    int
	}{
		{
			desc: "top right corner",
			x:    5,
			y:    3,
		},
		{
			desc: "bottom left corner",
			x:    0,
			y:    0,
		},
		{
			desc: "centre",
			x:    3,
			y:    2,
		},
		{
			desc: "top edge",
			x:    3,
			y:    3,
		},
		{
			desc: "left edge",
			x:    0,
			y:    2,
		},
		{
			desc: "right edge",
			x:    5,
			y:    2,
		},
		{
			desc: "bottom edge",
			x:    3,
			y:    0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			robotHasFallenOffTheGrid := grid.hasRobotFallenOff(tC.x, tC.y)
			if robotHasFallenOffTheGrid {
				t.Error("robot was on the grid but HasRobotFallenOff said it wasn't")
			}
		})
	}
}

func TestRobotOffTheGrid(t *testing.T) {
	grid := grid{5, 3}
	testCases := []struct {
		desc string
		x    int
		y    int
	}{
		{
			desc: "above",
			x:    3,
			y:    4,
		},
		{
			desc: "below",
			x:    3,
			y:    -1,
		},
		{
			desc: "left",
			x:    -1,
			y:    2,
		},
		{
			desc: "right",
			x:    6,
			y:    2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			robotHasFallenOffTheGrid := grid.hasRobotFallenOff(tC.x, tC.y)
			if !robotHasFallenOffTheGrid {
				t.Error("robot was off the grid but HasRobotFallenOff said it wasn")
			}
		})
	}
}
