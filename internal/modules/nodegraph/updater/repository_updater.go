package updater

import (
	"sync"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/repository"
)

type RepositoryUpdater struct {
	repo repository.Repository[model.ConnectionItem]
	lock *sync.RWMutex
}

func NewUpdater(repo repository.Repository[model.ConnectionItem]) *RepositoryUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (updater *RepositoryUpdater) Update(src string, srcName string, srcNamespace string, dst string, dstName string, dstNamespace string, persistent bool, bytesSent float64, bytesReceived float64, duration float64, closed bool) {
	_ = "STUB: not implemented"
	return
}
