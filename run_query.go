package messaging

import "github.com/google/uuid"

const TASK_RUN_QUERY = "tasks.run_query"

/**
 * Run query
 */
type RunQueryTask struct {
	Task
	Source Source `json:"source"`
	Query  string `json:"query"`
}

type Source struct {
	Uuid uuid.UUID `json:"uuid"`
}
