package mock

type repository struct {
	CloneURL string
	GitURL   string
	HTMLURL  string
	ID       int64
	Name     string
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetCloneURL() string {
	return r.CloneURL
}

func (r *repository) GetGitURL() string {
	return r.GitURL
}

func (r *repository) GetHTMLURL() string {
	return r.HTMLURL
}

func (r *repository) GetID() int64 {
	return r.ID
}

func (r *repository) GetName() string {
	return r.Name
}
