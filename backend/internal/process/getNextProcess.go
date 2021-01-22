package process

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"math/rand"
)

func GetNext() *common.Video {
	enabled := queries.GetEnabled()
	index := rand.Intn(len(enabled))
	chosen := enabled[index]
	return &common.Video{
		Id: chosen.YoutubeID,
		LengthSeconds: chosen.LengthSeconds,
	}
}
