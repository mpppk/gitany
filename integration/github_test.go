// +build integration

package integration

import (
	"os"

	"testing"

	"context"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/github"
)

func TestGitHubIntegration(t *testing.T) {
	token := os.Getenv("GITANY_GITHUB_TEST_TOKEN")
	if token == "" {
		t.Log("!!! No OAuth token. Some tests won't run. !!!")
	}
	gitany.RegisterClientGenerator(&github.ClientGenerator{})

	serviceConfig := &gitany.ServiceConfig{
		Host:     "github.com",
		Type:     "github",
		Token:    token,
		Protocol: "https",
	}

	owner := "mpppk-test"
	repoName := "test-repo"
	org := "gitany-test-org"

	ctx := context.Background()

	client, err := gitany.GetClient(context.Background(), serviceConfig)
	if err != nil {
		t.Fatal(err)
	}

	repos, res, err := client.GetRepositories().ListByOrg(ctx, org, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(repos) <= 0 {
		t.Error("failed to fetch repositories from " + org)
	}

	repoIssues, res, err := client.GetIssues().ListByRepo(ctx, owner, repoName, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(repoIssues) <= 0 {
		t.Errorf("failed to fetch repositories from %s/%s", owner, repoName)
	}

	opt := &gitany.IssueListOptions{
		Filter: "all", // see https://developer.github.com/v3/issues/#list-issues
	}
	issues, res, err := client.GetIssues().ListByOrg(ctx, org, opt)
	if err != nil {
		t.Fatal(err, res.GetHTTPResponse())
	}

	if len(issues) <= 0 {
		t.Error("failed to fetch issues from " + org)
	}

	labels, res, err := client.GetIssues().ListLabels(ctx, owner, repoName, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(labels) <= 0 {
		t.Errorf("failed to fetch labels from %s/%s", owner, repoName)
	}
	// ListMilestonesByOrg always return err because group milestone does not implemented in github
	_, _, err = client.GetIssues().ListMilestonesByOrg(ctx, org, nil)
}
