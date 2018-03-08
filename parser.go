package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string, t *Matrix, e *Matrix, image *Image) error {
	f, err := os.Open(filename)
	if err != nil {
		return errors.New("Couldn't open file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		switch c := strings.TrimSpace(scanner.Text()); c {
		case "line":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 6 {
				fmt.Errorf("Line: Incorrect # of args. Got: %d, expected: 6\n", len(args))
			}
			fargs := numerize(args)
			e.AddEdge(fargs[0], fargs[1], fargs[2], fargs[3], fargs[4], fargs[5])

		case "ident":
			t.Ident()

		case "scale":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 3 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			fargs := numerize(args)
			scale := MakeScale(fargs[0], fargs[1], fargs[2])
			t, _ = t.Mult(scale)

		case "move":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 3 {
				fmt.Errorf("Translate: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			fargs := numerize(args)
			translate := MakeTranslate(fargs[0], fargs[1], fargs[2])
			t, _ = t.Mult(translate)

		case "rotate":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 2 {
				fmt.Errorf("Rotate: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			// TODO: Error handling
			deg, _ := strconv.ParseFloat(args[1], 64)
			fmt.Printf("rotating %s by %.2f\n", args[0], deg)
			switch args[0] {
			case "x":
				rot := MakeRotX(deg)
				fmt.Println(rot)
				t, _ = t.Mult(rot)
			case "y":
				rot := MakeRotY(deg)
				fmt.Println(rot)
				t, _ = t.Mult(rot)
			case "z":
				rot := MakeRotZ(deg)
				fmt.Println(rot)
				t, _ = t.Mult(rot)
			default:
				// TODO: Error handling
				fmt.Println("Rotate fail")
				continue
			}

		case "apply":
			fmt.Println(t)
			// TODO: Error handling
			e, _ = e.Mult(t)

		case "display":
			image.Clear()
			image.DrawLines(e, Color{r: 0, b: 0, g: 255})
			image.Display()

		case "save":
			scanner.Scan()
			arg := strings.TrimSpace(scanner.Text())
			args := strings.Fields(arg)
			if len(args) != 1 {
				fmt.Errorf("Scale: Incorrect # of args. Got: %d, expected: 3\n", len(args))
			}
			image.Clear()
			image.DrawLines(e, Color{r: 0, b: 0, g: 255})
			image.SavePPM(args[0])

		case "quit":
			break

		default:
			return errors.New("Invalid command")
		}
	}
	return nil
}

// Error handling?
func numerize(args []string) []float64 {
	fargs := make([]float64, len(args))
	for i, val := range args {
		fargs[i], _ = strconv.ParseFloat(val, 64)
	}
	return fargs
}
