package aleo

import (
	"github.com/Anurupa03/testing1/chainService/chain"
	"github.com/Anurupa03/testing1/chainService/relay"
)

var completedCh chan *chain.Packet
var retryCh chan *chain.Packet

func init() {
	relay.RegisteredClients["aleo"] = NewClient
	completedCh = make(chan *chain.Packet)
	relay.RegisteredCompleteChannels["aleo"] = completedCh
	retryCh = make(chan *chain.Packet)
	relay.RegisteredRetryChannels["aleo"] = retryCh
}
