package engine

import (
	"embed"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/*
var fonts embed.FS

// //go:embed images/*
// var images embed.FS

var faces = map[string]*font.Face{}

func LoadFont(name string) *font.Face {
	bytes, err := fonts.ReadFile("fonts/" + name + ".ttf")

	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(bytes)

	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, nil)
	faces[name] = &face
	return &face
}

func GetFont(name string) *font.Face {
	return faces[name]
}
