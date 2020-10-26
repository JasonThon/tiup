package spec

import (
	"github.com/coreos/etcd/pkg/proxy"
	"net"
	"net/url"
	"strconv"
)

type Partition struct {
	server		proxy.Server
	Name		string
	from		url.URL
	fromPort	int
	to			url.URL
	toPort		int
}

func NewPartition(name string, cfg proxy.ServerConfig) *Partition {
	p := &Partition {
		server: proxy.NewServer(cfg),
		Name: name,
		from: cfg.From,
		to: cfg.To,
	}

	_, fromPort, err := net.SplitHostPort(cfg.From.Host)
	if err != nil {
		p.fromPort, _ = strconv.Atoi(fromPort)
	}
	_, toPort, err := net.SplitHostPort(cfg.To.Host)
	if err != nil {
		p.toPort, _ = strconv.Atoi(toPort)
	}

	return p
}
