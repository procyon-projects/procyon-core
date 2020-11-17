package core

import (
	"crypto/rand"
	"encoding/binary"
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
}

func GenerateUUID(uuid []byte) {
	timeMu.Lock()
	now := uint64(time.Now().UnixNano()/100) + g1582ns100

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

	if now <= lastTime {
		clockSeq = ((clockSeq + 1) & 0x3fff) | 0x8000
	}
	lastTime = now
	timeMu.Unlock()

	timeLow := uint32(now & 0xffffffff)
	timeMid := uint16((now >> 32) & 0xffff)
	timeHi := uint16((now >> 48) & 0x0fff)
	timeHi |= 0x1000

	binary.BigEndian.PutUint32(uuid[0:], timeLow)
	binary.BigEndian.PutUint16(uuid[4:], timeMid)
	binary.BigEndian.PutUint16(uuid[6:], timeHi)
	binary.BigEndian.PutUint16(uuid[8:], clockSeq)
	copy(uuid[10:], nodeID[:])
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
