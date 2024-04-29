package util

import (
	"context"
	"net/http"
	"net/url"
)

func GetContextQueries(ctx *context.Context) *url.Values {
	queryValues := (*ctx).Value("CTX_QUERIES")
	ErrorIf(queryValues == nil, http.StatusInternalServerError, "MSG_QUERIES_MISSING")
	return queryValues.(*url.Values)
}

func GetContextParams(ctx *context.Context) *map[string]string {
	paramValues := (*ctx).Value("CTX_PARAMS")
	ErrorIf(paramValues == nil, http.StatusInternalServerError, "MSG_PARAMS_MISSING")
	return paramValues.(*map[string]string)
}
