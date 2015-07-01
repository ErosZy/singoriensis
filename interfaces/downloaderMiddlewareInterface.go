package interfaces

import (
	"net/http"
)

type DownloaderMiddlewareInterface interface {
	SetClient(*http.Client)
	SetRequest(*http.Request)
	GetResponse(*http.Response)
	GetData([]byte)
	Error(error)
}
