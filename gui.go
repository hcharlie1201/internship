package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
    "os"
)

func GUI() {
	windowRect := sciter.NewRect(200, 200, 900, 600)
	win, _ := window.New(sciter.SW_MAIN | sciter.SW_TITLEBAR, windowRect)
    rootDir, _ := os.Getwd()
	win.SetTitle("Load main UI\n")
    win.LoadFile("file:///" + rootDir + "/html/index.html")
	win.Show()
	win.Run()
}
