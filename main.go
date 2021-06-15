package main

import (
	"fmt"
	"hook/parser"
	"os"
	"strconv"
)

func main() {
	var grid [][]string
	var stacks []*Stack

	loc := Location{0, 0, &grid}
	vel := Velocity{1, 0}

	var handle *os.File

	if len(os.Args) >= 2 {
		handle, _ = os.Open(os.Args[1])
	} else {
		panic("Please give a filename containing your fish program")
	}

	defer handle.Close()
	grid = parser.Tokenize(handle)

	stacks = append(stacks, NewStack())

	var (
		running       = true
		skip          = false
		string_mode   = false
		string_opener = "'"
	)

	for running {
		sym := grid[loc.X][loc.Y]
		stack := stacks[len(stacks)-1]

		if skip {
			skip = false
			loc.Move(vel)
			continue
		}

		if string_mode {
			switch sym {
			case string_opener:
				string_mode = false
			default:
				stack.Push(int(sym[0]))
			}

			loc.Move(vel)
			continue
		}

		switch sym {
		case ";":
			running = false

			// Directions
		case "v":
			vel = Velocity{0, 1}
		case "^":
			vel = Velocity{0, -1}
		case "<":
			vel = Velocity{-1, 0}
		case ">":
			vel = Velocity{1, 0}

			// Stack allocation
		case "[":
			new_stack := stack.Split()

			stacks = append(stacks, new_stack)
		case "]":
			previous_stack := stacks[len(stacks)-2]
			previous_stack.Join(stack)

			stacks = stacks[0 : len(stacks)-1]

			// Mirrors
		case "/":
			if vel.X != 0 {
				vel.Y = vel.X * -1
				vel.X = 0
			} else if vel.Y != 0 {
				vel.X = vel.Y * -1
				vel.Y = 0
			}
		case "\\":
			if vel.X != 0 {
				vel.Y = vel.X * -1
				vel.X = 0
			} else if vel.Y != 0 {
				vel.X = vel.Y
				vel.Y = 0
			}
		case "_":
			vel.Y *= -1
		case "|":
			vel.X *= -1
		case "#":
			vel.X *= -1
			vel.Y *= -1

			// Jumps
		case "!":
			skip = true
		case "?":
			val := stack.Pop()

			if val == 0 {
				skip = true
			}
		case ".":
			loc.X, loc.Y = stack.Pop(), stack.Pop()

			// Operators
		case "+":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a + b)
		case "-":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a - b)
		case "*":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a * b)
		case ",":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a / b)
		case "%":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a % b)
		case "=":
			a, b := stack.Pop(), stack.Pop()

			if a == b {
				stack.Push(1)
			} else {
				stack.Push(0)
			}
		case ")":
			a, b := stack.Pop(), stack.Pop()

			if b > a {
				stack.Push(1)
			} else {
				stack.Push(0)
			}
		case "(":
			a, b := stack.Pop(), stack.Pop()

			if b < a {
				stack.Push(1)
			} else {
				stack.Push(0)
			}

		// More stack manipulation
		case ":":
			val := stack.Pop()
			stack.Push(val)
			stack.Push(val)
		case "~":
			_ = stack.Pop()
		case "$":
			a, b := stack.Pop(), stack.Pop()
			stack.Push(a)
			stack.Push(b)
		case "@":
			vals := []int{stack.Pop(), stack.Pop(), stack.Pop()}
			reordered := []int{vals[2], vals[0], vals[1]}

			for _, v := range reordered {
				stack.Push(v)
			}
		case "{":
			stack.Reverse()
			vals := []int{}

			for stack.Count > 0 {
				vals = append(vals, stack.Pop())
			}

			reordered := append(vals[1:], vals[0])

			for _, val := range reordered {
				stack.Push(val)
			}

		case "}":
			stack.Reverse()
			vals := []int{}

			for stack.Count > 0 {
				vals = append(vals, stack.Pop())
			}

			reordered := []int{vals[len(vals)-1]}
			reordered = append(reordered, vals[:len(vals)-1]...)

			for _, val := range reordered {
				stack.Push(val)
			}

		case "l":
			stack.Push(stack.Count)

		case "r":
			stack.Reverse()
		case "o":
			fmt.Printf("%c", rune(stack.Pop()))
		case "n":
			fmt.Print(stack.Pop())
		case "x":
			fmt.Println(len(stack.Nodes))

			// Strings
		case "'":
			fallthrough
		case "\"":
			string_mode = true
			string_opener = sym

		case "g":
			x, y := stack.Pop(), stack.Pop()
			i, _ := strconv.ParseInt(grid[x][y], 16, 32)
			stack.Push(int(i))

		case "p":
			x, y, v := stack.Pop(), stack.Pop(), stack.Pop()

			grid[x][y] = strconv.Itoa(v)

		default:
			i, err := strconv.ParseInt(sym, 16, 32)
			if err == nil {
				stack.Push(int(i))
			}
		}

		loc.Move(vel)
	}
}
