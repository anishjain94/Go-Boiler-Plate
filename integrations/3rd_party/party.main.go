package rdparty

import (
	"context"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var karzaClient *resty.Client

func InitializeKarzaService() {

	karzaUrl := ""
	karzaClient = resty.NewWithClient(&http.Client{}).
		SetBaseURL(karzaUrl).
		SetHeader("Content-Type", "application/json")
}

func GSTVerify(ctx *context.Context, gstDto interface{}) (any, *resty.Response, error) {

	// request := karzaClient.R().SetBody(gstDto).SetContext(*ctx)
	// resp, apiErr := request.SetError(&gstDetailResponse).SetResult(&gstDetailResponse).Post(URL_GST_DETAILS)

	// return &gstDetailResponse, resp, apiErr
	return nil, nil, nil
}
