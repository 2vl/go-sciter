package main

import (
	sciter "github.com/wj008/go-sciter"
	res "github.com/wj008/go-sciter/res"
	window "github.com/wj008/go-sciter/window"
	"log"
)

/*
A go.rice usage example.

To build it you need an actual `go.rice` package and its tool:

* install go.gice: `go get github.com/GeertJohan/go.rice && go get github.com/GeertJohan/go.rice/rice`
* append resources to executable via `go build && rice append /exec restest.exe` (omit .exe if not on Windows)
* or you can embed it via `rice embed-go && go build` (in that order)

Now the resulting executable is completely stand-alone.
*/
func main() {
	// As usual, create a sciter window.
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Handle resources via sciter.rice loader.
	// It handles URLs like `file://box/` and `rice://box/`.
	res.HandleDataLoad(w.Sciter)

	// 3. Load a packaged resource.
	err = w.LoadFile("res://html/simple.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Show()
	w.Run()
}
