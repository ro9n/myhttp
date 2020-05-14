package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func checksum(text string) (string, error) {
	h := md5.New()

	if _, err := io.WriteString(h, text); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
