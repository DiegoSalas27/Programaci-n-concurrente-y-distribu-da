package main

import "fmt"

// Game is the game structure
type Game struct {
	Pieces []rune
	Turn   int
	Tab    [][]rune
}

// Move returns true when a move is made
func (g *Game) Move(i, j int) bool {
	if g.Tab[i][j] != ' ' {
		return false
	}
	g.Tab[i][j] = g.Pieces[g.Turn]
	g.Turn = (g.Turn + 1) % len(g.Pieces)
	return true
}

// Print the game board
func (g *Game) Print() {
	for _ = range g.Tab[0] {
		fmt.Print("+---")
	}
	fmt.Println("+")
	for _, row := range g.Tab {
		for _, p := range row {
			fmt.Printf("| %c ", p)
		}
		fmt.Println("|")
		for _ = range row {
			fmt.Print("+---")
		}
		fmt.Println("+")
	}
}

// Check who won
func (g *Game) Check() rune {
	win := make([]bool, len(g.Pieces))
	for k := range g.Pieces {
		win[k] = true
	}
	for i := range g.Tab {
		for j := range g.Tab[i] {
			for k, p := range g.Pieces {
				win[k] = win[k] && g.Tab[i][j] == p
			}
		}
	}
	for k, p := range g.Pieces {
		if win[k] {
			return p
		}
	}
	for k := range g.Pieces {
		win[k] = true
	}
	for i := range g.Tab {
		for j := range g.Tab[i] {
			for k, p := range g.Pieces {
				win[k] = win[k] && g.Tab[j][i] == p
			}
		}
	}
	for k, p := range g.Pieces {
		if win[k] {
			return p
		}
	}
	for k := range g.Pieces {
		win[k] = true
	}
	for i := range g.Tab {
		for k, p := range g.Pieces {
			win[k] = win[k] && g.Tab[i][i] == p
		}
	}
	for k, p := range g.Pieces {
		if win[k] {
			return p
		}
	}
	for k := range g.Pieces {
		win[k] = true
	}
	for i := range g.Tab {
		for k, p := range g.Pieces {
			win[k] = win[k] && g.Tab[len(g.Tab)-1-i][i] == p
		}
	}
	for k, p := range g.Pieces {
		if win[k] {
			return p
		}
	}
	return ' '
}
