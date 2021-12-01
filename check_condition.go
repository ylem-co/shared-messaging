package messaging

const CHECK_CONDITION = "tasks.check_condition"

/**
 * Run query
 */
type CheckConditionTask struct {
	Task
	Expression string `json:"expression"`
}
