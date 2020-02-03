// Package hmoj provides a function for turning arbitrary byte slices (hashes)
// into strings prefixed with a deterministic emoji
package hmoj

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// Encoder contains flags which specify how the hash will be encoded
type Encoder struct {
	IncludeShortHash bool
}

var std = &Encoder{IncludeShortHash: true}

// Encode turns a []byte into an emoji-prefixed hex string,
// using the standard emoji encoder
func Encode(b []byte) string {
	return std.Moji(b)
}

// Moji turns a []byte into an emoji, optionally with a short string representation
// of that []byte
func (e *Encoder) Moji(b []byte) string {
	var shorty string
	if len(b) < 8 {
		shorty = hex.EncodeToString(b)
	} else {
		shorty = hex.EncodeToString(b[:7])
	}

	dec, _ := strconv.ParseInt(shorty, 16, 64)
	n := int(dec) % len(emoji)

	if e.IncludeShortHash {
		return fmt.Sprintf("%s %s", emoji[n], shorty)
	}

	return emoji[n]
}
