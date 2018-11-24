package mock

import (
	"context"

	"github.com/mpppk/gitany"
)

type ClientGenerator struct {
	Client *Client
	Type   string
}

func (c *ClientGenerator) New(ctx context.Context, serviceConfig *gitany.ServiceConfig) (gitany.Client, error) {
	return c.Client, nil
}

func (c *ClientGenerator) NewViaBasicAuth(ctx context.Context, serviceConfig *gitany.ServiceConfig, username, pass string) (gitany.Client, error) {
	return c.Client, nil
}

func (c *ClientGenerator) GetType() string {
	return c.Type
}

type Client struct {
	Repositories   *repositoriesService
	PullRequests   *pullRequestsService
	Issues         *issuesService
	Projects       *projectsService
	Authorizations *authorizationsService
}

func NewClient() *Client {
	return &Client{
		Repositories:   NewRepositoriesSerivce(),
		PullRequests:   NewPullRequestsService(),
		Issues:         NewIssuesService(),
		Projects:       NewProjectsService(),
		Authorizations: NewAuthorizationService(),
	}
}

func (c *Client) GetRepositories() gitany.RepositoriesService {
	return c.Repositories
}

func (c *Client) GetPullRequests() gitany.PullRequestsService {
	return c.PullRequests
}

func (c *Client) GetIssues() gitany.IssuesService {
	return c.Issues
}

func (c *Client) GetProjects() gitany.ProjectsService {
	return c.Projects
}

func (c *Client) GetAuthorizations() gitany.AuthorizationsService {
	return c.Authorizations
}
