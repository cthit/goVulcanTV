package process

import (
	"github.com/swexbe/govulcantv/internal/db/queries"
	"math/rand"
)

type NextVideo struct {
	Id string `json:"id"`
	LengthSeconds uint32 `json:"playLengthSeconds"`
}

func GetNext() *NextVideo {
	enabled := queries.GetEnabled()
	index := rand.Intn(len(enabled))
	chosen := enabled[index]
	return &NextVideo{
		Id: chosen.YoutubeID,
		LengthSeconds: chosen.LengthSeconds,
	}
}
