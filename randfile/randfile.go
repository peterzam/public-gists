package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames [1000]string
	var totalFiles int

	for i, f := range files {
		fileNames[i] = f.Name()
		totalFiles = i
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(fileNames[(rand.Intn(totalFiles))])
}
