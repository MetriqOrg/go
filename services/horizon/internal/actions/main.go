package actions

import "github.com/lantah/go/services/orbitr/internal/corestate"

type CoreStateGetter interface {
	GetCoreState() corestate.State
}
