package resx

import (
	"embed"
	"github.com/wj008/go-sciter"
	"log"
	"strings"
)

func HandleDataLoad(s *sciter.Sciter, f embed.FS) {
	header := &sciter.CallbackHandler{
		OnLoadData: func(ld *sciter.ScnLoadData) int {
			uri := ld.Uri()
			path := ""
			log.Println("loading:", uri)
			// file:// or rice://
			if strings.HasPrefix(uri, "resx://") {
				path = uri[7:]
			} else {
				// // do not handle schemes other than file:// or rice://
				return sciter.LOAD_OK
			}
			log.Println("resx loading:", path)
			dat, err := f.ReadFile(path)
			if err != nil {
				// box locating failed, return to Sciter loading
				return sciter.LOAD_OK
			} else {
				// using rice found data
				s.DataReady(uri, dat)
			}
			return sciter.LOAD_OK
		},
	}
	s.SetCallback(header)
}
