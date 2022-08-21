package response

const (
	RequestSuccess = "Request successfully processed"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
