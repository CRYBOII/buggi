package buggi

import (
	"testing"
	"time"
)

func TestDetect(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				HeartbeatCheckingTest()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	<-quit

}
