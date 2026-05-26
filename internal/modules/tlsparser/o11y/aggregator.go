package o11y

import (
	"context"

	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	httpclient "github.com/k8spacket/k8spacket/internal/thirdparty/http"
)

// aggregateTLSResponses fetches TLS responses from peer k8spacket pods concurrently and merges them.
func aggregateTLSResponses[T model.TLSDetails | []model.TLSConnection](ctx context.Context, podIPs []string, urlTemplate string, client httpclient.Client, zero T, merge func(dst T, src T) T) (T, []error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
