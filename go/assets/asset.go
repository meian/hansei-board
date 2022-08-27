package assets

import (
	"embed"
)

//go:embed font/*
var font embed.FS

// Font は指定したリソースのフォントバイナリを取得する
func Font(name string) ([]byte, error) {
	return font.ReadFile("font/" + name)
}

//go:embed img/*
var img embed.FS

// Img は指定したリソースの画像バイナリを取得する
func Img(name string) ([]byte, error) {
	return img.ReadFile("img/" + name)
}
