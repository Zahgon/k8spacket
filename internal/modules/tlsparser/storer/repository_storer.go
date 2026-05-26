package storer

import (
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/repository"
	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/update"
)

type RepositoryStorer struct {
	repo    repository.Repository
	updater update.Updater
}

func NewStorer(repo repository.Repository, updater update.Updater) Storer {
	_ = "STUB: not implemented"
	return *new(Storer)
}

func (storer *RepositoryStorer) StoreInDatabase(tlsConnection *model.TLSConnection, tlsDetails *model.TLSDetails) {
	_ = "STUB: not implemented"
	return
}
