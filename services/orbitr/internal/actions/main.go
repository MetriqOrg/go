package actions

import "github.com/metriqorg/go/services/orbitr/internal/corestate"

type CoreStateGetter interface {
	GetCoreState() corestate.State
}
