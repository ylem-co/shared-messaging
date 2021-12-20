package messaging

const TaskCallApiMessageName = "tasks.call_api"

type CallApiTask struct {
	Task
	Type             string      `json:"type"`
	Payload          string      `json:"payload"`
	QueryString      string      `json:"query_string"`
	Headers          string      `json:"headers"`
	AttachedFileName string      `json:"attached_file_name"`
	Destination      Destination `json:"destination"`
}
