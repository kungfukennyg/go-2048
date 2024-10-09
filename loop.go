package main

import (
	"bufio"
	"fmt"
	"os"
)

var controls map[rune]Dir = map[rune]Dir{
	'w': DirUp,
	'a': DirLeft,
	's': DirDown,
	'd': DirRight,
}

func Start(b *Board) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\t   2048")
	fmt.Println(b)
	showPrompt := true
	for {
		if showPrompt {
			fmt.Print("(WASD): ")
			showPrompt = false
		}
		r, _, err := reader.ReadRune()
		if err != nil {
			continue
		}
		dir, ok := controls[r]
		if !ok {
			fmt.Println()
			continue
		}

		b.swipe(dir)
		fmt.Println(b)
		showPrompt = true
	}
}
