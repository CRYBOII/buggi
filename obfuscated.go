package buggi

import (
	"unsafe"
)

func eax() uint8 {
	return uint8(unsafe.Sizeof(true))
}

//  [ RFY1EFDC8 ] ====>  BurpSuite

func RFY1EFDC8() string {
	return string(
		[]byte{
			(eax()<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
		},
	)
}

//  [ RFY7F4D74 ] ====>  BurpSuiteFree

func RFY7F4D74() string {
	return string(
		[]byte{
			(eax()<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
		},
	)
}

//  [ FYSDCFEDC ] ====>  Charles

func FYSDCFEDC() string {
	return string(
		[]byte{
			((eax()<<eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ VRFE97BC4 ] ====>  dumpcap

func VRFE97BC4() string {
	return string(
		[]byte{
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
		},
	)
}

//  [ RFYF23E95 ] ====>  Fiddler

func RFYF23E95() string {
	return string(
		[]byte{
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}

//  [ FYSEF812B ] ====>  httpsMon

func FYSEF812B() string {
	return string(
		[]byte{
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
		},
	)
}

//  [ VRFEA01F6 ] ====>  httpwatchstudiox64

func VRFEA01F6() string {
	return string(
		[]byte{
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ VRF416959 ] ====>  mitmdump

func VRF416959() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
		},
	)
}

//  [ VRF341043 ] ====>  mitmweb

func VRF341043() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
		},
	)
}

//  [ RFYCFD4FD ] ====>  NetworkMiner

func RFYCFD4FD() string {
	return string(
		[]byte{
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}

//  [ RFY30337D ] ====>  Proxifier

func RFY30337D() string {
	return string(
		[]byte{
			(eax()<<eax()<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}

//  [ VRF79BB7F ] ====>  rpcapd

func VRF79BB7F() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ RFY2BF72F ] ====>  smsniff

func RFY2BF72F() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
		},
	)
}

//  [ FYS3312BC ] ====>  tshark

func FYS3312BC() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ VRF9079ED ] ====>  WinDump

func VRF9079ED() string {
	return string(
		[]byte{
			((((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
		},
	)
}

//  [ FYSD94F39 ] ====>  Wireshark

func FYSD94F39() string {
	return string(
		[]byte{
			((((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ RFYB28A61 ] ====>  WSockExpert

func RFYB28A61() string {
	return string(
		[]byte{
			((((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
		},
	)
}

//  [ FYSC85FE6 ] ====>  x96dbg

func FYSC85FE6() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ FYS6414CD ] ====>  ollydbg

func FYS6414CD() string {
	return string(
		[]byte{
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ VRF69C100 ] ====>  ida64

func VRF69C100() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ FYS53DF7A ] ====>  idag

func FYS53DF7A() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ VRFA69B20 ] ====>  idag64

func VRFA69B20() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ YSQ63437F ] ====>  idaw

func YSQ63437F() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ VRF6B5768 ] ====>  idaw64

func VRF6B5768() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ RFY4FEEC0 ] ====>  idaq

func RFY4FEEC0() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()<<eax() | eax()),
		},
	)
}

//  [ VRF18988F ] ====>  idaq64

func VRF18988F() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ FYS13BD0B ] ====>  idau

func FYS13BD0B() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
		},
	)
}

//  [ VRF2101F7 ] ====>  idau64

func VRF2101F7() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ VRF4C92A5 ] ====>  scylla_x64

func VRF4C92A5() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((((eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ RFY1B8C30 ] ====>  scylla_x86

func RFY1B8C30() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((((eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
		},
	)
}

//  [ VRF1A31DA ] ====>  protection_id

func VRF1A31DA() string {
	return string(
		[]byte{
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(((((eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
		},
	)
}

//  [ FYSD6F778 ] ====>  windbg

func FYSD6F778() string {
	return string(
		[]byte{
			(((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ RFY1FCD77 ] ====>  reshacker

func RFY1FCD77() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}

//  [ VRFE7BBC1 ] ====>  ImportREC

func VRFE7BBC1() string {
	return string(
		[]byte{
			((eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			((eax()<<eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()),
		},
	)
}

//  [ VRFEA261D ] ====>  IMMUNITYDEBUGGER

func VRFEA261D() string {
	return string(
		[]byte{
			((eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			(eax()<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()),
			(((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()),
			((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
		},
	)
}

//  [ VRF005751 ] ====>  HTTPDebuggerUI

func VRF005751() string {
	return string(
		[]byte{
			(eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			(eax()<<eax()<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
		},
	)
}

//  [ RFYB04AAD ] ====>  HTTPDebuggerSvc

func RFYB04AAD() string {
	return string(
		[]byte{
			(eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax() << eax(),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			(eax()<<eax()<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
		},
	)
}

//  [ RFYD5C728 ] ====>  OLLYDBG

func RFYD5C728() string {
	return string(
		[]byte{
			((((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()<<eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(eax()<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			(((eax()<<eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()),
		},
	)
}

//  [ VRF8B4D53 ] ====>  ida

func VRF8B4D53() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
		},
	)
}

//  [ FYSFB53AA ] ====>  disassembly

func FYSFB53AA() string {
	return string(
		[]byte{
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
		},
	)
}

//  [ VRF4F5E78 ] ====>  scylla

func VRF4F5E78() string {
	return string(
		[]byte{
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
		},
	)
}

//  [ FYS0810AB ] ====>  Immunity

func FYS0810AB() string {
	return string(
		[]byte{
			((eax()<<eax()<<eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()|eax())<<eax()<<eax()<<eax() | eax()),
		},
	)
}

//  [ FYS329AF9 ] ====>  WinDbg

func FYS329AF9() string {
	return string(
		[]byte{
			((((eax()<<eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ FYSF0D9B5 ] ====>  x32dbg

func FYSF0D9B5() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ FYS46BF24 ] ====>  x64dbg

func FYS46BF24() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax() | eax()) << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax() | eax()) << eax(),
			((eax()<<eax()^eax())<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
		},
	)
}

//  [ RFY30D5CD ] ====>  reconstructor

func RFY30D5CD() string {
	return string(
		[]byte{
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()^eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()) << eax() << eax(),
			(((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()|eax())<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}

//  [ RFYE4A285 ] ====>  MegaDumper

func RFYE4A285() string {
	return string(
		[]byte{
			(((eax()<<eax()<<eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()|eax())<<eax() | eax()),
			((eax()<<eax()^eax())<<eax()<<eax()<<eax()<<eax()<<eax() ^ eax()),
			(eax()<<eax()<<eax()<<eax()<<eax() ^ eax()) << eax() << eax(),
			((((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()|eax())<<eax()<<eax() | eax()),
			((((eax()<<eax()^eax())<<eax()<<eax()^eax())<<eax()|eax())<<eax()<<eax() | eax()),
			((eax()<<eax()^eax())<<eax() ^ eax()) << eax() << eax() << eax() << eax(),
			(((eax()<<eax()^eax())<<eax()<<eax()<<eax()^eax())<<eax()<<eax() | eax()),
			(((eax()<<eax()^eax())<<eax()^eax())<<eax()<<eax()<<eax() | eax()) << eax(),
		},
	)
}
