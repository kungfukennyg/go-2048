package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Dir string

const (
	DirLeft  Dir = Dir("left")
	DirRight Dir = Dir("right")
	DirUp    Dir = Dir("up")
	DirDown  Dir = Dir("down")
)

func (d Dir) transform(p Point) Point {
	switch d {
	case DirDown:
		return p.add(0, 1)
	case DirLeft:
		return p.add(-1, 0)
	case DirRight:
		return p.add(1, 0)
	case DirUp:
		return p.add(0, -1)
	default:
		panic(fmt.Sprintf("unexpected main.Dir: %#v", d))
	}
}

const SIZE = 4

type Board struct {
	squares [][]int
	dirty   bool
	score   int
}

type Point struct {
	x, y int
}

func (p Point) add(x, y int) Point {
	return Point{p.x + x, p.y + y}
}

func new() Board {
	b := Board{
		squares: make([][]int, SIZE),
	}

	for i := range SIZE {
		b.squares[i] = make([]int, SIZE)
	}

	// fill two random squares at start of game
	b.randomEmpty()
	b.randomEmpty()
	return b
}

func (b *Board) swipe(dir Dir) bool {
	b.dirty = false
	switch dir {
	case DirDown:
		for x := SIZE - 1; x >= 0; x-- {
			dirty := true
			foundMerge := false
			for dirty {
				dirty = false
				for y := SIZE - 1 - 1; y >= 0; y-- {
					cell := b.squares[y][x]
					if cell == 0 {
						continue
					}

					cur := Point{x, y}
					neighbor := dir.transform(cur)
					neighborCell := b.get(neighbor)
					if neighborCell == 0 {
						b.set(neighbor, cell)
						b.set(cur, 0)
						dirty = true
					} else if !foundMerge && neighborCell == cell {
						b.set(neighbor, cell*2)
						b.set(cur, 0)
						dirty = true
						foundMerge = true
					}
				}
			}
		}
	case DirUp:
		for x := 0; x < SIZE; x++ {
			dirty := true
			foundMerge := false
			for dirty {
				dirty = false
				for y := 1; y < SIZE; y++ {
					cell := b.squares[y][x]
					if cell == 0 {
						continue
					}

					cur := Point{x, y}
					neighbor := dir.transform(cur)
					neighborCell := b.get(neighbor)
					if neighborCell == 0 {
						b.set(neighbor, cell)
						b.set(cur, 0)
						dirty = true
					} else if !foundMerge && neighborCell == cell {
						b.set(neighbor, cell*2)
						b.set(cur, 0)
						dirty = true
						foundMerge = true
					}
				}
			}
		}
	case DirLeft:
		for y := 0; y < SIZE; y++ {
			row := b.squares[y]
			dirty := true
			foundMerge := false
			for dirty {
				dirty = false
				for x := 1; x < len(row); x++ {
					cell := row[x]
					if cell == 0 {
						continue
					}

					cur := Point{x, y}
					neighbor := dir.transform(cur)
					neighborCell := b.get(neighbor)
					if neighborCell == 0 {
						b.set(neighbor, cell)
						b.set(cur, 0)
						dirty = true
					} else if !foundMerge && neighborCell == cell {
						b.set(neighbor, cell*2)
						b.set(cur, 0)
						dirty = true
						foundMerge = true
					}
				}
			}
		}
	case DirRight:
		for y := SIZE - 1; y >= 0; y-- {
			row := b.squares[y]
			dirty := true
			foundMerge := false
			for dirty {
				dirty = false
				for x := SIZE - 1 - 1; x >= 0; x-- {
					cell := row[x]
					if cell == 0 {
						continue
					}

					cur := Point{x, y}
					neighbor := dir.transform(cur)
					neighborCell := b.get(neighbor)
					if neighborCell == 0 {
						b.set(neighbor, cell)
						b.set(cur, 0)
						dirty = true
					} else if !foundMerge && neighborCell == cell {
						b.set(neighbor, cell*2)
						b.set(cur, 0)
						dirty = true
						foundMerge = true
					}
				}
			}
		}
	default:
		panic(fmt.Sprintf("unexpected main.Dir: %#v", dir))
	}
	if b.dirty {
		b.randomEmpty()
		b.score++
	}
	return b.dirty
}

func (b *Board) get(p Point) int {
	return b.squares[p.y][p.x]
}

func (b *Board) set(p Point, v int) {
	b.squares[p.y][p.x] = v
	b.dirty = true
}

func (b *Board) randomEmpty() {
	foundEmpty := false
OUTER:
	for _, row := range b.squares {
		for _, elem := range row {
			if elem == 0 {
				foundEmpty = true
				break OUTER
			}
		}
	}
	if !foundEmpty {
		return
	}

	for {
		x, y := rand.Intn(SIZE), rand.Intn(SIZE)
		p := Point{x, y}
		if b.get(p) == 0 {
			b.set(p, 2)
			return
		}
	}
}

func (b *Board) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\tScore: %d\n", b.score))
	sb.WriteString("-------------------------------------\n")
	for _, row := range b.squares {
		for _, cell := range row {
			sb.WriteString(fmt.Sprintf("%6d ", cell))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("-------------------------------------")
	return sb.String()
}
