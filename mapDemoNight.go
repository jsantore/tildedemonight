package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

const mapPath = "demo.tmx"

type MapDemo struct {
	Level    *tiled.Map
	tileDict map[uint32]*ebiten.Image
}

func (m MapDemo) Update() error {
	return nil
}

func main() {

}
