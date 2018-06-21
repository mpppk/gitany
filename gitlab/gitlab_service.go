package gitlab

import (
	"context"
	"fmt"

	"strings"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/etc"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	rawClient     RawClient
	host          string
	serviceConfig *etc.ServiceConfig
	ListOptions   *gitlab.ListOptions
}

type ClientGenerator struct{}

func (c *Client) GetRepositories() gitany.RepositoriesService {
	return gitany.RepositoriesService(&repositoriesService{
		raw:           c.rawClient,
		host:          c.host,
		serviceConfig: c.serviceConfig,
	})
}

func (c *Client) GetIssues() gitany.IssuesService {
	return gitany.IssuesService(&issuesService{
		raw:             c.rawClient.GetIssues(),
		projectsService: c.GetRepositories(),
		ListOptions:     c.ListOptions,
	})
}

func (c *Client) GetPullRequests() gitany.PullRequestsService {
	return gitany.PullRequestsService(&pullRequestsService{
		raw:                 c.rawClient.GetMergeRequests(),
		repositoriesService: c.GetRepositories(),
		ListOptions:         c.ListOptions,
	})
}

func (c *Client) GetAuthorizations() gitany.AuthorizationsService {
	return gitany.AuthorizationsService(&authorizationsService{})
}

func (c *Client) GetProjects() gitany.ProjectsService {
	return gitany.ProjectsService(&projetsService{
		repositoriesService: c.GetRepositories(),
	})
}

func (cb *ClientGenerator) New(ctx context.Context, serviceConfig *etc.ServiceConfig) (gitany.Client, error) {
	rawClient := newGitLabRawClient(serviceConfig)
	return newClientFromRawClient(serviceConfig, rawClient), nil
}

func (cb *ClientGenerator) NewViaBasicAuth(ctx context.Context, serviceConfig *etc.ServiceConfig, user, pass string) (gitany.Client, error) {
	panic("gitlab.ClientGenerator.NewViaBasicAuth is not implemented yet")
}

func (cb *ClientGenerator) GetType() string {
	return "gitlab"
}

func newGitLabRawClient(serviceConfig *etc.ServiceConfig) *rawClient {
	client := gitlab.NewClient(nil, serviceConfig.Token)
	client.SetBaseURL(serviceConfig.Protocol + "://" + serviceConfig.Host)
	return &rawClient{Client: client}
}

func newClientFromRawClient(serviceConfig *etc.ServiceConfig, rawClient RawClient) gitany.Client {
	listOpt := &gitlab.ListOptions{PerPage: 100}
	return &Client{rawClient: rawClient, serviceConfig: serviceConfig, host: serviceConfig.Host, ListOptions: listOpt}
}

func checkOwnerAndRepo(owner, repo string) error {
	m := map[string]string{"owner": owner, "repo": repo}

	for strType, name := range m {
		if err := validateOwnerOrRepo(strType, name); err != nil {
			return err
		}
	}
	return nil
}

func validateOwnerOrRepo(strType, name string) error {
	if name == "" {
		return errors.New(strType + " is empty")
	}
	if strings.Contains(name, "/") {
		return errors.New(fmt.Sprintf("invalid %v: %v", strType, name))
	}
	return nil
}
