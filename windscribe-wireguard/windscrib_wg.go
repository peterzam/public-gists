// Replace cookie string
// Run - "go run wg.go"
// ** Sometimes the config files cannot be imported because of long filenames. **
// ** You can rename the config and reimport. **

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var COOKIE string = "<replace with ws_session_auth_hash cookie here>"
var PORT string = "443"

var LOCATIONS [1000]string
var TOTAL int = 0

func init() {
	GetLocations()
	_ = os.Mkdir("./configs", 0700)
}

func main() {

start:
	for TOTAL > 0 {
		config := CurlReq(false, LOCATIONS[TOTAL-1])
		if config == "" {
			fmt.Println("Wait for 5 second @ - " + strings.Split(LOCATIONS[TOTAL-1], ":")[1] + " to wait timeout")
			time.Sleep(5 * time.Second)
			goto start
		} else {
			fmt.Println(strings.Split(LOCATIONS[TOTAL-1], ":")[1])
			WriteFile(strings.ReplaceAll(strings.Split(LOCATIONS[TOTAL-1], ":")[1]+".conf", " ", ""), config)
			TOTAL--
		}

	}
}

func CurlReq(getLocations bool, location string) string {

	params := url.Values{}
	params.Add("location", location)
	params.Add("port", PORT)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://windscribe.com/getconfig/wireguard", body)
	if err != nil {
		log.Fatal(err)
	}

	if getLocations == false {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("Cookie", "ws_session_auth_hash="+COOKIE)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""

}

func GetLocations() {
	response := CurlReq(true, "")
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(response))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "<option value=\"") {
			LOCATIONS[TOTAL] = strings.Split(strings.TrimLeft(line, "<option value=\""), "\"")[0]
			TOTAL++
		}
	}
}

func WriteFile(fileName string, content string) {

	f, err := os.OpenFile("./configs/"+fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Println(err)
	}
}
