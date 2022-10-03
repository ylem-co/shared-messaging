package messaging

const (
	TaskFilterMessageName = "tasks.filter"

	ErrorFilterTaskFailure = 10800
)

type FilterTask struct {
	Task
	Expression string `json:"expression"`
}
