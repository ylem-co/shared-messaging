package messaging

const (
	TaskMergeMessageName = "tasks.merge"

	ErrorMergeTaskFailure = 10700
)

type MergeTask struct {
	Task
	FieldNames string `json:"field_names"`
}
