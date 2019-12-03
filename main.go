package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./state"

	flags "github.com/jessevdk/go-flags"
)

var opts options

func main() {
	flags.Parse(&opts)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("grid size: ")
	scanner.Scan()
	gridSize := strings.Split(scanner.Text(), " ")
	gridWidth, _ := strconv.Atoi(gridSize[0])
	gridHeight, _ := strconv.Atoi(gridSize[1])

	fmt.Print("initial possition: ")
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		s := state.NewState(gridHeight, gridWidth, x, y, split[2])
		if opts.Debug {
			fmt.Println(s.ReportPosition())
		}
		fmt.Print("movement commands: ")
		scanner.Scan()
		for _, char := range scanner.Text() {
			err := move(string(char), &s)
			if opts.Debug {
				fmt.Println(s.ReportPosition())
			}
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Println("final possition: ", s.ReportPosition())
		fmt.Print("initial possition: ")
	}
}

type options struct {
	Debug bool `long:"debug"`
}

func move(m string, s *state.State) error {
	switch m {
	case Forward:
		s.MoveForward()
		return nil
	case Left:
		s.RotateLeft()
		return nil
	case Right:
		s.RotateRight()
		return nil
	default:
		return errors.New("incorrect char")
	}
}
