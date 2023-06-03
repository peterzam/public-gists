package Gorand

import (
	"math/rand"
	"time"
)

var base string = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "0123456789"

func New(length int) string {
	charset := base
	return (RandStringGen(length, charset))

}

func Special(length int) string {
	charset := base + "`!@#$%^&*()-_+={}|?/:;'<>,." + "\\" + "\""
	return (RandStringGen(length, charset))
}

func Base64(length int) string {
	charset := base + "+/"
	return (RandStringGen(length, charset) + "=")
}

func RandStringGen(length int, charset string) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))]
	}
	return (string(b))
}
