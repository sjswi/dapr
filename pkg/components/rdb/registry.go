package rdb

import (
	"fmt"
	"github.com/dapr/components-contrib/rdb"
	"github.com/dapr/kit/logger"
	"strings"
)

// Registry is the interface of a components that allows callers to get registered instances of input and output bindings.
type Registry struct {
	Logger logger.Logger
	tests  map[string]func(logger.Logger) rdb.RDB
}

func (r *Registry) RegisterComponent(factory func(logger.Logger) rdb.RDB, ss string) {
	r.tests[ss] = factory
}

func (r *Registry) Create(name, version, logName string) (rdb.RDB, error) {
	nameLower := strings.ToLower(name)
	versionLower := strings.ToLower(version)
	key := nameLower + "/" + versionLower
	if method, ok := r.tests[key]; ok {
		l := r.Logger
		if logName != "" && l != nil {
			l = l.WithFields(map[string]any{
				"component": logName,
			})
		}
		return method(l), nil
	}
	return nil, fmt.Errorf("couldn't find test  %s/%s", name, version)
}

// DefaultRegistry is the singleton with the registry.
var DefaultRegistry *Registry = NewRegistry()

func NewRegistry() *Registry {

	return &Registry{
		tests: map[string]func(logger.Logger) rdb.RDB{},
	}

}
