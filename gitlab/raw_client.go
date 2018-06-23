package gitlab

import "github.com/xanzy/go-gitlab"

type RawClient interface {
	GetProjects() ProjectsService
	GetGroups() GroupsService
	GetMergeRequests() MergeRequestsService
	GetIssues() IssuesService
	//GetTags() tagsService
	SetBaseURL(baseUrl string) error
}

type ProjectsService interface {
	GetProject(pid interface{}, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error)
	CreateProject(opt *gitlab.CreateProjectOptions, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error)
}

type GroupsService interface {
	GetGroup(gid interface{}, options ...gitlab.OptionFunc) (*gitlab.Group, *gitlab.Response, error)
	ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}

type IssuesService interface {
	ListProjectIssues(pid interface{}, opt *gitlab.ListProjectIssuesOptions, options ...gitlab.OptionFunc) ([]*gitlab.Issue, *gitlab.Response, error)
	ListGroupIssues(pid interface{}, opt *gitlab.ListGroupIssuesOptions, options ...gitlab.OptionFunc) ([]*gitlab.Issue, *gitlab.Response, error)
}

type LabelsService interface {
	ListLabels(pid interface{}, opt *gitlab.ListLabelsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Label, *gitlab.Response, error)
}

type MergeRequestsService interface {
	ListProjectMergeRequests(pid interface{}, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.OptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error)
	CreateMergeRequest(pid interface{}, opt *gitlab.CreateMergeRequestOptions, options ...gitlab.OptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error)
}

//type tagsService interface {
//	CreateTag(pid interface{}, opt *gitlab.CreateTagOptions, options ...gitlab.OptionFunc) (*gitlab.Tag, *gitlab.Response, error)
//}

type rawClient struct {
	*gitlab.Client
}

func (r *rawClient) GetProjects() ProjectsService {
	return r.Projects
}

func (r *rawClient) GetGroups() GroupsService {
	return r.Groups
}

func (r *rawClient) GetIssues() IssuesService {
	return r.Issues
}

func (r *rawClient) GetMergeRequests() MergeRequestsService {
	return r.MergeRequests
}

//func (r *rawClient) GetTags() tagsService {
//	return tagsService(r.Tags)
//}
