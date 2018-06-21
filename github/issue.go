package github

import (
	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
)

type Issue struct {
	*github.Issue
}

func (i *Issue) GetRepository() gitany.Repository {
	repo := i.Issue.GetRepository()
	return gitany.Repository(repo)
}
