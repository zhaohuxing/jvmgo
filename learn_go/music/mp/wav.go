package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) //假装正在播放
		fmt.Print(". ")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
