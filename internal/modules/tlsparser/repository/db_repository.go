package repository

import (
	"time"

	"github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	"github.com/k8spacket/k8spacket/internal/thirdparty/db"
)

type DbRepository struct {
	dbConnectionHandler db.Db[model.TLSConnection]
	dbDetailsHandler    db.Db[model.TLSDetails]
}

func NewDbRepository(db db.Db[model.TLSConnection], dbDetails db.Db[model.TLSDetails]) *DbRepository {
	_ = "STUB: not implemented"
	return nil
}

func (repository *DbRepository) Query(from time.Time, to time.Time) []model.TLSConnection {
	_ = "STUB: not implemented"
	return nil
}

func (repository *DbRepository) UpsertConnection(key string, value *model.TLSConnection) {
	_ = "STUB: not implemented"
	return
}

func (repository *DbRepository) Read(key string) model.TLSDetails {
	_ = "STUB: not implemented"
	return *new(model.TLSDetails)
}

//can happen, silent

type Fn func(newValue *model.TLSDetails, oldValue *model.TLSDetails)

func (repository *DbRepository) UpsertDetails(key string, value *model.TLSDetails, fn Fn) {
	_ = "STUB: not implemented"
	return
}
