package github

import (
	"time"

	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
)

type Issue struct {
	*github.Issue
}

func (i *Issue) GetDueDate() *time.Time {
	return nil
}

func (i *Issue) GetRepository() gitany.Repository {
	repo := i.Issue.GetRepository()
	return gitany.Repository(repo)
}

func (i *Issue) GetMilestone() gitany.Milestone {
	if i.Milestone == nil {
		return nil
	}

	return &Milestone{Milestone: i.Milestone}
}

func (i *Issue) GetRepositoryID() int64 {
	repo := i.Issue.GetRepository()
	return *repo.ID
}

func (i *Issue) GetLabels() (labelNames []string) {
	for _, label := range i.Labels {
		labelNames = append(labelNames, *label.Name)
	}
	return
}
