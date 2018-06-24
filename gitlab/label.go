package gitlab

import "github.com/xanzy/go-gitlab"

type Label struct {
	*gitlab.Label
}

func (l *Label) GetID() int64 {
	return int64(l.ID)
}

func (l *Label) GetName() string {
	return l.Name
}

func (l *Label) GetDescription() string {
	return l.Description
}
