package mp

import "testing"

func TestMP3(t *testing.T) {
	var p Player
	p = &MP3Player{stat: 1}
	p.Play("www.xxx.com")
}
