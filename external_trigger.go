package messaging

const (
	TaskExternalTriggerMessageName = "tasks.external_trigger"

	ErrorExternalTriggerTaskFailure = 10900
)

type ExternalTriggerTask struct {
	Task
	Input []byte `json:"input"`
}
