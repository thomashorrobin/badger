package main

import "testing"

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
