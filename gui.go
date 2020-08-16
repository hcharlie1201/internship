package main

import ("github.com/sciter-sdk/go-sciter"
"github.com/sciter-sdk/go-sciter/window"
)

func GUI() {
    windowRect := sciter.NewRect(400, 400, 400, 400)
    win, _ := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN, windowRect)
    win.SetTitle("Load main UI\n")
    win.LoadFile("./html/index.html")
    win.Show()
    win.Run()
}
