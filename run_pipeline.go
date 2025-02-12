package messaging

const (
	TaskRunPipelineMessageName = "tasks.run_pipeline"

	ErrorRunPipelineTaskFailure = 11300
)

type RunPipelineTask struct {
	Task
	PipelineToRunUuid   string `json:"pipeline_to_run_uuid"`
}

type RunPipelineResult struct {
	Result        []byte `json:"result"`
	OriginalInput []byte `json:"original_input"`
}
