package main

import "testing"

func TestDataRaceCondition(t *testing.T) {
	var counter int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			counter += int32(i)
		}(i)
	}
}

//
// $ go test --race
// ==================
// WARNING: DATA RACE
// Read at 0x00c00001a25c by goroutine 8:
//
