package orm

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

type GormProvider struct {
}

func (h *GormProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeGorm
}

func (h *GormProvider) Boot(container framework.Container) error {
	return nil
}

func (h *GormProvider) IsDefer() bool {
	return true
}

func (h *GormProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (h *GormProvider) Name() string {
	return contract.ORMKey
}
