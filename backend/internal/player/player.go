package player

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"math/rand"
	"sync"
	"time"
)

var current *common.Video
var channel chan bool
var lock sync.Mutex

func Start() {
	current = GetNext()
	channel = make(chan bool)
	go waitForNext()
}

func GetCurrent() *common.Video {
	return current
}

func ForcePlayVideo(video *common.Video) {
	channel<-true
	lock.Lock()
	current = video
	lock.Unlock()
	go waitForNext()
}

func waitForNext() {
	for {
		time.Sleep(time.Duration(current.LengthSeconds) * time.Second)
		if v := <- channel; v {
			return
		}
		lock.Lock()
		current = GetNext()
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