package o11y

import (
	"net/http"

	httpclient "github.com/k8spacket/k8spacket/internal/thirdparty/http"
	k8sclient "github.com/k8spacket/k8spacket/internal/thirdparty/k8s"

	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
)

const connectionDetailsUri = "/tlsparser/api/data/"

type O11yHandler struct {
	httpClient httpclient.Client
	k8sClient  k8sclient.Client
}

func NewO11yHandler(httpClient httpclient.Client, k8sClient k8sclient.Client) *O11yHandler {
	_ = "STUB: not implemented"
	return nil
}

func (handler *O11yHandler) TLSParserConnectionsHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *O11yHandler) TLSParserConnectionDetailsHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *O11yHandler) buildConnectionsResponse(url string) []model.TLSConnection {
	_ = "STUB: not implemented"
	return nil
}

func (handler *O11yHandler) buildDetailsResponse(url string) model.TLSDetails {
	_ = "STUB: not implemented"
	return *new(model.TLSDetails)
}

func buildResponse[T model.TLSDetails | []model.TLSConnection](handler *O11yHandler, url string, t T, resultFunc func(d T, s T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func prepareResponse[T model.TLSDetails | []model.TLSConnection](w http.ResponseWriter, out T) {
	_ = "STUB: not implemented"
	return
}
