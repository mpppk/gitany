package gitlab

import (
	"github.com/mpppk/gitany"
)

func init() {
	gitany.RegisterClientGenerator(&ClientGenerator{})

	gitany.RegisterDefaultServiceConfig(&gitany.ServiceConfig{
		Host:     "gitlab.com",
		Type:     "gitlab",
		Protocol: "https",
	})
}
