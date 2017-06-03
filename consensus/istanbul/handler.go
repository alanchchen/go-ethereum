package istanbul

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

const (
	PBFTMsg = 100
)

func HandleMsg(peerID func() discover.NodeID, msg p2p.Msg, fallback func(p2p.Msg) error) error {
	// Handle the message depending on its contents
	switch {
	case msg.Code == PBFTMsg:
		_, err := peerID().Pubkey()
		if err != nil {
			return err
		}
		var data []byte
		if err := msg.Decode(&data); err != nil {
			return err
		}

		// call our handler or fallback
	}

	return fallback(msg)
}
