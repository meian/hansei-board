package main

type AppConfig struct {
	Port int `default:"3000"`
}

type PreviewConfig struct {
	Scale   float64 `default:"0.5"`
	Quality int     `default:"70"`
}
