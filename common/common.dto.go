package common

type SuccessDto struct {
	Meta AckDto `json:"meta"`
	Data any
}

type ErrorDto struct {
	Meta AckDto `json:"meta"`
	Data any
}

type AckDto struct {
	Success bool    `json:"success"`
	Code    string  `json:"code"`
	Message *string `json:"message"`
}

func ToErrorDto(msg string) *ErrorDto {
	errorDto := &ErrorDto{
		Meta: AckDto{
			Success: false,
			Message: &msg,
		},
		Data: nil,
	}

	return errorDto
}
