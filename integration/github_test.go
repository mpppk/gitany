package integration

import (
	"os"

	"testing"

	"context"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/etc"
	"github.com/mpppk/gitany/github"
)

func TestGitHubIntegration(t *testing.T) {
	token := os.Getenv("GITANY_GITHUB_TEST_TOKEN")
	if token == "" {
		t.Log("!!! No OAuth token. Some tests won't run. !!!")
	}
	gitany.RegisterClientGenerator(&github.ClientGenerator{})

	serviceConfig := &etc.ServiceConfig{
		Host:     "github.com",
		Type:     "github",
		Token:    token,
		Protocol: "https",
	}

	org := "gitany-test-org"

	ctx := context.Background()

	client, err := gitany.GetClient(context.Background(), serviceConfig)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = client.GetRepositories().ListByOrg(ctx, org, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = client.GetIssues().ListByRepo(ctx, "mpppk-test", "test-repo", nil)
	if err != nil {
		t.Fatal(err)
	}

	opt := &gitany.IssueListOptions{
		Filter: "all", // see https://developer.github.com/v3/issues/#list-issues
	}
	issues, res, err := client.GetIssues().ListByOrg(ctx, org, opt)
	if err != nil {
		t.Fatal(err, res.GetHTTPResponse())
	}

	if len(issues) <= 0 {
		t.Fatal("failed to fetch issues from "+org, res.GetHTTPResponse())
	}

	_, _, err = client.GetIssues().ListLabels(ctx, "mpppk-test", "test-repo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ListMilestonesByOrg always return err because group milestone does not implemented in github
	_, _, err = client.GetIssues().ListMilestonesByOrg(ctx, org, nil)
}
