package tlsparser

import (
	"net/http"

	"github.com/k8spacket/k8spacket/internal/modules"
)

func Init(mux *http.ServeMux) modules.Listener[modules.TLSEvent] {
	_ = "STUB: not implemented"
	return nil
}
