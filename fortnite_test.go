package main

import (
	"testing"
)

func TestRandomLandingZone(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go randomLandingZone()
	}
}
