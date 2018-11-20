package github

import "github.com/mpppk/gitany"

func init() {
	gitany.RegisterClientGenerator(&ClientGenerator{})

	gitany.RegisterDefaultServiceConfig(&gitany.ServiceConfig{
		Host:     "github.com",
		Type:     "github",
		Protocol: "https",
	})
}
