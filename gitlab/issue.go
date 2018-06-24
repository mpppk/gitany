package gitlab

import (
	"github.com/mpppk/gitany"
	"github.com/xanzy/go-gitlab"
)

type Issue struct {
	*gitlab.Issue
	project *gitlab.Project
}

func (issue *Issue) GetBody() string {
	return issue.Description
}

func (issue *Issue) GetID() int64 {
	return int64(issue.ID)
}

func (issue *Issue) GetNumber() int {
	return issue.IID
}

func (issue *Issue) GetRepository() gitany.Repository {
	return &Repository{Project: issue.project}
}

func (issue *Issue) GetRepositoryID() int64 {
	return int64(issue.ProjectID)
}

func (issue *Issue) GetTitle() string {
	return issue.Title
}

func (issue *Issue) GetLabels() []string {
	return issue.Labels
}

func (issue *Issue) GetHTMLURL() string {
	return issue.WebURL
}
