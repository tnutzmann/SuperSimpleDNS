package server

import (
	"log"

	"github.com/miekg/dns"
)

func Start() {
	dns.HandleFunc(".", handleDNSRequest)

	udpServer := &dns.Server{Addr: ":53", Net: "udp"}
	log.Println("DNS server started on port 53...")
	log.Fatal(udpServer.ListenAndServe())
}
