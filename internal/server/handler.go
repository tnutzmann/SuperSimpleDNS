package server

import (
	"log"

	"github.com/miekg/dns"
	"github.com/tnutzmann/SuperSimpleDNS/internal/config"
	"github.com/tnutzmann/SuperSimpleDNS/internal/resolver"
)

var zones = []config.Zone{
	config.Zone{Name: "example.com.", A: []string{"127.0.0.1", "127.0.0.2"}, AAAA: []string{"::1"}},
}

var resolvers = []resolver.Resolver{
	&resolver.LocalResolver{Zones: zones},
	&resolver.UpstreamResolver{UpstreamAddress: "8.8.8.8:53"},
}

func handleDNSRequest(writer dns.ResponseWriter, msg *dns.Msg) {
	var err error
	reply := dns.Msg{}
	reply.SetReply(msg)

	for _, resolver := range resolvers {
		reply.Answer, err = resolver.Resolve(msg)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Printf("Error: Failed to resolve DNS request for '%v'", msg.Question[0].Name)
		failureResponse(writer, msg)
	}

	if err := writer.WriteMsg(&reply); err != nil {
		log.Printf("Failed to write DNS response: %v", err)
	}
}

func failureResponse(writer dns.ResponseWriter, msg *dns.Msg) {
	failMsg := new(dns.Msg)
	failMsg.SetRcode(msg, dns.RcodeServerFailure)
	writer.WriteMsg(failMsg)
}
