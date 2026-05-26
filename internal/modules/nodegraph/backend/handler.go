package backend

import (
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/repository"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
)

var reMatchAll = regexp.MustCompile("")

type Handler struct {
	repo repository.Repository[model.ConnectionItem]
}

func NewHandler(repo repository.Repository[model.ConnectionItem]) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (handler *Handler) ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *Handler) filterConnections(query url.Values) []model.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}

func (handler *Handler) getConnections(from time.Time, to time.Time, patternNs *regexp.Regexp, patternIn *regexp.Regexp, patternEx *regexp.Regexp) []model.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}
