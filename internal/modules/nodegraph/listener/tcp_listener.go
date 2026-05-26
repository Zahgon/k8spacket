package listener

import (
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/updater"

	"github.com/k8spacket/k8spacket/internal/modules"
)

type TcpListener struct {
	updater           updater.Updater
	tcpMetricsEnabled bool
}

func NewListener(updater updater.Updater) modules.Listener[modules.TCPEvent] {
	_ = "STUB: not implemented"
	return nil
}

func (listener *TcpListener) Listen(event modules.TCPEvent) { _ = "STUB: not implemented"; return }

func sendPrometheusMetrics(event modules.TCPEvent, persistent bool, tcpMetricsEnabled bool) {
	_ = "STUB: not implemented"
	return
}
