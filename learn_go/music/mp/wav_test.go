package mp

import "testing"

func TestWAV(t *testing.T) {
	var p Player
	p = &WAVPlayer{}
	p.Play("music.163.com")

}
