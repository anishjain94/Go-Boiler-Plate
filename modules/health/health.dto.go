package health

type HealthResponseDto struct {
	Message      string            `json:"message"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
}
