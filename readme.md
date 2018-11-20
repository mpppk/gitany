# gitany - Go library for accessing to multiple git repository hosting services

## Example1: Copy GitHub issues to GitLab 

```go
ctx := context.Background()

githubClient, _ := gitany.GetClient(ctx, gitany.NewGitHubConfig("github-token")) 
gitlabClient, _ := gitany.GetClient(ctx, gitany.NewGitLabConfig("gitlab-token")) 

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

## Example2: List repositories on AWS CodeCommit and Google Cloud Source Repositories
```go
ctx := context.Background()

awsClient, _ := gitany.GetClient(ctx, gitany.NewCodeCommitConfig("aws-token")) 
gcpClient, _ := gitany.GetClient(ctx, gitany.NewCloudSourceRepositoriesConfig("gcp-token")) 
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

## Example3: Access to git service on your own server

```go
ctx := context.Background()

config := &gitany.ServiceConfig {
    Host:     "my.domain.com",
    Type:     "gitlab",
    Token:    "xxxxxxx",
    Protocol: "http",
}

client, _ := gitany.GetClient(ctx, config) 
```
