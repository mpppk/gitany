package gitlab

import (
	"context"
	"fmt"

	"strings"

	"github.com/mpppk/gitany"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	rawClient     RawClient
	host          string
	serviceConfig *gitany.ServiceConfig
	ListOptions   *gitlab.ListOptions
}

type ClientGenerator struct{}

func (c *Client) GetRepositories() gitany.RepositoriesService {
	return &repositoriesService{
		raw:           c.rawClient,
		host:          c.host,
		serviceConfig: c.serviceConfig,
	}
}

func (c *Client) GetIssues() gitany.IssuesService {
	return &issuesService{
		client:      c,
		rawClient:   c.rawClient,
		ListOptions: c.ListOptions,
	}
}

func (c *Client) GetPullRequests() gitany.PullRequestsService {
	return &pullRequestsService{
		raw:                 c.rawClient.GetMergeRequests(),
		repositoriesService: c.GetRepositories(),
		ListOptions:         c.ListOptions,
	}
}

func (c *Client) GetAuthorizations() gitany.AuthorizationsService {
	return &authorizationsService{}
}

func (c *Client) GetProjects() gitany.ProjectsService {
	return &projetsService{
		repositoriesService: c.GetRepositories(),
	}
}

func (cb *ClientGenerator) New(ctx context.Context, serviceConfig *gitany.ServiceConfig) (gitany.Client, error) {
	rawClient := newGitLabRawClient(serviceConfig)
	return newClientFromRawClient(serviceConfig, rawClient), nil
}

func (cb *ClientGenerator) NewViaBasicAuth(ctx context.Context, serviceConfig *gitany.ServiceConfig, user, pass string) (gitany.Client, error) {
	panic("gitlab.ClientGenerator.NewViaBasicAuth is not implemented yet")
}

func (cb *ClientGenerator) GetType() string {
	return "gitlab"
}

func newGitLabRawClient(serviceConfig *gitany.ServiceConfig) *rawClient {
	client := gitlab.NewClient(nil, serviceConfig.Token)
	client.SetBaseURL(serviceConfig.Protocol + "://" + serviceConfig.Host)
	return &rawClient{Client: client}
}

func newClientFromRawClient(serviceConfig *gitany.ServiceConfig, rawClient RawClient) gitany.Client {
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
