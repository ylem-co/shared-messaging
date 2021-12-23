package messaging

const (
	TaskAggregateDataMessageName = "tasks.aggregate_data"

	ErrorAggregateDataTaskFailure = 10500
)

type AggregateDataTask struct {
	Task
	Expression string `json:"expression"`
}
