package messaging

const (
	TaskRunPipelineMessageName = "tasks.run_pipeline"

	ErrorRunPipelineTaskFailure = 11300
)

type RunPipelineTask struct {
	Task
	PipelineToRunUuid   string `json:"pipeline_uuid"`
}
