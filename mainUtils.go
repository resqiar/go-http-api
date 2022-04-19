package main

import (
	"crypto/rand"
	"unsafe"
)

func generateRandomString(size int) string {
	alpha := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alpha[b[i]/5]
	}

	return *(*string)(unsafe.Pointer(&b))
}
