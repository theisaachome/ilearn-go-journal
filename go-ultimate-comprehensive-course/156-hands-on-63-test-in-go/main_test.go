package main

import "testing"

// func TestX...()
func TestAdd(t *testing.T) {
	total := Add(1, 2)
	if total != 3 {
		t.Error("Expected 3, got ", total)
	}
}
