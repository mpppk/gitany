package gitany

import (
	"context"

	"github.com/mpppk/gitany/etc"
)

type NewPullRequest struct {
	Title      string
	Body       string
	BaseBranch string
	HeadBranch string
	BaseOwner  string
	HeadOwner  string
}

type RepositoriesService interface {
	Get(ctx context.Context, owner, repo string) (Repository, Response, error)
	GetURL(owner, repo string) (string, error)
	GetWikisURL(owner, repo string) (string, error)
	GetMilestonesURL(owner, repo string) (string, error)
	GetMilestoneURL(owner, repo string, no int) (string, error)
	GetCommitsURL(owner, repo string) (string, error)
	Create(ctx context.Context, org string, repo Repository) (Repository, Response, error)
	CreateRelease(ctx context.Context, owner, repo string, newRelease *NewRelease) (Release, Response, error)
	ListByOrg(ctx context.Context, org string, opt *RepositoryListByOrgOptions) ([]Repository, Response, error)
}

type IssuesService interface {
	ListByRepo(ctx context.Context, owner, repo string) ([]Issue, error)
	//ListByOrg(ctx context.Context, org string) ([]*Issue, error) // FIXME
	ListByOrg(ctx context.Context, org string, opt *IssueListOptions) ([]Issue, Response, error)
	ListLabels(ctx context.Context, owner string, repo string) ([]Label, error)
	GetIssuesURL(owner, repo string) (string, error)
	GetURL(owner, repo string, no int) (string, error)
}

type PullRequestsService interface {
	List(ctx context.Context, owner, repo string) ([]PullRequest, error)
	Create(ctx context.Context, repo string, pull *NewPullRequest) (PullRequest, error)
	GetPullRequestsURL(owner, repo string) (string, error)
	GetURL(owner, repo string, no int) (string, error)
}

type AuthorizationsService interface {
	CreateToken(ctx context.Context) (string, error)
}

type ProjectsService interface {
	GetProjectsURL(owner, repo string) (string, error)
	GetURL(owner, repo string, no int) (string, error)
}

type Client interface {
	GetRepositories() RepositoriesService
	GetPullRequests() PullRequestsService
	GetIssues() IssuesService
	GetProjects() ProjectsService
	GetAuthorizations() AuthorizationsService
}

type ClientGenerator interface {
	New(ctx context.Context, serviceConfig *etc.ServiceConfig) (Client, error)
	NewViaBasicAuth(ctx context.Context, serviceConfig *etc.ServiceConfig, username, pass string) (Client, error)
	GetType() string
}
