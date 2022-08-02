package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func myId() string {
	f, err := os.Open("clifford")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	cliparser()
}
