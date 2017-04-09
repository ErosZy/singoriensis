package interfaces

import (
	"net/http"

	"github.com/ErosZy/singoriensis/common"
	sErr "github.com/ErosZy/singoriensis/error"
)

type DownloaderMiddlewareInterface interface {
	SetClient(*bool, *http.Client)
	SetRequest(*bool, *http.Request, *sErr.RequestError)
	GetResponse(*bool, *common.Page, *sErr.ResponseError)
	Error(*bool, *http.Client, error)
}
