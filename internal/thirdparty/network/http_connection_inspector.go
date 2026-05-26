package network

import (
	"crypto/x509"
)

type HttpConnectionInspector struct {
	ConnectionInspector
}

func (inspector *HttpConnectionInspector) IsDomainReachable(domain string) bool {
	_ = "STUB: not implemented"
	return false
}

func (inspector *HttpConnectionInspector) GetPeerCertificates(address string, port uint16) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
