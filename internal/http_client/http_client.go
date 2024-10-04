package http

import (
	"go-boiler-plate/config"
	"net/http"
	"net/url"
	"strings"

	"github.com/gojek/heimdall/httpclient"
	"github.com/joomcode/errorx"
)

var HC *httpclient.Client

func InitializeHttpClient(cfg *config.HttpClientConfig) {
	var hcopts []httpclient.Option
	useProxy := strings.EqualFold(config.Environment, "development") && cfg.ProxyURL != ""
	if useProxy {
		pu, err := url.Parse(cfg.ProxyURL)
		if err != nil {
			panic(errorx.Decorate(err, "failed to parse proxy url"))
		}
		hcopts = append(hcopts, httpclient.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(pu),
			},
		}))
	}
	HC = httpclient.NewClient(hcopts...)
}
