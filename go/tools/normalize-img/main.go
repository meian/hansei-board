package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const (
	rootDir = "./img"
	org     = "original.jpg"
	dst     = "normalized.jpg"
)

type param struct {
	dir     string
	quality int
}

func main() {
	p := parseParam(os.Args[1:])
	dir := filepath.Join(rootDir, p.dir)
	if fi, err := os.Stat(dir); err != nil {
		panic(err)
	} else if !fi.IsDir() {
		panic(fmt.Sprintf("not found directory: %s", dir))
	}
	op := filepath.Join(dir, org)
	ob, err := ioutil.ReadFile(op)
	if err != nil {
		panic(err)
	}
	img, err := imaging.Decode(bytes.NewReader(ob), imaging.AutoOrientation(true))
	if err != nil {
		panic(err)
	}
	dp := filepath.Join(dir, dst)
	df, err := os.Create(dp)
	if err != nil {
		panic(err)
	}
	jopts := &jpeg.Options{Quality: p.quality}
	if err := jpeg.Encode(df, img, jopts); err != nil {
		panic(err)
	}
	cmd := exec.Command("ls", "-l", dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Print(cmd.Run())
}

func parseParam(args []string) *param {
	fs := flag.NewFlagSet("normalize-img", flag.PanicOnError)
	var p param
	fs.IntVar(&p.quality, "q", 80, "JPEG quality")
	fs.Parse(args)
	p.dir = fs.Arg(0)
	return &p
}
