package integration

import (
	"os"

	"testing"

	"context"

	"github.com/mpppk/gitany"
	"github.com/mpppk/gitany/etc"
	"github.com/mpppk/gitany/gitlab"
)

func TestGitLabIntegration(t *testing.T) {
	token := os.Getenv("GITANY_GITLAB_TEST_TOKEN")
	if token == "" {
		t.Log("!!! No OAuth token. Some tests won't run. !!!")
	}
	gitany.RegisterClientGenerator(&gitlab.ClientGenerator{})

	serviceConfig := &etc.ServiceConfig{
		Host:     "gitlab.com",
		Type:     "gitlab",
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
