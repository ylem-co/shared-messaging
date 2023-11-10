package messaging

const (
	TaskExecuteCodeMessageName = "tasks.execute_code"

	ErrorExecuteCodeFailure = 11000
)

type ExecuteCodeTask struct {
	Task
	Code string `json:"code"`
	Type string `json:"type"`
}
