package repository

import (
	"regexp"
	"time"

	"github.com/k8spacket/k8spacket/internal/modules/nodegraph/model"
)

type Repository[T model.ConnectionItem] interface {
	Read(key string) T
	Query(from time.Time, to time.Time, patternNs *regexp.Regexp, patternIn *regexp.Regexp, patternEx *regexp.Regexp) []T
	Set(key string, value *T)
}
