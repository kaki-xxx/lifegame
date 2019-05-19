package main

import (
	"testing"
)

func TestCountAlive(t *testing.T) {
	table1 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, false, false},
		[]bool{false, false, false},
	}
	if CountAlive(table1, 1, 1) != 3 {
		t.Error("誕生の場合が通りません")
	}
	table2 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	if CountAlive(table2, 1, 1) != 3 {
		t.Error("生存の場合が通りません")
	}
	table3 := [][]bool{
		[]bool{false, false, false},
		[]bool{false, true, true},
		[]bool{false, false, false},
	}
	if CountAlive(table3, 1, 1) != 1 {
		t.Error("過疎の場合が通りません")
	}
	table4 := [][]bool{
		[]bool{true, true, true},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	if CountAlive(table4, 1, 1) != 4 {
		t.Error("過密の場合が通りません")
	}

	if CountAlive(table4, 0, 0) != 3 {
		t.Error("角の場合が通りません")
	}
}
