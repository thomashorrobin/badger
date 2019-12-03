package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func askGridSize(scanner *bufio.Scanner) (*int, *int, error) {
	fmt.Print("grid size: ")
	scanner.Scan()
	gridSize := strings.Split(scanner.Text(), " ")
	gridWidth, err1 := strconv.Atoi(gridSize[0])
	if err1 != nil {
		return nil, nil, err1
	}
	gridHeight, err2 := strconv.Atoi(gridSize[1])
	if err2 != nil {
		return nil, nil, err2
	}
	return &gridWidth, &gridHeight, nil
}

func askInitialPossition(scanner *bufio.Scanner) (*int, *int, *string, error) {
	fmt.Print("initial possition: ")
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	if len(split) < 3 {
		return nil, nil, nil, errors.New("wrong number of arguements")
	}
	xStartingPoint, err1 := strconv.Atoi(split[0])
	if err1 != nil {
		return nil, nil, nil, err1
	}
	yStartingPoint, err2 := strconv.Atoi(split[1])
	if err2 != nil {
		return nil, nil, nil, err2
	}
	return &xStartingPoint, &yStartingPoint, &split[2], nil
}
