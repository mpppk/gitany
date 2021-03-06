package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mpppk/gitany"
	_ "github.com/mpppk/gitany/github"
)

func main() {
	token, ok := os.LookupEnv("GITANY_GITHUB_TEST_TOKEN")
	if !ok {
		_, err := fmt.Fprintln(os.Stderr, "GITHUB_TOKEN does not set")
		gitany.PanicIfErrorExist(err)
	}

	ctx := context.Background()
	config, ok := gitany.NewDefaultServiceConfig("github")
	if !ok {
		fmt.Println("github plugin of gitany does not imported")
	}

	config.Token = token
	githubClient, err := gitany.NewClient(ctx, config)
	gitany.PanicIfErrorExist(err)

	r := gitany.NewRepository("test")
	repo, _, err := githubClient.GetRepositories().Create(ctx, "", r)
	gitany.PanicIfErrorExist(err)

	fmt.Println(repo)
}
