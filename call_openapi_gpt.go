package messaging

const (
	TaskCallOpenapiGptMessageName = "tasks.call_openapi_gpt"

	ErrorCallOpenapiGptTaskFailure = 11100
)

type CallOpenapiGptTask struct {
	Task
	Prompt string `json:"type"`
}
