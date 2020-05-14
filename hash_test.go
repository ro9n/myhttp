package main

import "testing"

func TestChecksum(t *testing.T) {
	const (
		empty  = "d41d8cd98f00b204e9800998ecf8427e"
		adjust = "3b7770f7743e8f01f0fd807f304a21d0"
	)

	if result, _ := checksum(""); result != empty {
		t.Errorf("checksum(\"\") failed, expected %v, got %v", empty, result)
	}

	if result, _ := checksum("adjust"); result != adjust {
		t.Errorf("checksum(\"adjust\") failed, expected %v, got %v", adjust, result)
	}
}
