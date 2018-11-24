package mock

import (
	"context"

	"github.com/mpppk/gitany"
)

type pullRequestsService struct {
	RepoPullRequests   []gitany.PullRequest
	CreatedPullRequest gitany.PullRequest
	PullRequestsURL    string
	PullRequestURL     string
}

func NewPullRequestsService() *pullRequestsService {
	return &pullRequestsService{}
}

func (ps *pullRequestsService) List(ctx context.Context, owner, repo string) ([]gitany.PullRequest, error) {
	return ps.RepoPullRequests, nil
}

func (ps *pullRequestsService) Create(ctx context.Context, repo string, pull *gitany.NewPullRequest) (gitany.PullRequest, error) {
	return ps.CreatedPullRequest, nil
}

func (ps *pullRequestsService) GetPullRequestsURL(owner, repo string) (string, error) {
	return ps.PullRequestsURL, nil
}

func (ps *pullRequestsService) GetURL(owner, repo string, no int) (string, error) {
	return ps.PullRequestURL, nil
}
