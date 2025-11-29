package uuidv7

import (
	"encoding/hex"
)

// UUIDv7 implement [rfc9562](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7) uuidV7
// the "randA" part is replaced by the number of nanoseconds that project to 12bit in the current millisecond
type UUIDv7 [16]byte

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (u UUIDv7) String() string {
	var buf [36]byte
	encodeHex(buf[:], u)
	return string(buf[:])
}

// Ref: https://github.com/google/uuid/blob/v1.6.0/uuid.go#L359
// dst must be 36 bytes (uuidv7 16bytes => 32 hex => plus 4 "-" => 36 bytes)
func encodeHex(dst []byte, uuid UUIDv7) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

// Parse uuid string to uuidv7
func Parse(uuidStr string) (UUIDv7, error) {
	if len(uuidStr) != 36 {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	if uuidStr[8] != '-' || uuidStr[13] != '-' || uuidStr[18] != '-' || uuidStr[23] != '-' {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	uuidv7Buf := [16]byte{}

	if _, err := hex.Decode(uuidv7Buf[:4], []byte(uuidStr[:8])); err != nil {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	if _, err := hex.Decode(uuidv7Buf[4:6], []byte(uuidStr[9:13])); err != nil {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	if _, err := hex.Decode(uuidv7Buf[6:8], []byte(uuidStr[14:18])); err != nil {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	if _, err := hex.Decode(uuidv7Buf[8:10], []byte(uuidStr[19:23])); err != nil {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	if _, err := hex.Decode(uuidv7Buf[10:], []byte(uuidStr[24:])); err != nil {
		return UUIDv7{}, ErrInvalidUUIDFormat
	}

	return UUIDv7(uuidv7Buf), nil
}

// IsZero check if uuid is zero (empty [16]byte)
func (u UUIDv7) IsZero() bool {
	return u == (UUIDv7{})
}
