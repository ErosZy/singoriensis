package interfaces

import (
	"net/http"
)

type DownloaderMiddlewareInterface interface {
	SetClient(*bool, *http.Client)
	SetRequest(*bool, *http.Request)
	GetResponse(*bool, *http.Response)
	Error(*bool, *http.Client, error)
}
