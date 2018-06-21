package github

import (
	"github.com/google/go-github/github"
)

type Label struct {
	*github.Label
}

func (l *Label) GetName() string {
	return l.GetName()
}
