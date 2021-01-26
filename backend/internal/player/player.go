package player

import (
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"log"
	"math/rand"
	"sync"
	"time"
)

type CurrentVideo struct {
	Video *common.Video
	StartedAt int64
}

var current CurrentVideo
var currentLock sync.Mutex

var override *CurrentVideo
var overrideLock sync.Mutex

func Start() {
	go waitForNext()
}

func GetCurrent() *CurrentVideo {
	return &current
}

func ForcePlayVideo(video *common.Video) {
	overrideLock.Lock()
	override = &CurrentVideo{
		Video: video,
		StartedAt: time.Now().Unix(),
	}
	overrideLock.Unlock()
}

func waitForNext() {
	for {
		if override != nil && override.Video != nil {
			log.Printf("Override video available, overriding with %s\n", override.Video.Id)
			// Override the current video with the override video
			overrideLock.Lock()
			current = *override
			current.StartedAt = time.Now().Unix()
			override = nil
			overrideLock.Unlock()
		} else if current.Video == nil || time.Now().Unix() >= current.StartedAt + int64(current.Video.LengthSeconds) {
			// Video is done or missing, play a new one
			next := GetNext()
			if next != nil {

				currentLock.Lock()
				current = CurrentVideo{
					Video:     next,
					StartedAt: time.Now().Unix(),
				}
				log.Printf("Selected video ID %s\n", current.Video.Id)
				currentLock.Unlock()
			} else {
				log.Println("No enabled videos found")
			}
		}

		time.Sleep(time.Second)
	}
}

func GetNext() *common.Video {
	enabled := queries.GetEnabled()
	if len(enabled) == 0 {
		return nil
	}

	index := rand.Intn(len(enabled))
	chosen := enabled[index]
	return &common.Video{
		Id: chosen.YoutubeID,
		LengthSeconds: chosen.LengthSeconds,
	}
}