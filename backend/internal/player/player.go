package player

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"math/rand"
	"sync"
	"time"
)

type CurrentVideo struct {
	Video *common.Video
	StartedAt int64
}

var current CurrentVideo

var channel chan bool
var lock sync.Mutex

func Start() {
	current = CurrentVideo{
		Video: GetNext(),
		StartedAt: time.Now().Unix(),
	}
	channel = make(chan bool)
	go waitForNext()
}

func GetCurrent() *CurrentVideo {
	return &current
}

func ForcePlayVideo(video *common.Video) {
	channel<-true
	lock.Lock()
	current = CurrentVideo{
		Video: video,
		StartedAt: time.Now().Unix(),
	}
	lock.Unlock()
	go waitForNext()
}

func waitForNext() {
	for {
		// Update startedAt in case any time has gone by since current was last updated
		lock.Lock()
		current = CurrentVideo{
			Video:     current.Video,
			StartedAt: time.Now().Unix(),
		}
		lock.Unlock()
		time.Sleep(time.Duration(current.Video.LengthSeconds) * time.Second)
		if v := <- channel; v {
			return
		}
		lock.Lock()
		current = CurrentVideo{
			Video: GetNext(),
			StartedAt: time.Now().Unix(),
		}
		lock.Unlock()
	}
}

func GetNext() *common.Video {
	enabled := queries.GetEnabled()
	index := rand.Intn(len(enabled))
	chosen := enabled[index]
	return &common.Video{
		Id: chosen.YoutubeID,
		LengthSeconds: chosen.LengthSeconds,
	}
}