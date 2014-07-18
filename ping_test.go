package main

import "testing"
import "encoding/hex"
import "time"

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

func TestComputeStatistics(t *testing.T) {
	oneSec := 1 * time.Second
	pingTimes := []time.Duration{oneSec}
	_, min, max, _ := ComputeStats(pingTimes)
	if min != oneSec {
		t.Errorf("%v != %v", min, oneSec)
	}

	if max != oneSec {
		t.Errorf("%v != %v", max, oneSec)
	}
}

func TestComputeStatisticsRange(t *testing.T) {
	twenty := 30 * time.Second
	forty := 40 * time.Second
	pingTimes := []time.Duration{twenty, forty}
	total, min, max, avg := ComputeStats(pingTimes)

	if min != twenty {
		t.Errorf("%v = %v", min, twenty)
	}

	if max != forty {
		t.Errorf("%v = %v", min, forty)
	}

	if total != twenty+forty {
		t.Errorf("Total is not 60 seconds: %v", total)
	}

	if avg != (min+max)/2 {
		t.Errorf("Average is not 30 seconds: %v", avg)
	}
}
