package ethereum

import (
	"github.com/Anurupa03/testing1/chainService/chain"
	"github.com/Anurupa03/testing1/chainService/relay"
)

var completedCh chan *chain.Packet
var retryCh chan *chain.Packet

func init() {
	relay.RegisteredClients["ethereum"] = NewClient
	completedCh = make(chan *chain.Packet)
	relay.RegisteredCompleteChannels["ethereum"] = completedCh
	retryCh = make(chan *chain.Packet)
	relay.RegisteredRetryChannels["ethereum"] = retryCh
}
