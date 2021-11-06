package buggi

import (
	"sync"
	"time"

	"github.com/CRYBOII/buggi/obfuscated"
	"github.com/sqweek/dialog"
)

var (
	firstRun  sync.Once
	blackList []string
)

func init() {

	firstRun.Do(
		func() {
			blackList = []string{obfuscated.RFY1EFDC8(), obfuscated.RFY7F4D74(), obfuscated.FYSDCFEDC(), obfuscated.VRFE97BC4(), obfuscated.RFYF23E95(), obfuscated.FYSEF812B(), obfuscated.VRFEA01F6(), obfuscated.VRF416959(), obfuscated.VRF341043(), obfuscated.RFYCFD4FD(), obfuscated.RFY30337D(), obfuscated.VRF79BB7F(), obfuscated.RFY2BF72F(), obfuscated.FYS3312BC(), obfuscated.VRF9079ED(), obfuscated.FYSD94F39(), obfuscated.RFYB28A61(), obfuscated.FYSC85FE6(), obfuscated.FYS6414CD(), obfuscated.VRF69C100(), obfuscated.FYS53DF7A(), obfuscated.VRFA69B20(), obfuscated.YSQ63437F(), obfuscated.VRF6B5768(), obfuscated.RFY4FEEC0(), obfuscated.VRF18988F(), obfuscated.FYS13BD0B(), obfuscated.VRF2101F7(), obfuscated.VRF4C92A5(), obfuscated.RFY1B8C30(), obfuscated.VRF1A31DA(), obfuscated.FYSD6F778(), obfuscated.RFY1FCD77(), obfuscated.VRFE7BBC1(), obfuscated.VRFEA261D(), obfuscated.VRF005751(), obfuscated.RFYB04AAD(), obfuscated.RFYD5C728(), obfuscated.VRF8B4D53(), obfuscated.FYSFB53AA(), obfuscated.VRF4F5E78(), obfuscated.FYS0810AB(), obfuscated.FYS329AF9(), obfuscated.FYSF0D9B5(), obfuscated.FYS46BF24(), obfuscated.RFY30D5CD(), obfuscated.RFYE4A285()}

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

func SimpleRun(delay int) {

	ticker := time.NewTicker(time.Duration(delay) * time.Second)
	quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				DetectAndClose()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
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
