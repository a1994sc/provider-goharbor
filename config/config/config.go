package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_config_auth", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
