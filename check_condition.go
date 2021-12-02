package messaging

const TASK_CHECK_CONDITION = "tasks.check_condition"

/**
 * Run query
 */
type CheckConditionTask struct {
	Task
	Expression string `json:"expression"`
}
