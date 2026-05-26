package ebpf

import (
	ebpf_inet "github.com/k8spacket/k8spacket/internal/ebpf/inet"
	ebpf_socketfilter "github.com/k8spacket/k8spacket/internal/ebpf/socketfilter"
	ebpf_tc "github.com/k8spacket/k8spacket/internal/ebpf/tc"
)

type EbpfLoader struct {
	inetEbpf         ebpf_inet.Inet
	tcEbpf           ebpf_tc.Tc
	socketFilterEbpf ebpf_socketfilter.SocketFilter
	interfaces       []string
}

func Init(inetEbpf ebpf_inet.Inet, tcEbpf ebpf_tc.Tc, socketFilterEbpf ebpf_socketfilter.SocketFilter) *EbpfLoader {
	_ = "STUB: not implemented"
	return nil
}

func (loader *EbpfLoader) Load() {
	_ = "STUB: not implemented"
	// load inet_sock_set_state ebpf program
	return
}

func interfacesRefresher(loader EbpfLoader) { _ = "STUB: not implemented"; return }

// load traffic control ebpf program (qdisc filter)

// looking for network interfaces on cluster nodes regarding started containers based on the command `ip address`
func findInterfaces() []string { _ = "STUB: not implemented"; return nil }
