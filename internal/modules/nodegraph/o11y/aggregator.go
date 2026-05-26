package o11y

import (
	"context"
	"net/url"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/stats"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	httpclient "github.com/k8spacket/k8spacket/internal/thirdparty/http"
)

func aggregateConnections(ctx context.Context, podIPs []string, query url.Values, port string, client httpclient.Client) []model.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}

func prepareConnections(connectionItems map[string]model.ConnectionItem, connectionEndpoints map[string]model.ConnectionEndpoint) {
	_ = "STUB: not implemented"
	return
}

func buildApiResponse(connectionItems map[string]model.ConnectionItem, connectionEndpoints map[string]model.ConnectionEndpoint, statsImpl stats.Stats) model.NodeGraph {
	_ = "STUB: not implemented"
	return *new(model.NodeGraph)
}

func fillNodesArray(id string, nodeArray []model.Node, connectionEndpoints map[string]model.ConnectionEndpoint, statsImpl stats.Stats) []model.Node {
	_ = "STUB: not implemented"
	return nil
}

func fillEdgesArray(id string, edgeArray []model.Edge, connectionItems map[string]model.ConnectionItem, statsImpl stats.Stats) []model.Edge {
	_ = "STUB: not implemented"
	return nil
}
