package messaging

const TaskCheckConditionMessageName = "tasks.check_condition"

type CheckConditionTask struct {
	Task
	Expression string `json:"expression"`
}
