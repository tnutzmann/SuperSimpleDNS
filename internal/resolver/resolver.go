package resolver

import "github.com/miekg/dns"

type Resolver interface {
	Resolve(request *dns.Msg) ([]dns.RR, error)
}
