package player

import (
	"github.com/swexbe/govulcantv/internal/process"
	"sync"
	"time"
)

var current *process.Video
var channel chan bool
var lock sync.Mutex

func Start() {
	current = process.GetNext()
	channel = make(chan bool)
	go waitForNext()
}

func GetCurrent() *process.Video {
	return current
}

func ForcePlayVideo(video *process.Video) {
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
		current = process.GetNext()
		lock.Unlock()
	}
}