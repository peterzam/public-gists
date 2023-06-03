// ref - https://bruinsslot.jp/post/golang-crypto/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"math"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/scrypt"
	"golang.org/x/term"
)

type TOTP struct {
	OTP
	interval int
}

type OTP struct {
	secret string
	digits int
	hasher *Hasher
}

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

func NewDefaultTOTP(secret string) *TOTP {
	otp := NewOTP(secret, 6, nil)
	return &TOTP{OTP: otp, interval: 30}
}

func NewOTP(secret string, digits int, hasher *Hasher) OTP {
	if hasher == nil {
		hasher = &Hasher{
			HashName: "sha1",
			Digest:   sha1.New,
		}
	}
	return OTP{
		secret: secret,
		digits: digits,
		hasher: hasher,
	}
}

func (t *TOTP) Now() string {
	return t.generateOTP(int((int(time.Now().Unix())) / t.interval))
}

func (o *OTP) generateOTP(input int) string {
	if input < 0 {
		panic("input must be positive integer")
	}
	hasher := hmac.New(o.hasher.Digest, o.byteSecret())
	hasher.Write(Itob(input))
	hmacHash := hasher.Sum(nil)

	offset := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := ((int(hmacHash[offset]) & 0x7f) << 24) |
		((int(hmacHash[offset+1] & 0xff)) << 16) |
		((int(hmacHash[offset+2] & 0xff)) << 8) |
		(int(hmacHash[offset+3]) & 0xff)

	code = code % int(math.Pow10(o.digits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", o.digits), code)
}

func (o *OTP) byteSecret() []byte {
	missingPadding := len(o.secret) % 8
	if missingPadding != 0 {
		o.secret = o.secret + strings.Repeat("=", 8-missingPadding)
	}
	bytes, err := base32.StdEncoding.DecodeString(o.secret)
	if err != nil {
		panic("decode secret failed")
	}
	return bytes
}

func Itob(integer int) []byte {
	byteArr := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		byteArr[i] = byte(integer & 0xff)
		integer = integer >> 8
	}
	return byteArr
}

func Decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]
	key, _, err := DeriveKey(key, salt)
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	return key, salt, nil
}

func Key(password string, data string) string {

	p := []byte(password)
	d, _ := hex.DecodeString(data)

	t, err := Decrypt(p, d)
	if err != nil {
		log.Fatal(err)
	}
	return string(t)
}

func main() {
	data := ""
	fmt.Printf("Enter secret: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	secret := string(bytepw)

	fmt.Println(NewDefaultTOTP(Key(secret, data)).Now())
}
