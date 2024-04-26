package middleware

import (
	"context"
	"encoding/json"
	"go-boiler-plate/common"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		r = r.WithContext(ctx)

		defer HandleError(w, r)
		next.ServeHTTP(w, r)
	})
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		ctx := r.Context()

		msg := getMessageFromError(err)
		statusCode, msg := SeparateCodeFromMsg(&ctx, msg)

		setStatusCode(w, statusCode)

		dto := common.ToErrorDto(msg)

		json.NewEncoder(w).Encode(dto)
	}
}

func getMessageFromError(err any) string {

	var msg string

	if reflect.TypeOf(err).Kind().String() == "error" {
		msg = err.(error).Error()
	} else if reflect.TypeOf(err).Kind() == reflect.String {
		msg = reflect.ValueOf(err).String()
	} else {
		msg = reflect.TypeOf(err).String()
	}

	stackField := zap.Stack("stack")
	zap.L().Error(msg + "\n" + stackField.String)

	return msg
}

func SeparateCodeFromMsg(ctx *context.Context, msg string) (int, string) {

	statusCode := 200

	parts := strings.Split(msg, ":")

	if len(parts) > 1 {
		if codeInt, err := strconv.Atoi(parts[0]); err == nil {
			statusCode = codeInt
			msg = strings.Join(parts[1:], ":")
		}
	}

	return statusCode, msg
}

func setStatusCode(w http.ResponseWriter, code int) {
	if code >= 100 {
		w.WriteHeader(code)
	}
}
