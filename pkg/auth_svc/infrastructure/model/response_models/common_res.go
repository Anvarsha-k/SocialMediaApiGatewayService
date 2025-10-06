package responsemodels_auth_apigw

type CommonResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`//The omitempty means → if Message is empty, it won’t appear in the JSON.
	Data		interface{}	`json:"after execution,omitempty"`//interface{} means it can hold any type (string, struct, list, etc.).
	Error 		interface{}	`json:"error,omitempty"` 		
}
