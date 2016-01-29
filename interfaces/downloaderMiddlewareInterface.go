package interfaces

import (
	"net/http"
	sErr "singoriensis/error"
	"singoriensis/common"
)

type DownloaderMiddlewareInterface interface {
	SetClient(*bool, *http.Client)
	SetRequest(*bool, *http.Request, *sErr.RequestError)
	GetResponse(*bool, *common.Page, *sErr.ResponseError)
	Error(*bool, *http.Client, error)
}
