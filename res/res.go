package res

import (
	"embed"
	"strings"
	"github.com/wj008/go-sciter"
)

//go:embed html/*
var f embed.FS

func OnLoadData(s *sciter.Sciter) func(ld *sciter.ScnLoadData) int {
	return func(ld *sciter.ScnLoadData) int {
		uri := ld.Uri()
		path := ""
		//log.Println("loading:", uri)
		// file:// or rice://
		if strings.HasPrefix(uri, "res://") {
			path = uri[7:]
		} else {
			// // do not handle schemes other than file:// or rice://
			return sciter.LOAD_OK
		}
		//log.Println("rice loading:", path)
		dat, err := f.ReadFile(path)
		if err != nil {
			// box locating failed, return to Sciter loading
			return sciter.LOAD_OK
		} else {
			// using rice found data
			s.DataReady(uri, dat)
		}
		return sciter.LOAD_OK
	}
}

func newHandler(s *sciter.Sciter) *sciter.CallbackHandler {
	return &sciter.CallbackHandler{
		OnLoadData: OnLoadData(s),
	}
}

func HandleDataLoad(s *sciter.Sciter) {
	s.SetCallback(newHandler(s))
}
