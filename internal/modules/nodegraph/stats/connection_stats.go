package stats

import (
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
)

type ConnectionStats struct {
	Stats
}

func (connection *ConnectionStats) GetConfig() model.Config {
	_ = "STUB: not implemented"
	return *new(model.Config)
}

func (connection *ConnectionStats) FillNodeStats(node *model.Node, connEndpoint model.ConnectionEndpoint) {
	_ = "STUB: not implemented"
	return
}

func (connection *ConnectionStats) FillEdgeStats(edge *model.Edge, connItem model.ConnectionItem) {
	_ = "STUB: not implemented"
	return
}
