package interfaces

import (
	"net/http"
	sErr "singoriensis/error"
)

type DownloaderMiddlewareInterface interface {
	SetClient(*bool, *http.Client)
	SetRequest(*bool, *http.Request, *sErr.RequestError)
	GetResponse(*bool, *http.Response, *sErr.ResponseError)
	Error(*bool, *http.Client, error)
}
