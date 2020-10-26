package main

import "go.etcd.io/etcd/pkg/proxy"

type Partition struct {
	server 	proxy.Server
}

func NewPartition(cfg proxy.ServerConfig) *Partition {
	return &Partition {
		server: proxy.NewServer(cfg),
	}
}

func (p *Partition) PauseAccept() {
	p.server.PauseAccept()
}

func (p *Partition) UnpauseAccept() {
	p.server.UnpauseAccept()
}
