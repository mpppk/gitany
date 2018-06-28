package github

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/etc"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	VALIDATION_FAILED_MSG = "Validation Failed"
	NO_COMMITS_MSG_PREFIX = "No commits between"
	CODE_INVALID          = "invalid"
)

type Client struct {
	rawClient   RawClient
	host        string
	ListOptions *github.ListOptions
}

func (c *Client) GetRepositories() gitany.RepositoriesService {
	return &repositoriesService{
		raw:  c.rawClient.GetRepositories(),
		host: c.host,
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
		raw:                 c.rawClient.GetPullRequests(),
		repositoriesService: c.GetRepositories(),
		ListOptions:         c.ListOptions,
	}
}

func (c *Client) GetProjects() gitany.ProjectsService {
	return &projectsService{
		repositoriesService: c.GetRepositories(),
	}
}

func (c *Client) GetAuthorizations() gitany.AuthorizationsService {
	return &authorizationsService{
		raw: c.rawClient.GetAuthorizations(),
	}
}

type ClientGenerator struct{}

func (cb *ClientGenerator) New(ctx context.Context, serviceConfig *etc.ServiceConfig) (gitany.Client, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: serviceConfig.Token})
	tc := oauth2.NewClient(ctx, ts)
	return newServiceFromClient(serviceConfig, &rawClient{Client: github.NewClient(tc)})
}

func (cb *ClientGenerator) NewViaBasicAuth(ctx context.Context, serviceConfig *etc.ServiceConfig, user, pass string) (gitany.Client, error) {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(user),
		Password: strings.TrimSpace(pass),
	}
	return newServiceFromClient(serviceConfig, &rawClient{Client: github.NewClient(tp.Client())})
}

func (cb *ClientGenerator) GetType() string {
	return "github"
}

func newServiceFromClient(serviceConfig *etc.ServiceConfig, client RawClient) (gitany.Client, error) {
	urlStr := serviceConfig.Protocol + "://api." + serviceConfig.Host
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse base URL based on serviceConfig in github.Client.newServiceFromClient")
	}
	client.SetBaseURL(baseUrl)
	listOpt := &github.ListOptions{PerPage: 100}
	return gitany.Client(&Client{rawClient: client, host: serviceConfig.Host, ListOptions: listOpt}), nil
}

func hasAuthNote(auths []*github.Authorization, note string) bool {
	for _, a := range auths {
		if a.Note != nil && note == *a.Note {
			return true
		}
	}
	return false
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
