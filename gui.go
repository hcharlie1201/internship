package main

import (
	"os"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func GUI() {
	windowRect := sciter.NewRect(200, 200, 900, 600)
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_TITLEBAR, windowRect)
	win.SetTitle("Load main UI")
    dir, _ := os.Getwd()
    dir += "/html/index.html"
	win.LoadFile(dir)
	win.Show()
	win.Run()
}
