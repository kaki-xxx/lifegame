package main

// CountAlive returns next cell status
func CountAlive(table [][]bool, x, y int) int {
	h, w := len(table), len(table[0])
	count := 0
	for yi := y - 1; yi <= y+1; yi++ {
		if !(yi >= 0 && yi <= h-1) {
			continue
		}
		for xi := x - 1; xi <= x+1; xi++ {
			if !(xi >= 0 && xi <= w-1) {
				continue
			}
			if xi == x && yi == y {
				continue
			}
			if table[yi][xi] {
				count++
			}
		}
	}
	return count
}
