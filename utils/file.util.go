package utils

import "math/rand"

var carset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = carset[rand.Intn(len(carset))]
	}
	return string(b)
}