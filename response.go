package gitany

import (
	"net/http"
	"time"
)

type Response interface {
	GetHTTPResponse() *http.Response
	GetNextPage() int
	GetPrevPage() int
	GetFirstPage() int
	GetLastPage() int
	GetRate() Rate
}

type Rate interface {
	GetLimit() int
	GetRemaining() int
	GetReset() *time.Time
}
