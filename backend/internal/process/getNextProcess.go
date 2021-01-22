package process

import (
	"github.com/swexbe/govulcantv/internal/db/queries"
	"math/rand"
)

func GetNext() string {
	enabled := queries.GetEnabled()
	index := rand.Intn(len(enabled))
	chosen := enabled[index]
	return chosen.YoutubeID
}
