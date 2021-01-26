package process

import (
	"github.com/swexbe/govulcantv/internal/player"
	"time"
)

type CurrentVideoResponse struct {
	*player.CurrentVideo
	SecondsRemaining int64
}

func GetCurrent() *CurrentVideoResponse {
	current := player.GetCurrent()
	if current != nil {
		diff := current.StartedAt + int64(current.Video.LengthSeconds) - time.Now().Unix()
		return &CurrentVideoResponse{
			CurrentVideo:     current,
			SecondsRemaining: diff,
		}
	}
	return nil
}