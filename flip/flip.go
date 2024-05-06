package main

import (
	"encoding/binary"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	crand "crypto/rand"
	rand "math/rand"

	"github.com/labstack/echo"
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func main() {
	e := echo.New()
	e.GET(os.Getenv("LISTEN_PATH"), func(c echo.Context) error {
		var src cryptoSource
		var qbit = readCsv("qbit.csv")
		if qbit[rand.New(src).Intn(len(qbit))] != (rand.New(src).Intn(2) == 0) {
			return c.String(http.StatusOK, "head")
		} else {
			return c.String(http.StatusOK, "tail")
		}
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("LISTEN_PORT")))
}

func readCsv(filename string) []bool {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	values := strings.Split(strings.ReplaceAll(string(f), "\r\n", "\n"), "\n")
	arr := make([]bool, len(values))
	for i, val := range values {
		b, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		arr[i] = b
	}
	return arr
}
