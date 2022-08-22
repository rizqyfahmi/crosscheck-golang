package response

const (
	ResponseMessageSuccess = "Request successfully processed"
	ResponseStatusSuccess  = "success"
	ResponseStatusError    = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
