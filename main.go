package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func Generate(p string, k string) string {
	h := hmac.New(sha256.New, []byte(k))
	h.Write([]byte(p))

	res := hex.EncodeToString(h.Sum(nil))
	log.Printf("\n secret_key = %v\n password = %v\n hash = %v\n len = %v\n", k, p, res, len(res))

	return res
}

func main() {
	Generate("password", "secret_key")
}
