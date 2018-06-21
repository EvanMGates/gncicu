package main

import (
	"math/rand"
	"time"
)

var fortniteLandingZones = [...]string{
	"Junk Junction",
	"Haunted Hills",
	"Anarchy Acres",
	"Risky Reels",
	"Wailing Woods",
	"Tomato Town",
	"Lonely Lodge",
	"Pleasant Park",
	"Loot Lake",
	"Snobby Shores",
	"Tilted Towers",
	"Dusty Divot",
	"Retail Row",
	"Greasy Grove",
	"Shifty Shafts",
	"Salty Springs",
	"Fatal Fields",
	"Moisty Mire",
	"Flush Factory",
	"Lucky Landing",
}

func randomLandingZone() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := int(r.Int31()) % len(fortniteLandingZones)
	return fortniteLandingZones[i]
}
