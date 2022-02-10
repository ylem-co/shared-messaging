package messaging

const (
	TaskRunForEachMessageName = "tasks.run_for_each"

	ErrorRunForEachTaskFailure = 10600
)

type RunForEachTask struct {
	Task
}
