package messaging

const RUN_QUERY_TASK = "tasks.run_query"

/**
 * Run query
 */
type RunQueryTask struct {
	Task
	Source Source `json:"source"`
	Query  string `json:"query"`
}

type Source struct {
	DSN string `json:"dsn"`
}
