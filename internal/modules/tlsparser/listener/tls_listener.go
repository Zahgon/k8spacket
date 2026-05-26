package listener

import (
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/storer"

	"github.com/k8spacket/k8spacket/internal/modules"
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
)

type TlsListener struct {
	storer                      storer.Storer
	tlsRecordsMeticsEnabled     bool
	tlsExpirationMetricsEnabled bool
}

func NewListener(storer storer.Storer) modules.Listener[modules.TLSEvent] {
	_ = "STUB: not implemented"
	return nil
}

func (listener *TlsListener) Listen(tlsEvent modules.TLSEvent) { _ = "STUB: not implemented"; return }

func sendPrometheusMetrics(tlsConnection model.TLSConnection, tlsDetails model.TLSDetails, tlsRecordsMeticsEnabled bool, tlsExpirationMetricsEnabled bool) {
	_ = "STUB: not implemented"
	return
}
