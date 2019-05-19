package main

import (
	"testing"
)

func TestUpdateTable(t *testing.T) {
	test1 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, false, false},
		[]bool{false, false, false},
	}
	expected1 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	actual1 := UpdateTable(test1)
	for y := range expected1 {
		for x := range expected1[0] {
			if actual1[y][x] != expected1[y][x] {
				t.Error("誕生の場合が通りません")
			}
		}
	}

	test2 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	expected2 := [][]bool{
		[]bool{true, true, false},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	actual2 := UpdateTable(test2)
	for y := range expected2 {
		for x := range expected2[0] {
			if actual2[y][x] != expected2[y][x] {
				t.Error("生存の場合が通りません")
			}
		}
	}

	test3 := [][]bool{
		[]bool{false, false, false},
		[]bool{false, true, true},
		[]bool{false, false, false},
	}
	expected3 := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	actual3 := UpdateTable(test3)
	for y := range expected3 {
		for x := range expected3[0] {
			if actual3[y][x] != expected3[y][x] {
				t.Error("過疎の場合が通りません")
			}
		}
	}

	test4 := [][]bool{
		[]bool{true, true, true},
		[]bool{true, true, false},
		[]bool{false, false, false},
	}
	expected4 := [][]bool{
		[]bool{true, false, true},
		[]bool{true, false, true},
		[]bool{false, false, false},
	}
	actual4 := UpdateTable(test4)
	for y := range expected4 {
		for x := range expected4[0] {
			if actual4[y][x] != expected4[y][x] {
				t.Error("過密の場合が通りません")
			}
		}
	}
}
