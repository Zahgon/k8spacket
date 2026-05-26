package stats

import (
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
)

type BytesStats struct {
	Stats
}

func (bytes *BytesStats) GetConfig() model.Config {
	_ = "STUB: not implemented"
	return *new(model.Config)
}

func (bytes *BytesStats) FillNodeStats(node *model.Node, connEndpoint model.ConnectionEndpoint) {
	_ = "STUB: not implemented"
	return
}

func (bytes *BytesStats) FillEdgeStats(edge *model.Edge, connItem model.ConnectionItem) {
	_ = "STUB: not implemented"
	return
}
