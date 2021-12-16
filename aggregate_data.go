package messaging

const TaskAggregateDataMessageName = "tasks.aggregate_data"

type AggregateDataTask struct {
	Task
	Expression string `json:"expression"`
}
