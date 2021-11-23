package main

import (
	"fmt"
	"time"
)

type placeholder [5]string

var (
	zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	siven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	empty = placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	digits = []placeholder{zero, one, two, three, four, five, six, siven, eight, nine}
)

func main() {
	for {
		printTime()
		fmt.Println()

		time.Sleep(time.Second)
	}
}

func printTime() {
	ct := time.Now()
	h, m, s := ct.Hour(), ct.Minute(), ct.Second()

	clock := []placeholder{
		digits[h/10], digits[h%10],
		colon,
		digits[m/10], digits[m%10],
		colon,
		digits[s/10], digits[s%10],
	}

	ClearScreen()
	printDigitArray(clock)
	MoveCursorTopLeft()
}

func printDigitArray(digits []placeholder) {
	for i := 0; i < 5; i++ {
		for _, digit := range digits {
			if digit == colon && time.Now().Second()%2 == 0 {
				digit = empty
			}
			fmt.Print(digit[i])
			fmt.Print("  ")
		}
		fmt.Println()
	}
}

func ClearScreen() {
	fmt.Print("\033[2J")
}

func MoveCursorTopLeft() {
	fmt.Print("\033[H")
}
