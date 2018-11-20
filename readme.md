# gitany - Go library for accessing to multiple git repository hosting services


## Example1: Access to GitHub.com

First, import gitany and plugins for access to git hosting service that you want to connect.

```go
import "github.com/mpppk/gitany"
import _ "github.com/mpppk/gitany/github" // Load github plugin for gitany
```

Then you can access to the git hosting service!

```go
ctx := context.Background()

config := gitany.NewDefaultServiceConfig("github") // you can load github default config because github plugin is imported
config.Token = "xxxxx"
client, _ := gitany.NewClient(ctx, config) 
repos, _, _ := client.GetRepositories().ListByOrg(ctx, "org_name", nil) // you can access to github resources 
```

## Example2: Access to GitLab on your own server

```go
import "github.com/mpppk/gitany"
import _ "github.com/mpppk/gitany/gitlab" // Load gitlab plugin for gitany
```

```go
ctx := context.Background()

config := &gitany.ServiceConfig {
	Host:     "my.domain.com",
	Type:     "gitlab",
	Token:    "xxxxx",
	Protocol: "https", // or http
}

client, _ := gitany.NewClient(ctx, config) 
repos, _, _ := client.GetRepositories().ListByOrg(ctx, "org_name", nil)
```

## Example3: Copy GitHub issues to GitLab (Future Work)

```go
ctx := context.Background()

githubConfig := gitany.NewDefaultServiceConfig("github")
githubConfig.Token = "xxxxx"
githubClient, _ := gitany.NewClient(ctx, githubConfig) 

gitlabConfig := gitany.NewDefaultServiceConfig("gitlab")
gitlabConfig.Token = "xxxxx"
gitlabClient, _ := gitany.NewClient(ctx, gitlabConfig) 

issueOpt := &gitany.IssueListByRepoOptions{
	State: "open",
	Labels: []string{"bug"},
}

issues, _, _ := githubClient.GetIssues().ListByRepo(ctx, "repo_name", issueOpt)

// Copy GitHub issues to GitLab
for _, issue := range issues {
	gitlabClient.GetIssues().CreateIssue("owner/repo", issue)
}
```

## Example4: List repositories on AWS CodeCommit and Google Cloud Source Repositories (Future Work)
```go
ctx := context.Background()

awsConfig := gitany.NewDefaultServiceConfig("aws-code-commit")
awsConfig.Token = "xxxxx"
awsClient, _ := gitany.NewClient(ctx, awsConfig) 

gcpConfig := gitany.NewDefaultServiceConfig("gcp-cloud-source-repositories")
gcpConfig.Token = "xxxxx"
gcpClient, _ := gitany.NewClient(ctx, gcpConfig) 

clients := []*gitany.Client{awsClient, gcpClient}

func listRepos(clients []*gitany.Client, org string) (repos []gitany.Repository) {
	for _, c := range clients {
		rs, _, _ := c.GetRepositories().ListByOrg(org)
		repos = append(repos, rs)
	}
	return
}

fmt.Println(listRepos(clients, "org_name"))
// => Print repositories of specified org
//    on AWS CodeCommit and Google Cloud Source Repositories
```


