package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func getGNC() *gnc {
	return &gnc{
		config: &Config{
			Server: Server{
				Assets: ".",
			},
		},
	}
}

func TestRandomLandingZone(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go randomLandingZone()
	}
}

func TestBuildImage(t *testing.T) {
	outFile, err := os.Create("out.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	gnc := getGNC()
	gnc.writeLandingZoneImage(randomLandingZone(), b)
}
