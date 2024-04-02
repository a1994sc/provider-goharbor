package account

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("harbor_robot_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("harbor_user", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
	})
}
