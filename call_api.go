package messaging

const TaskCallApiMessageName = "tasks.call_api"

type CallApiTask struct {
	Task
	Type             string         `json:"type"`
	Payload          string         `json:"payload"`
	QueryString      string         `json:"query_string"`
	Headers          string         `json:"headers"`
	AttachedFileName string         `json:"attached_file_name"`
	Destination      ApiDestination `json:"destination"`
}

type ApiDestination struct {
	Destination
	AuthType              string `json:"auth_type"`
	AuthBearerToken       string `json:"auth_bearer_token"`
	AuthBasicUserName     string `json:"auth_basic_user_name"`
	AuthBasicUserPassword string `json:"auth_basic_user_password"`
	AuthHeaderName        string `json:"auth_header_name"`
	AuthHeaderValue       string `json:"auth_header_value"`
}
