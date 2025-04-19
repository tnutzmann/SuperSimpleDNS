package resolver

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
	"github.com/tnutzmann/SuperSimpleDNS/internal/config"
)

type LocalResolver struct {
	Zones []config.Zone
}

func (resolver *LocalResolver) Resolve(msg *dns.Msg) ([]dns.RR, error) {
	if len(msg.Question) > 1 {
		return nil, fmt.Errorf("ERROR: Too many DNS Questions")
	}

	q := msg.Question[0]
	for _, zone := range resolver.Zones {
		if zone.Name == q.Name {
			switch q.Qtype {
			case dns.TypeA:
				return resolver.buildRR(q.Name, zone.A, "A")
			case dns.TypeAAAA:
				return resolver.buildRR(q.Name, zone.AAAA, "AAAA")
			case dns.TypeCNAME:
				return resolver.buildRR(q.Name, []string{zone.CNAME}, "CNAME")
			case dns.TypeTXT:
				return resolver.buildRR(q.Name, zone.TXT, "TXT")
			case dns.TypeMX:
				return resolver.buildRR(q.Name, zone.MX, "MX")
			case dns.TypeNS:
				return resolver.buildRR(q.Name, zone.NS, "NS")
			case dns.TypeSOA:
				return resolver.buildRR(q.Name, []string{zone.SOA}, "SOA")
			case dns.TypeSRV:
				return resolver.buildRR(q.Name, zone.SRV, "SRV")
			case dns.TypePTR:
				return resolver.buildRR(q.Name, zone.PTR, "PTR")
			case dns.TypeCAA:
				return resolver.buildRR(q.Name, zone.CAA, "CAA")
			}
		}
	}
	return nil, fmt.Errorf("ERROR: DNS-Type not supported")
}

func (resolver *LocalResolver) buildRR(domainName string, records []string, dnsType string) (rr []dns.RR, err error) {
	for _, record := range records {
		r, err := dns.NewRR(fmt.Sprintf("%s %s %s", domainName, dnsType, record))
		if err != nil {
			log.Fatalf("ERROR: Something went wrong: %v", err)
			return nil, err
		}
		rr = append(rr, r)
	}
	return rr, nil
}
