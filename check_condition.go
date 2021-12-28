package messaging

const (
	TaskCheckConditionMessageName = "tasks.check_condition"

	ErrorCheckConditionTaskFailure = 10400
)

type CheckConditionTask struct {
	Task
	Expression string `json:"expression"`
}

type ConditionResult struct {
	Result        bool
	OriginalInput []byte
}
