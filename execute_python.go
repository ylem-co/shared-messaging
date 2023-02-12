package messaging

const (
	TaskExecutePythonMessageName = "tasks.execute_python"

	ErrorExecutePythonFailure = 11000
)

type ExecutePythonTask struct {
	Task
	Code string `json:"code"`
}
