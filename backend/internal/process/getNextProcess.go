package process

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/player"
)

func GetNext() *common.Video {
	return player.GetNext()
}
