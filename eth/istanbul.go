package eth

import (
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

func istanbulProtocols(
	run func(p *p2p.Peer, rw p2p.MsgReadWriter, version int, handler HandlerFunc) error,
	nodeInfo func() interface{},
	peerInfo func(id discover.NodeID) interface{}) (protocols []p2p.Protocol) {
	version := 15

	protocols = append(protocols, p2p.Protocol{
		Name:    "istanbul",
		Version: uint(version),
		Length:  15,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			return run(p, rw, int(version), HandlerFunc(istanbul.HandleMsg))
		},
		NodeInfo: nodeInfo,
		PeerInfo: peerInfo,
	})

	return protocols
}
