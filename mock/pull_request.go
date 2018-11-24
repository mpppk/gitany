package mock

type pullRequest struct {
	Number  int
	Title   string
	HTMLURL string
}

func NewPullRequest() *pullRequest {
	return &pullRequest{}
}

func (p *pullRequest) GetNumber() int {
	return p.Number
}

func (p *pullRequest) GetTitle() string {
	return p.Title
}

func (p *pullRequest) GetHTMLURL() string {
	return p.HTMLURL
}
