// +build integration

package integration

import (
	"os"

	"testing"

	"context"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/gitlab"
)

func TestGitLabIntegration(t *testing.T) {
	token := os.Getenv("GITANY_GITLAB_TEST_TOKEN")
	if token == "" {
		t.Log("!!! No OAuth token. Some tests won't run. !!!")
	}
	gitany.RegisterClientGenerator(&gitlab.ClientGenerator{})

	serviceConfig := &gitany.ServiceConfig{
		Host:     "gitlab.com",
		Type:     "gitlab",
		Token:    token,
		Protocol: "https",
	}

	ctx := context.Background()

	org := "gitany-test"
	owner := "mpppk-test"
	repoName := "test-repo"

	client, err := gitany.NewClient(context.Background(), serviceConfig)
	if err != nil {
		t.Fatal(err)
	}

	orgRepos, res, err := client.GetRepositories().ListByOrg(ctx, org, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(orgRepos) <= 0 {
		t.Errorf("failed to fetch org repos from %s", org)
	}

	repoIssues, res, err := client.GetIssues().ListByRepo(ctx, owner, repoName, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(repoIssues) <= 0 {
		t.Errorf("failed to fetch repo issues from %s/%s", owner, repoName)
	}

	orgIssues, res, err := client.GetIssues().ListByOrg(ctx, org, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(orgIssues) <= 0 {
		t.Errorf("failed to fetch org issues from %s", org)
	}

	repoLabels, res, err := client.GetIssues().ListLabels(ctx, owner, repoName, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(repoLabels) <= 0 {
		t.Errorf("failed to fetch repo labels from %s/%s", owner, repoName)
	}

	milestones, res, err := client.GetIssues().ListMilestonesByOrg(ctx, org, nil)
	if err != nil {
		t.Fatal(err, res)
	}

	if len(milestones) <= 0 {
		t.Errorf("failed to fetch org milestones from %s", org)
	}
}
