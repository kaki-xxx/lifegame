package main

// UpdateTable returns next table status
func UpdateTable(table [][]bool) [][]bool {
	nexttable := make([][]bool, len(table))
	for i := range nexttable {
		nexttable[i] = make([]bool, len(table[0]))
	}
	for y := range nexttable {
		for x := range nexttable {
			if alive := CountAlive(table, x, y); alive == 3 {
				nexttable[y][x] = true
			} else if alive == 2 {
				nexttable[y][x] = table[y][x]
			} else {
				nexttable[y][x] = false
			}
		}
	}
	return nexttable
}
