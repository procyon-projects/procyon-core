package core

import (
	"crypto/rand"
	"io"
	"net"
	"sync"
	"time"
)

/* Customized google.uuid
 * as defers is removed, it makes it faster than google.uuid
 */

const (
	lillian    = 2299160          // Julian day of 15 Oct 1582
	unix       = 2440587          // Julian day of 1 Jan 1970
	epoch      = unix - lillian   // Days between epochs
	g1582      = epoch * 86400    // seconds between epochs
	g1582ns100 = g1582 * 10000000 // 100s of a nanoseconds between epochs
)

var (
	timeMu   sync.Mutex
	lastTime uint64
	clockSeq uint16

	nodeID [6]byte
	rander = rand.Reader
)

func init() {
	iname, addr := getHardwareInterface()

	if iname != "" && addr != nil {
		copy(nodeID[:], addr)
	} else {
		randomBits(nodeID[:])
	}

	if clockSeq == 0 {
		var b [2]byte
		randomBits(b[:])

		seq := int(b[0])<<8 | int(b[1])
		oldSeq := clockSeq
		clockSeq = uint16(seq&0x3fff) | 0x8000

		if oldSeq != clockSeq {
			lastTime = 0
		}

	}
}

const hexTable = "0123456789abcdef"

func GenerateUUID(uuidBuffer []byte) {
	timeMu.Lock()
	now := uint64(time.Now().UnixNano()/100) + g1582ns100

	if now <= lastTime {
		clockSeq = ((clockSeq + 1) & 0x3fff) | 0x8000

	}

	lastTime = now
	timeMu.Unlock()

	timeLow := uint32(now & 0xffffffff)
	timeMid := uint16((now >> 32) & 0xffff)
	timeHi := uint16((now >> 48) & 0x0fff)
	timeHi |= 0x1000

	b := byte(timeLow >> 24)
	uuidBuffer[0] = hexTable[b>>4]
	uuidBuffer[1] = hexTable[b&0x0f]

	b = byte(timeLow >> 16)
	uuidBuffer[2] = hexTable[b>>4]
	uuidBuffer[3] = hexTable[b&0x0f]

	b = byte(timeLow >> 8)
	uuidBuffer[4] = hexTable[b>>4]
	uuidBuffer[5] = hexTable[b&0x0f]

	b = byte(timeLow)
	uuidBuffer[6] = hexTable[b>>4]
	uuidBuffer[7] = hexTable[b&0x0f]

	uuidBuffer[8] = '-'

	b = byte(timeMid >> 8)
	uuidBuffer[9] = hexTable[b>>4]
	uuidBuffer[10] = hexTable[b&0x0f]

	b = byte(timeMid)
	uuidBuffer[11] = hexTable[b>>4]
	uuidBuffer[12] = hexTable[b&0x0f]

	uuidBuffer[13] = '-'
	b = byte(timeHi >> 8)
	uuidBuffer[14] = hexTable[b>>4]
	uuidBuffer[15] = hexTable[b&0x0f]

	b = byte(timeHi)
	uuidBuffer[16] = hexTable[b>>4]
	uuidBuffer[17] = hexTable[b&0x0f]

	uuidBuffer[18] = '-'
	b = byte(clockSeq >> 8)
	uuidBuffer[19] = hexTable[b>>4]
	uuidBuffer[20] = hexTable[b&0x0f]

	b = byte(clockSeq)
	uuidBuffer[21] = hexTable[b>>4]
	uuidBuffer[22] = hexTable[b&0x0f]

	uuidBuffer[23] = '-'
	b = nodeID[0]
	uuidBuffer[24] = hexTable[b>>4]
	uuidBuffer[25] = hexTable[b&0x0f]

	b = nodeID[1]
	uuidBuffer[26] = hexTable[b>>4]
	uuidBuffer[27] = hexTable[b&0x0f]

	b = nodeID[2]
	uuidBuffer[28] = hexTable[b>>4]
	uuidBuffer[29] = hexTable[b&0x0f]

	b = nodeID[3]
	uuidBuffer[30] = hexTable[b>>4]
	uuidBuffer[31] = hexTable[b&0x0f]

	b = nodeID[4]
	uuidBuffer[32] = hexTable[b>>4]
	uuidBuffer[33] = hexTable[b&0x0f]

	b = nodeID[5]
	uuidBuffer[34] = hexTable[b>>4]
	uuidBuffer[35] = hexTable[b&0x0f]
}

func getHardwareInterface() (string, []byte) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return "", nil
	}

	for _, ifs := range interfaces {

		if len(ifs.HardwareAddr) >= 6 {
			return ifs.Name, ifs.HardwareAddr
		}

	}
	return "", nil
}

func randomBits(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err.Error())
	}
}
