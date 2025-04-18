package resolver

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

type UpstreamResolver struct {
	UpstreamAddress string
}

func (resolver *UpstreamResolver) Resolve(original *dns.Msg) ([]dns.RR, error) {
	client := new(dns.Client)
	response, _, err := client.Exchange(original, resolver.UpstreamAddress)
	if err != nil || response == nil || len(response.Answer) == 0 {
		log.Printf("Upstream resolution failed: %v", err)
		return nil, fmt.Errorf("ERROR: Upstream resolution failed")
	}
	return response.Answer, nil
}
