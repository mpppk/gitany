package github

import (
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/mpppk/gitany"
)

type Response struct {
	*github.Response
}

type Rate struct {
	*github.Rate
}

func (r *Rate) GetLimit() int {
	return r.Limit
}

func (r *Rate) GetRemaining() int {
	return r.Remaining
}

func (r *Rate) GetReset() *time.Time {
	return &r.Reset.Time
}

func (r *Response) GetHTTPResponse() *http.Response {
	return r.Response.Response
}

func (r *Response) GetNextPage() int {
	return r.NextPage
}

func (r *Response) GetPrevPage() int {
	return r.PrevPage
}

func (r *Response) GetFirstPage() int {
	return r.FirstPage
}

func (r *Response) GetLastPage() int {
	return r.LastPage
}

func (r *Response) GetRate() gitany.Rate {
	return &Rate{Rate: &r.Rate}
}
