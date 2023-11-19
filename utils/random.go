package utils

import (
	"crypto/rand"
	"strconv"
)


// generate random 6 digit code
func GenCode() (int, error) {
	var err error
	codes := make([]byte, 6)
	if _, err = rand.Read(codes); err != nil {
		return 0, err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	i, err := strconv.Atoi(string(codes))
	return i, err
}
