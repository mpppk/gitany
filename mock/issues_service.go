package mock

import (
	"context"

	"github.com/mpppk/gitany"
)

type issuesService struct {
	RepoIssues    []gitany.Issue
	OrgIssues     []gitany.Issue
	OrgMilestones []gitany.Milestone
	RepoLabels    []gitany.Label
	IssueURL      string
	URL           string
}

func NewIssuesService() *issuesService {
	return &issuesService{}
}

func (is *issuesService) ListByRepo(ctx context.Context, owner, repo string, opt *gitany.IssueListByRepoOptions) ([]gitany.Issue, gitany.Response, error) {
	return is.RepoIssues, nil, nil
}

func (is *issuesService) ListByOrg(ctx context.Context, org string, opt *gitany.IssueListOptions) ([]gitany.Issue, gitany.Response, error) {
	return is.OrgIssues, nil, nil
}

func (is *issuesService) ListMilestonesByOrg(ctx context.Context, org string, opt *gitany.MilestoneListOptions) ([]gitany.Milestone, gitany.Response, error) {
	return is.OrgMilestones, nil, nil
}

func (is *issuesService) ListLabels(ctx context.Context, owner string, repo string, opt *gitany.ListOptions) ([]gitany.Label, gitany.Response, error) {
	return is.RepoLabels, nil, nil
}

func (is *issuesService) GetIssuesURL(owner, repo string) (string, error) {
	return is.IssueURL, nil
}

func (is *issuesService) GetURL(owner, repo string, no int) (string, error) {
	return is.URL, nil
}
