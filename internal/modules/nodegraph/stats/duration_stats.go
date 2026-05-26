package stats

import (
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
)

type DurationStats struct {
	Stats
}

func (duration *DurationStats) GetConfig() model.Config {
	_ = "STUB: not implemented"
	return *new(model.Config)
}

func (duration *DurationStats) FillNodeStats(node *model.Node, connEndpoint model.ConnectionEndpoint) {
	_ = "STUB: not implemented"
	return
}

func (duration *DurationStats) FillEdgeStats(edge *model.Edge, connItem model.ConnectionItem) {
	_ = "STUB: not implemented"
	return
}
