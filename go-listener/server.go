package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

var godaemon = flag.Bool("d", false, "Run background (default \"false\")")
var protocol = flag.String("t", "tcp4", "Protocol type")
var address = flag.String("a", "0.0.0.0:1000", "Listen address & port")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *godaemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}

func main() {

	switch *protocol {
	case "tcp4":
		TCP(*protocol, *address)

	case "tcp6":
		TCP(*protocol, *address)

	case "udp4":
		UDP(*protocol, *address)

	case "udp6":
		UDP(*protocol, *address)

	default:
		log.Fatal("Protocol error")
	}
}

func TCP(proto string, addr string) {
	l, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal("TCP Listen Error")
	}
	for {
		l.Accept()
	}
}

func UDP(proto string, addr string) {
	s, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal("UDP Resolve Address Error")
	}
	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal("UDP Listen error")
	}
	buffer := make([]byte, 1024)

	for {
		_, addr, _ := c.ReadFromUDP(buffer)

		t := time.Now()
		c.WriteToUDP([]byte(t.String()), addr)

	}
}
