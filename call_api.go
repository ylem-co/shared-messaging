package messaging

const (
	TaskCallApiMessageName = "tasks.call_api"

	ApiTypeGeneric   = "generic"
	ApiTypePagerDuty = "pager_duty"

	ApiAuthTypeBasic  = "Basic"
	ApiAuthTypeBearer = "Bearer"
	ApiAuthTypeHeader = "Header"

	ErrorCallApiTaskFailure            = 10200
	ErrorCallApiTaskDestinationOffline = 10201
)

type CallApiTask struct {
	Task
	Type             string            `json:"type"`
	Payload          string            `json:"payload"`
	QueryString      string            `json:"query_string"`
	Headers          map[string]string `json:"headers"`
	Severity         string            `json:"severity"`
	AttachedFileName string            `json:"attached_file_name"`
	Destination      ApiDestination    `json:"destination"`
}

type ApiDestination struct {
	Destination
	Type                  string                  `json:"type"`
	Method                string                  `json:"method"`
	AuthType              string                  `json:"auth_type"`
	AuthBearerToken       string                  `json:"auth_bearer_token"`
	AuthBasicUserName     string                  `json:"auth_basic_user_name"`
	AuthBasicUserPassword string                  `json:"auth_basic_user_password"`
	AuthHeaderName        string                  `json:"auth_header_name"`
	AuthHeaderValue       string                  `json:"auth_header_value"`
	PagerDuty             PagerDutyApiDestination `json:"pager_duty"`
}

type PagerDutyApiDestination struct {
	From               string   `json:"from"`
	AssigneeIds        []string `json:"assignee_ids"`
	ServiceId          string   `json:"service_id"`
	EscalationPolicyId string   `json:"escalation_policy_id"`
}
