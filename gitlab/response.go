package gitlab

import (
	"net/http"

	"github.com/mpppk/gitany"
	"github.com/xanzy/go-gitlab"
)

type Response struct {
	*gitlab.Response
}

func (r *Response) GetHTTPResponse() *http.Response {
	return r.Response.Response
}

func (r *Response) GetNextPage() int {
	return r.NextPage
}

func (r *Response) GetPrevPage() int {
	return r.PreviousPage
}

func (r *Response) GetFirstPage() int {
	panic("gitlab.Response.GetFirstPage is not implemented")
}

func (r *Response) GetLastPage() int {
	panic("gitlab.Response.GetLastPage is not implemented")
}

func (r *Response) GetRate() gitany.Rate {
	panic("gitlab.Response.GetRate is not implemented")
}
