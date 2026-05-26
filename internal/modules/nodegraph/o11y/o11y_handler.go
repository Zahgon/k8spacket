package o11y

import (
	"net/http"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/stats"
	httpclient "github.com/k8spacket/k8spacket/internal/thirdparty/http"
	k8sclient "github.com/k8spacket/k8spacket/internal/thirdparty/k8s"
	"github.com/k8spacket/k8spacket/internal/thirdparty/resource"
)

type O11yHandler struct {
	factory    stats.Factory
	httpClient httpclient.Client
	k8sClient  k8sclient.Client
	resource   resource.Resource
}

func NewO11yHandler(factory stats.Factory, httpClient httpclient.Client, k8sClient k8sclient.Client, resource resource.Resource) *O11yHandler {
	_ = "STUB: not implemented"
	return nil
}

func (handler *O11yHandler) Health(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *O11yHandler) NodeGraphFieldsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *O11yHandler) NodeGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *O11yHandler) getO11yStatsConfig(statsType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (handler *O11yHandler) buildO11yResponse(r *http.Request) (model.NodeGraph, error) {
	_ = "STUB: not implemented"
	return *new(model.NodeGraph), nil
}
