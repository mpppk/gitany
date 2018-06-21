package gitlab

import "github.com/xanzy/go-gitlab"

type Label struct {
	*gitlab.Label
}

func (l *Label) GetName() string {
	return l.GetName()
}
