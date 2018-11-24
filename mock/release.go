package mock

type release struct {
	Body    string
	HTMLURL string
	ID      int64
	Name    string
	TagName string
}

func NewRelease() *release {
	return &release{}
}

func (r *release) GetBody() string {
	return r.Body
}

func (r *release) GetHTMLURL() string {
	return r.HTMLURL
}

func (r *release) GetID() int64 {
	return r.ID
}

func (r *release) GetName() string {
	return r.Name
}

func (r *release) GetTagName() string {
	return r.TagName
}
