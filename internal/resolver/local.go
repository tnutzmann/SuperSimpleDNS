package resolver

import (
	"fmt"

	"github.com/miekg/dns"
)

type LocalResolver struct{}

func (*LocalResolver) Resolve(msg *dns.Msg) ([]dns.RR, error) {
	if len(msg.Question) > 1 {
		return nil, fmt.Errorf("ERROR: Too many DNS Questions")
	}

	q := msg.Question[0]

	switch q.Qtype {
	case dns.TypeA:
		if q.Name == "example.com." {
			rr, err := dns.NewRR(fmt.Sprintf("%s A 127.0.0.1", q.Name))

			if err != nil {
				return nil, err
			}
			return []dns.RR{rr}, nil
		}
	case dns.TypeAAAA:
		if q.Name == "example.com." {
			rr, err := dns.NewRR(fmt.Sprintf("%s AAAA ::1", q.Name))

			if err != nil {
				return nil, err
			}
			return []dns.RR{rr}, nil
		}
	}

	return nil, fmt.Errorf("ERROR: DNS-Type not supported")
}
