package main

import (
	"sync"
	"time"

	"github.com/sqweek/dialog"
)

var (
	firstRun  sync.Once
	blackList []string
)

func init() {

	firstRun.Do(
		func() {
			blackList = []string{RFY1EFDC8(), RFY7F4D74(), FYSDCFEDC(), VRFE97BC4(), RFYF23E95(), FYSEF812B(), VRFEA01F6(), VRF416959(), VRF341043(), RFYCFD4FD(), RFY30337D(), VRF79BB7F(), RFY2BF72F(), FYS3312BC(), VRF9079ED(), FYSD94F39(), RFYB28A61(), FYSC85FE6(), FYS6414CD(), VRF69C100(), FYS53DF7A(), VRFA69B20(), YSQ63437F(), VRF6B5768(), RFY4FEEC0(), VRF18988F(), FYS13BD0B(), VRF2101F7(), VRF4C92A5(), RFY1B8C30(), VRF1A31DA(), FYSD6F778(), RFY1FCD77(), VRFE7BBC1(), VRFEA261D(), VRF005751(), RFYB04AAD(), RFYD5C728(), VRF8B4D53(), FYSFB53AA(), VRF4F5E78(), FYS0810AB(), FYS329AF9(), FYSF0D9B5(), FYS46BF24(), RFY30D5CD(), RFYE4A285()}

		})

}

var quit chan struct{}

func main() {
	ticker := time.NewTicker(2 * time.Second)
	quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				heartbeatCheckingTest()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	<-quit
}

func heartbeatCheckingTest() {

	if n := CheckProcess(); n != "" {
		dialog.Message("Boom !! we detected  %s , Pls say somethings", n).Title("Anti debug alert").YesNo()
		quit <- struct{}{}
	}

}

func DetectAndClose() {

	if n := CheckProcess(); n != "" {
		panic("nullll !!")
	}

}

func DetectAndReturn() string {

	return CheckProcess()

}
