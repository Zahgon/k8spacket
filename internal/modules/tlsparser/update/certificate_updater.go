package update

import (
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	"github.com/k8spacket/k8spacket/internal/thirdparty/network"
)

type CertificateUpdater struct {
	network network.ConnectionInspector
}

func NewUpdater(network network.ConnectionInspector) Updater {
	_ = "STUB: not implemented"
	return *new(Updater)
}

func (updater *CertificateUpdater) Update(newValue *model.TLSDetails, oldValue *model.TLSDetails) {
	_ = "STUB: not implemented"
	return
}

// do update when it is the first time or time to live is exceeded

func scrapeCertificate(updater *CertificateUpdater, tlsDetails *model.TLSDetails) {
	_ = "STUB: not implemented"
	return
}

// check if domain is valid, if not - use destination IP
