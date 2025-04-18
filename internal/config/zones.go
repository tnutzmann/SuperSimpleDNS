package config

type Config struct {
	Upstream string `yaml:"upstream"`
	Zones    []Zone `yaml:"zones"`
}

type Zone struct {
	Name  string   `yaml:"name"`
	A     []string `yaml:"A,omitempty"`
	AAAA  []string `yaml:"AAAA,omitempty"`
	CNAME string   `yaml:"CNAME,omitempty"`
	TXT   []string `yaml:"TXT,omitempty"`
	MX    []string `yaml:"MX,omitempty"`
	NS    []string `yaml:"NS,omitempty"`
	SOA   string   `yaml:"SOA,omitempty"`
	SRV   []string `yaml:"SRV,omitempty"`
	PTR   []string `yaml:"PTR,omitempty"`
	CAA   []string `yaml:"CAA,omitempty"`
}
