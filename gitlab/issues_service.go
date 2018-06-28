package gitlab

import (
	"context"
	"fmt"

	"github.com/mpppk/gitany"
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type issuesService struct {
	projectURL  string
	rawClient   RawClient
	ListOptions *gitlab.ListOptions
}

func (i *issuesService) ListByRepo(ctx context.Context, owner, repo string, opt *gitany.IssueListByRepoOptions) (serviceIssues []gitany.Issue, res gitany.Response, err error) {
	gitlabOpt := toGitLabListProjectIssuesOptions(opt)
	issues, response, err := i.rawClient.GetIssues().ListProjectIssues(owner+"/"+repo, gitlabOpt)

	for _, issue := range issues {
		serviceIssues = append(serviceIssues, &Issue{Issue: issue})
	}

	return serviceIssues, &Response{Response: response}, errors.Wrap(err, "Failed to get Issues by raw client in gitlab.Client.GetIssues")
}

func (i *issuesService) ListByOrg(ctx context.Context, org string, opt *gitany.IssueListOptions) ([]gitany.Issue, gitany.Response, error) {
	gitlabOpt := toGitLabListGroupIssuesOptions(opt)
	gitlabIssues, res, err := i.rawClient.GetIssues().ListGroupIssues(org, gitlabOpt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get gitlab group issue")
	}

	var issues []gitany.Issue
	for _, gitlabIssue := range gitlabIssues {
		issues = append(issues, &Issue{Issue: gitlabIssue})
	}

	return issues, &Response{Response: res}, nil
}

func (i *issuesService) ListMilestonesByOrg(ctx context.Context, org string, opt *gitany.MilestoneListOptions) (milestones []gitany.Milestone, res gitany.Response, err error) {
	gitlabOpt := toGitLabListGroupMilestonesOptions(opt)
	rawMilestones, response, err := i.rawClient.GetGroupMilestones().ListGroupMilestones(org, gitlabOpt)
	if err != nil {
		return nil, &Response{Response: response}, errors.Wrap(err, "failed to fetch gitlab group milestone")
	}

	for _, rawMilestone := range rawMilestones {
		milestones = append(milestones, &GroupMilestone{GroupMilestone: rawMilestone})
	}

	return milestones, &Response{Response: response}, nil
}

func (i *issuesService) ListLabels(ctx context.Context, owner string, repo string, opt *gitany.ListOptions) (labels []gitany.Label, res gitany.Response, err error) {
	gitlabOpt := gitlab.ListLabelsOptions(toGitLabListOptions(opt))

	gitlabLabels, response, err := i.rawClient.GetLabels().ListLabels(owner+"/"+repo, &gitlabOpt)
	if err != nil {
		return nil, &Response{Response: response}, errors.Wrap(err, "failed to fetch labels in gitlab.Client.ListLabels")
	}

	for _, githubLabel := range gitlabLabels {
		labels = append(labels, &Label{Label: githubLabel})
	}

	return labels, &Response{Response: response}, nil
}

func (i *issuesService) GetIssuesURL(owner, repo string) (string, error) {
	return i.projectURL + "/issues", nil
}

func (i *issuesService) GetURL(owner, repo string, id int) (string, error) {
	url, err := i.GetIssuesURL(owner, repo)
	return fmt.Sprintf("%s/%d", url, id), errors.Wrap(err, "Error occurred in gitlab.Client.GetIssueURL")
}
