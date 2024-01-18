package messaging

const (
	TaskProcessDataMessageName = "tasks.process_data"

	ErrorProcessDataTaskFailure = 11200
)

type ProcessDataTask struct {
	Task
	Expression   string `json:"expression"`
	Strategy     string `json:"strategy"`
}
