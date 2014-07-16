package main

import "testing"
import "encoding/hex"

func TestCheckSum(t *testing.T) {
	data := "45000047738840004006a2c4839f0e85839f0ea1"
	bin, _ := hex.DecodeString(data)
	actual := CheckSum(bin)
	var expected uint16 = 0
	if expected != actual {
		t.Errorf("%v = %v", expected, actual)
	}
}

func TestCheckSumWikipedia(t *testing.T) {
	data := "45000073000040004011b861c0a80001c0a800c7"
	bin, _ := hex.DecodeString(data)
	actual := CheckSum(bin)
	var expected uint16 = 0
	if expected != actual {
		t.Errorf("%v = %v", expected, actual)
	}
}
