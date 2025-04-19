package server

import (
	"log"

	"github.com/miekg/dns"
	"github.com/tnutzmann/SuperSimpleDNS/internal/config"
	"github.com/tnutzmann/SuperSimpleDNS/internal/resolver"
)

var configuration config.Config = *readConfig()

var resolvers = []resolver.Resolver{
	&resolver.LocalResolver{Zones: configuration.Zones},
	&resolver.UpstreamResolver{UpstreamAddress: configuration.Upstream},
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

func readConfig() *config.Config {
	con, err := config.LoadConfigFromFile("./zones.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return con
}

func failureResponse(writer dns.ResponseWriter, msg *dns.Msg) {
	failMsg := new(dns.Msg)
	failMsg.SetRcode(msg, dns.RcodeServerFailure)
	writer.WriteMsg(failMsg)
}
