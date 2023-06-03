package main

import (
	"flag"
	"log"

	"inet.af/tcpproxy"
)

var (
	to   string
	port string
)

func init() {
	flag.StringVar(&to, "to", "localhost:30000", "the target (<host>:<port>)")
	flag.StringVar(&port, "port", "50000", "the tunnelthing port")
}

func main() {
	var p tcpproxy.Proxy
	p.AddSNIRoute(":"+port, "localhost", tcpproxy.To(to))
	p.AddRoute(":"+port, tcpproxy.To(to)) // fallback
	log.Fatal(p.Run())
}
