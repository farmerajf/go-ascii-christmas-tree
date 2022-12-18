package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"
)

var (
	height *int
	lights *int
	flash  *bool
)

func main() {
	height = flag.Int("height", 10, "The height of the tree")
	lights = flag.Int("lights", 5, "The spacing for the lights")
	flash = flag.Bool("flash", false, "Turn on flashing lights")
	flag.Parse()

	fmt.Print(generateTree(*height, *lights, *flash))
}

func generateTree(height int, lights int, flash bool) string {
	tree := ""
	t := &tree
	lc := 0
	for row := 1; row <= height; row++ {
		p(t, rep(" ", height-row))
		if row == 1 {
			pc(t, "*", YELLOW)
		} else {
			for col := 0; col < (row*2)-1; col++ {
				if lc%lights == 0 {
					p(t, randc(flash)+"*")
				} else {
					pc(t, "*", GREEN)
				}
				lc++
			}
		}
		p(t, "\n")
	}
	for tr := 0; tr < height/3; tr++ {
		p(t, rep(" ", (((height*2)-3)/2)))
		pc(t, "|||\n", RED)
	}

	return tree
}

func rep(char string, count int) string {
	r := ""
	for c := 0; c < count; c++ {
		r = r + char
	}
	return r
}

func p(input *string, char string) {
	*input = *input + char
}

func pc(input *string, char string, colour string) {
	*input = *input + colour + char + RESET
}

func randc(flash bool) string {
	f := "m"
	if flash {
		f = ";5m"
	}
	r := rn(6) + 31
	return "\033[" + fmt.Sprint(r) + f
}

func rn(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max)
}
