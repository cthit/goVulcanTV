package process

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/player"
)

func GetCurrent() *common.Video {
	return player.GetCurrent()
}