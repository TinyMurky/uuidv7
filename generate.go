// Package uuidv7 implement [rfc9562](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7)
// 參考這邊的 code https://github.com/Hypersequent/uuid7
package uuidv7

import (
	"encoding/binary"
	"math"
	"math/rand/v2"
	"time"
)

// UUIDv7 Format as below
//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           unix_ts_ms                          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |          unix_ts_ms           |  ver  |       rand_a          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |var|                        rand_b                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                            rand_b                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// unix_ts_ms:
//     48-bit big-endian unsigned number of the Unix Epoch timestamp in milliseconds as per Section 6.1. Occupies bits 0 through 47 (octets 0-5).
// ver:
//     The 4-bit version field as defined by Section 4.2, set to 0b0111 (7). Occupies bits 48 through 51 of octet 6.
// rand_a:
//     12 bits of pseudorandom data to provide uniqueness as per Section 6.9 and/or optional constructs to guarantee additional monotonicity as per Section 6.2. Occupies bits 52 through 63 (octets 6-7).
// var:
//     The 2-bit variant field as defined by Section 4.1, set to 0b10. Occupies bits 64 and 65 of octet 8.
// rand_b:
//     The final 62 bits of pseudorandom data to provide uniqueness as per Section 6.9 and/or an optional counter to guarantee additional monotonicity as per Section 6.2. Occupies bits 66 through 127 (octets 8-15).
//
// Note: rand_a in this package is the number of nanoseconds that project to 12bit in the current millisecond

const (
	// versionInUint64 is uuid ver field, 4-bit (48~51)
	versionInUint64 = uint64(7)

	// variantInUint64 is variant field, 2 bit (64~65)
	variantInUint64 = uint64(2)
)

// millisecondInUint64 generate unix_ts_ms (0~47 bit)
func millisecondInUint64(uuidTime time.Time) uint64 {
	milliseconds := uint64(uuidTime.UnixNano() / int64(time.Millisecond))
	return milliseconds
}

// nanosecondAsRandAInUint64 generate rand_a (12 bit) 52~63 bit
// rand_a is the number of nanoseconds that project to 12bit in the current millisecond
func nanosecondAsRandAInUint64(uuidTime time.Time) uint64 {
	nanoseconds := uint64(uuidTime.UnixNano() % int64(time.Millisecond))

	portion := float64(nanoseconds) / float64(time.Millisecond)

	nanosecondsProjectToRandA := uint64(math.Floor(portion * 4096))

	return nanosecondsProjectToRandA
}

// randBInUint64 generate 62 bits random number (66~127)
func randBInUint64() (uint64, error) {
	// Below use crypto/rand, safe but slow

	// lim := new(big.Int).Lsh(big.NewInt(1), 62) // 2^62
	// n, err := rand.Int(rand.Reader, lim)       // [0, 2^62) ()
	// if err != nil {
	// 	return 0, err
	// }
	// rand.IntN()

	// Below use rand/v2, faster but unsafe
	const TwoPow62 = uint64(1) << 62

	// half-open interval [0,n), n is not included
	return rand.Uint64N(TwoPow62), nil
}

// first64Bits generate 64 bits of uuidv7, unix_ts_ms (48 bit) + ver (4 bit) + rand_a (12bit)
func first64Bits(uuidTime time.Time) uint64 {
	milliseconds := millisecondInUint64(uuidTime)
	version := versionInUint64
	randA := nanosecondAsRandAInUint64(uuidTime)
	sixtyFourBitField := (milliseconds << 16) | ((version << 12) & 0xF000) | (randA)
	return sixtyFourBitField
}

// last64Bits generate 64 bits of uuidv7, varient (2 bit) + randB (62 bit)
func last64Bits() (uint64, error) {
	randB, err := randBInUint64()

	if err != nil {
		return 0, err
	}

	sixtyFourBitField := variantInUint64<<62 | randB
	return sixtyFourBitField, nil
}

// FromTime use time.Time to generate UUIDv7
// RandB is random number
func FromTime(uuidTime time.Time) (UUIDv7, error) {
	uuid := UUIDv7{}

	firstSection := first64Bits(uuidTime)
	lastSection, err := last64Bits()

	if err != nil {
		return uuid, err
	}

	// uuid is [16]bytes, section is uint64
	// we can not use firstSection << 64 | lastSection
	// because firstSection will overflow and lost data
	binary.BigEndian.PutUint64(uuid[0:8], firstSection)
	binary.BigEndian.PutUint64(uuid[8:], lastSection)

	return uuid, nil
}

// New use time.Now to generate UUIDv7
func New() (UUIDv7, error) {
	return FromTime(time.Now())
}

// ZeroUUIDv7 generate [16]byte all 0 UUIDv7
func ZeroUUIDv7() UUIDv7 {
	return UUIDv7{}
}
