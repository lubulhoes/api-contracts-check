package internals

type AppDTO struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	EndpointsBody map[string]any `json:"endpoints_body"`
}
