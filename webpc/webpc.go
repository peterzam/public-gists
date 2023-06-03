package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/image/webp"
)

var (
	src = flag.String("s", "", "Source folder")
	n   = flag.Int("n", 8, "Concurrent")
)

func main() {
	flag.Parse()

	if len(*src) == 0 {
		flag.Usage()
		return
	}

	// Read Dir
	file_list, err := os.ReadDir(*src)
	if err != nil {
		log.Fatal(err)
	}

	// Check file
	reg, err := regexp.Compile(".webp$")
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan string, len(file_list))
	results := make(chan string, len(file_list))

	for w := 0; w < *n; w++ {
		go Convert(reg, jobs, results)
	}
	for _, job := range file_list {
		jobs <- job.Name()
	}
	close(jobs)
	for range file_list {
		fmt.Println(<-results)
	}
}

func Convert(reg *regexp.Regexp, jobs <-chan string, result chan<- string) {
	// If webp
	for file_name := range jobs {
		if reg.MatchString(file_name) {

			// Read file
			file_path := *src + "/" + file_name
			file, err := os.Open(file_path)
			if err != nil {
				log.Println(err)
			} else {
				os.Remove(file_path)
			}

			// Decode webp
			image, err := webp.Decode(file)
			if err != nil {
				log.Println(err)
			}
			file.Close()

			// Create png
			png_file, err := os.Create(strings.Split(file_path, ".")[0] + ".png")
			if err != nil {
				log.Println(err)
			}

			// Write to png
			if png.Encode(png_file, image) != nil {
				log.Println(err)
			}
			png_file.Close()
			// Show changes
			result <- (file_name + " -> " + strings.Split(file_name, ".")[0] + ".png")
		} else {
			result <- (file_name + " -> " + "Do not change")
		}
	}
}
