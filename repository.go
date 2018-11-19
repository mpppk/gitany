package gitany

type Repository interface {
	GetID() int64
	GetName() string
	GetHTMLURL() string
	GetGitURL() string
	GetCloneURL() string
}

type repository struct {
	Id       int64
	Name     string
	HtmlURL  string
	GitURL   string
	CloneURL string
}

func (nr *repository) GetID() int64 {
	return nr.Id
}

func (nr *repository) GetName() string {
	return nr.Name
}

func (nr *repository) GetHTMLURL() string {
	return nr.HtmlURL
}

func (nr *repository) GetGitURL() string {
	return nr.GitURL
}

func (nr *repository) GetCloneURL() string {
	return nr.CloneURL
}

func NewRepository(name string) *repository {
	return &repository{
		Name: name,
	}
}
