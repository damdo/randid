package randid

// code inspiration and credits to github.com/moby/moby/pkg/stringid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// DefaultLen is the default id length
var DefaultLen = 32
var hardLimitLen = 1000000

// ID generates an "defaultLen" long id
func ID() string {
	return generate(DefaultLen)
}

// SizedID generates an "size" long id
func SizedID(size int) string {
	return generate(size)
}

func generate(size int) string {
	if size > hardLimitLen {
		size = hardLimitLen
	}

	b := make([]byte, size)
	r := rand.Reader

	if _, err := io.ReadFull(r, b); err != nil {
		panic(err)
	}

	id := hex.EncodeToString(b)
	return truncateID(id, size)
}

func truncateID(id string, length int) string {
	trimTo := length
	if len(id) < length {
		trimTo = len(id)
	}
	return id[:trimTo]
}
