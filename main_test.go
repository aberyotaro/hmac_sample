package main

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	p := "password"
	sk := "secret_key"
	e := "2729f9405917a2614f4eff56aa960069ea566e8ddf5c59373f90538c60519b3b"
	el := 64

	t.Run("Success", func(t *testing.T) {
		res := Generate(p, sk)

		if res != e {
			log.Printf("FAILED. Unexpected result. expect: %v, actual: %v", e, res)
		}

		if len(res) != el {
			log.Printf("FAILED. Unexpected result. expect: %v, actual: %v", el, res)
		}
	})
}
