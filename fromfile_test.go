package main

import (
	"testing"
)

func TestFromFile(t *testing.T) {
	expected1 := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	actual1 := FromFile("./fromfile_test1.csv")
	for y := range expected1 {
		for x := range expected1[0] {
			if actual1[y][x] != expected1[y][x] {
				t.Error("全て死んだセルの場合が通りません")
			}
		}
	}

	expected2 := [][]bool{
		[]bool{true, true, true},
		[]bool{true, true, true},
		[]bool{true, true, true},
	}
	actual2 := FromFile("./fromfile_test2.csv")
	for y := range expected2 {
		for x := range expected2[0] {
			if actual2[y][x] != expected2[y][x] {
				t.Error("全て生きたセルの場合が通りません")
			}
		}
	}

	expected3 := [][]bool{
		[]bool{true, false, false},
		[]bool{false, true, true},
		[]bool{false, true, false},
	}
	actual3 := FromFile("./fromfile_test3.csv")
	for y := range expected3 {
		for x := range expected3[0] {
			if actual3[y][x] != expected3[y][x] {
				t.Error("平均的な場合が通りません")
			}
		}
	}
}
