package mock

type projectsService struct {
	ProjectsURL string
	URL         string
}

func NewProjectsService() *projectsService {
	return &projectsService{}
}

func (ps *projectsService) GetProjectsURL(owner, repo string) (string, error) {
	return ps.ProjectsURL, nil
}

func (ps *projectsService) GetURL(owner, repo string, no int) (string, error) {
	return ps.URL, nil
}
