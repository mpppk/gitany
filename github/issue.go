package github

import (
	"github.com/google/go-github/github"
	"github.com/mpppk/gitany/service"
)

type Issue struct {
	*github.Issue
}

func (i *Issue) GetRepository() service.Repository {
	repo := i.Issue.GetRepository()
	return service.Repository(repo)
}

