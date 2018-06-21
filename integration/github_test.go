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

	ctx := context.Background()

	t.Log(serviceConfig)

	client, err := gitany.GetClient(context.Background(), serviceConfig)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetIssues().ListByRepo(ctx, "mpppk-test", "test-repo")
	if err != nil {
		t.Fatal(err)
	}
}
