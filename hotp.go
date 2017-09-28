package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"math"
)

const (
	fourBitMask      = 0xf
	eightBitMask     = 0xff
	thirtyOneBitMask = 0x7fffffff
)

// Generator generates HMAC-based One-Time Password
type Generator struct {
	Counter uint64 // C in RFC4226
	Secret  string // K in RFC4226
	Digit   int
}

// Generate HOTP
func (g *Generator) Generate() int64 {
	hs := hmacSHA1([]byte(g.Secret), counterToBytes(g.Counter))
	snum := truncate(hs)
	d := int64(snum) % int64(math.Pow10(g.Digit))
	g.Counter++
	return d
}

func counterToBytes(c uint64) []byte {
	t := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		t[i] = byte(c & eightBitMask)
		c = c >> 8
	}
	return t
}

func hmacSHA1(k, c []byte) (hs []byte) {
	mac := hmac.New(sha1.New, k)
	mac.Write(c)
	hs = mac.Sum(nil)
	return hs
}

func truncate(hs []byte) int {
	offsetBits := hs[len(hs)-1] & fourBitMask
	offset := int(offsetBits)
	p := hs[offset : offset+4]
	return int(binary.BigEndian.Uint32(p)) & thirtyOneBitMask
}
