// Package modules
package modules

import (
	"go.uber.org/fx"
)

type AppModule struct {
	Name     string
	FxOption fx.Option
}

var registry []AppModule

func Register(m fx.Option, name string) {
	registry = append(registry, AppModule{
		name,
		m,
	})
}

// Meta returns all registred modules names
func Meta() []string {
	modsNames := []string{}
	for _, v := range registry {
		modsNames = append(modsNames, v.Name)
	}

	return modsNames
}

// Modules returns all registred fx modules
func Modules() []fx.Option {
	mods := []fx.Option{}
	for _, v := range registry {
		mods = append(mods, v.FxOption)
	}
	return mods
}
