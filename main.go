package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"./state"

	flags "github.com/jessevdk/go-flags"
)

var opts options

func main() {
	flags.Parse(&opts)

	scanner := bufio.NewScanner(os.Stdin)

	gridWidth, gridHeight, err1 := askGridSize(scanner)

	if err1 != nil {
		fmt.Println(1)
		return
	}

	for true {
		x, y, dir, err3 := askInitialPossition(scanner)
		if err3 != nil {
			fmt.Println(err3.Error())
			return
		}
		s := state.NewState(*gridHeight, *gridWidth, *x, *y, *dir)
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
