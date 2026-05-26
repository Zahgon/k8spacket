package main

import (
	"net/http"

	"github.com/k8spacket/k8spacket/internal/broker"
	"github.com/k8spacket/k8spacket/internal/ebpf"
	ebpf_inet "github.com/k8spacket/k8spacket/internal/ebpf/inet"
	ebpf_socketfilter "github.com/k8spacket/k8spacket/internal/ebpf/socketfilter"
	ebpf_tc "github.com/k8spacket/k8spacket/internal/ebpf/tc"
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph"
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser"
)

func main() {

	mux := http.NewServeMux()

	nodegraphListener := nodegraph.Init(mux)
	tlsParserListener := tlsparser.Init(mux)
	distributionBroker := broker.Init(nodegraphListener, tlsParserListener)

	inetEbpf := &ebpf_inet.EbpfInet{Broker: distributionBroker}
	tcEbpf := &ebpf_tc.EbpfTc{Broker: distributionBroker}
	socketFilterEbpf := &ebpf_socketfilter.EbpfSocketFilter{Broker: distributionBroker}
	loader := ebpf.Init(inetEbpf, tcEbpf, socketFilterEbpf)

	buildLogger()
	startApp(distributionBroker, loader, mux)
}

func startApp(broker broker.Broker, loader ebpf.Loader, mux *http.ServeMux) {
	_ = "STUB: not implemented"
	return
}

func startHttpServer(mux *http.ServeMux) { _ = "STUB: not implemented"; return }

// graceful shutdown

func buildLogger() { _ = "STUB: not implemented"; return }
