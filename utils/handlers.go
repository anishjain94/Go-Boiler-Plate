package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

type SuccessDto struct {
	Meta AckDto `json:"meta"`
	Data any
}

type AckDto struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
}

func HandleHTTPPost[InputDtoType any, OutputDtoType any](serviceFunc func(ctx *context.Context, dto *InputDtoType) *OutputDtoType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pCtx := &ctx

		var dto InputDtoType

		err := json.NewDecoder(r.Body).Decode(&dto)
		AssertError(err, http.StatusBadRequest, "Invalid Request Body.")

		response := serviceFunc(pCtx, &dto)

		w.Header().Set(string("Content-Type"), string("application/json"))

		// extract or

		responseEncodeError := json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Code:    GetContextErrorCode(&ctx),
			},
			Data: response,
		})

		AssertError(responseEncodeError, http.StatusInternalServerError, "Error while parsing")
	}
}

func HandleHTTPPut[InputDtoType any, OutputDtoType any](serviceFunc func(ctx *context.Context, dto *InputDtoType) *OutputDtoType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pCtx := &ctx

		var dto InputDtoType

		err := json.NewDecoder(r.Body).Decode(&dto)
		AssertError(err, http.StatusBadRequest, "Invalid Request Body")

		response := serviceFunc(pCtx, &dto)

		w.Header().Set(string("Content-Type"), string("application/json"))

		responseEncodeError := json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Code:    GetContextErrorCode(&ctx),
			},
			Data: response,
		})

		AssertError(responseEncodeError, http.StatusInternalServerError, "Error while parsing")

	}
}

func HandleHTTPGet[OutputDtoType any](serviceFunc func(ctx *context.Context) *OutputDtoType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		response := serviceFunc(&ctx)

		w.Header().Set(string("Content-Type"), string("application/json"))
		responseEncodeError := json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Code:    GetContextErrorCode(&ctx),
			},
			Data: response,
		})

		AssertError(responseEncodeError, http.StatusInternalServerError, "Error while parsing")

	}
}

func GetContextErrorCode(ctx *context.Context) string {

	// Extract from ctx if any or else return Success

	return "SUCCESS"

}
