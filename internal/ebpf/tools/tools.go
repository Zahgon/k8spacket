package ebpf_tools

import (
	"os"
	"regexp"
	"sync"

	"github.com/k8spacket/k8spacket/internal/modules"
)

const (
	id_format string = "%s-%d"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

var domainsMap = &SafeMap{data: make(map[string]string)}
var reverseLookupMap = &SafeMap{data: make(map[string]string)}
var reReverseWhois = regexp.MustCompile(os.Getenv("K8S_PACKET_REVERSE_WHOIS_REGEXP"))

func EnrichAddress(addr *modules.Address) { _ = "STUB: not implemented"; return }

// try to find domain (https only), organization name and (if GeoLite2 Free Geolocation Data enabled) country and city by external IP
func reverseLookup(ip string, port uint16) string { _ = "STUB: not implemented"; return "" }

// Check if an IP is private.
func privateIPCheck(ip string) bool { _ = "STUB: not implemented"; return false }

func StoreDomain(ip string, port uint16, domain string) { _ = "STUB: not implemented"; return }

func IntToIP4(ipNum uint32, fn func(b []byte, c uint32)) string {
	_ = "STUB: not implemented"
	return ""
}

func Htons(v uint16) uint16 { _ = "STUB: not implemented"; return 0 }
