package broker

import (
	"github.com/k8spacket/k8spacket/internal/modules"
)

type DistributionBroker struct {
	Broker
	NodegraphListener modules.Listener[modules.TCPEvent]
	TlsParserListener modules.Listener[modules.TLSEvent]
	tcpEventChannel   chan modules.TCPEvent
	tlsEventChannel   chan modules.TLSEvent
}

func Init(nodegraphListener modules.Listener[modules.TCPEvent], tlsParserListener modules.Listener[modules.TLSEvent]) *DistributionBroker {
	_ = "STUB: not implemented"
	return nil
}

func (broker *DistributionBroker) TCPEvent(event modules.TCPEvent) {
	_ = "STUB: not implemented"
	return
}

func (broker *DistributionBroker) TLSEvent(event modules.TLSEvent) {
	_ = "STUB: not implemented"
	return
}

func (broker *DistributionBroker) DistributeEvents() { _ = "STUB: not implemented"; return }
