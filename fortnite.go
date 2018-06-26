package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
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

func writeLandingZoneImage(landingZone string, w io.Writer) {
	fontFile := "assets/Burbank_Big_Condensed_Black.ttf"
	imageFile := "assets/dropNoText.png"
	dpi := float64(72)
	size := float64(142)
	spacing := float64(1.5)

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	imageStream, err := os.Open(imageFile)
	if err != nil {
		log.Println(err)
		return
	}
	dropImage, _, _ := image.Decode(imageStream)
	rgba := dropImage.(draw.Image)
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(color.RGBA{0xD4, 0xB8, 0x20, 0xff}))
	c.SetHinting(font.HintingNone)

	// Draw the text.
	pt := freetype.Pt(264, 520+int(c.PointToFixed(size)>>6))
	_, err = c.DrawString(landingZone, pt)
	if err != nil {
		log.Println(err)
		return
	}
	pt.Y += c.PointToFixed(size * spacing)
	png.Encode(w, rgba)
	// Save that RGBA image to disk.
	/*
		outFile, err := os.Create("out.png")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer outFile.Close()
		b := bufio.NewWriter(outFile)
		err = png.Encode(b, rgba)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		err = b.Flush()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Println("Wrote out.png OK.")
		return
	*/
}
