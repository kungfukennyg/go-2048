package main

import "testing"

func TestSwipeOnlyMergesOneMatch(t *testing.T) {
	board := new()
	board.set(Point{0, 0}, 8)
	board.set(Point{0, 1}, 4)
	board.set(Point{0, 2}, 2)
	board.set(Point{0, 3}, 2)

	changed := board.swipe(DirUp)
	if !changed {
		t.Fatalf("expected board state to change after swipe up")
	}

	expect(t, board, Point{0, 0}, 8)
	expect(t, board, Point{0, 1}, 4)
	expect(t, board, Point{0, 2}, 4)
	expect(t, board, Point{0, 3}, 0)
}

func expect(t *testing.T, b Board, p Point, want int) {
	t.Helper()
	got := b.get(p)
	if got != want {
		t.Fatalf("expected point (%d, %d) to be %d, got %d", p.x, p.y, want, got)
	}
}
