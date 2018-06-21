package service

type Issue interface {
	GetBody() string
	GetHTMLURL() string
	GetID() int64
	GetNumber() int
	GetRepository() Repository
	GetTitle() string
}

