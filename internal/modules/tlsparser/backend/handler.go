package backend

import (
	"net/http"
	"net/url"

	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/repository"

	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler { _ = "STUB: not implemented"; return nil }

func (handler *Handler) TLSConnectionHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (handler *Handler) getConnection(id string) model.TLSDetails {
	_ = "STUB: not implemented"
	return *new(model.TLSDetails)
}

func (handler *Handler) filterConnections(query url.Values) []model.TLSConnection {
	_ = "STUB: not implemented"
	return nil
}
