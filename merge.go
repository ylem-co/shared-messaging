package messaging

const (
	TaskMergeMessageName = "tasks.merge"

	ErrorMergeTaskFailure = 10700
)

type MergeTask struct {
	Task
	FieldName string `json:"field_name"`
}
