package github

import "github.com/mpppk/gitany"

func init() {
	gitany.RegisterClientGenerator(&ClientGenerator{})
}
