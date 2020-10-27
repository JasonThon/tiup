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

// Drop all incoming and outgoing traffic
func (p *Partition) PauseTx() {
	p.server.PauseTx()
}

func (p *Partition) UnpauseTx() {
	p.server.UnpauseTx()
}

// Stop accept new connection
func (p *Partition) PauseAccept() {
	p.server.PauseAccept()
}

func (p *Partition) UnpauseAccept() {
	p.server.UnpauseAccept()
}
