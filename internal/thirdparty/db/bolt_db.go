package db

import (
	tcp_model "github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
	tls_model "github.com/k8spacket/k8spacket/internal/modules/tlsparser/model"
	"github.com/timshannon/bolthold"
)

type BoltDb[T tls_model.TLSDetails | tls_model.TLSConnection | tcp_model.ConnectionItem] struct {
	store *bolthold.Store
}

func New[T tls_model.TLSDetails | tls_model.TLSConnection | tcp_model.ConnectionItem](dbname string) (Db[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (boltDb *BoltDb[T]) Close() error { _ = "STUB: not implemented"; return nil }

func (boltDb *BoltDb[T]) Read(key string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (boltDb *BoltDb[T]) Query(query *bolthold.Query) ([]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (boltDb *BoltDb[T]) QueryMatchFunc(field string, matchFunc func(*T) (bool, error)) bolthold.Query {
	_ = "STUB: not implemented"
	return *new(bolthold.Query)
}

func (boltDb *BoltDb[T]) Upsert(key string, value *T) error { _ = "STUB: not implemented"; return nil }

func HashId(s string) uint32 { _ = "STUB: not implemented"; return 0 }
