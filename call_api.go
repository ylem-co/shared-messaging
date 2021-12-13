package messaging

const TASK_CALL_API = "tasks.call_api"

type CallApiTask struct {
	Task
	URL              string `json:"url"`
	Method           string `json:"method"`
	Payload          []byte `json:"payload"`
	ContentType      string `json:"content_type"`
	AttachedFileName string `json:"attached_file_name"`
}
