package repository

import (
	"regexp"
	"time"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	"github.com/k8spacket/k8spacket/internal/thirdparty/db"
)

type DbRepository struct {
	dbHandler db.Db[model.ConnectionItem]
}

func NewDbRepository(db db.Db[model.ConnectionItem]) *DbRepository {
	_ = "STUB: not implemented"
	return nil
}

func (repository *DbRepository) Read(key string) model.ConnectionItem {
	_ = "STUB: not implemented"
	return *new(model.ConnectionItem)
}

// can happen, silent

func (repository *DbRepository) Query(from time.Time, to time.Time, patternNs *regexp.Regexp, patternIn *regexp.Regexp, patternEx *regexp.Regexp) []model.ConnectionItem {
	_ = "STUB: not implemented"
	return nil
}

func (repository *DbRepository) Set(key string, value *model.ConnectionItem) {
	_ = "STUB: not implemented"
	return
}
