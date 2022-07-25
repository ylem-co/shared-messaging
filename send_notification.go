package messaging

const (
	TaskSendNotificationMessageName = "tasks.send_notification"

	TaskSeverityLowest   = "lowest"
	TaskSeverityLow      = "low"
	TaskSeverityMedium   = "medium"
	TaskSeverityHigh     = "high"
	TaskSeverityCritical = "critical"

	NotificationTypeSms        = "sms"
	NotificationTypeEmail      = "email"
	NotificationTypeSlack      = "slack"
	NotificationTypeJira       = "jira"
	NotificationTypeIncidentIo = "incidentio"

	ErrorSendNotificationTaskFailure            = 10100
	ErrorSendNotificationTaskDestinationOffline = 10101
	ErrorSendNotificationTaskUnconfirmedEmail   = 10102
	ErrorSendNotificationTaskUnconfirmedSms     = 10103
)

type SendNotificationTask struct {
	Task
	Type                    string                  `json:"type"`
	Body                    string                  `json:"body"`
	Destination             Destination             `json:"destination"`
	SlackConfiguration      SlackConfiguration      `json:"slack_configuration"`
	JiraConfiguration       JiraConfiguration       `json:"jira_configuration"`
	IncidentIoConfiguration IncidentIoConfiguration `json:"incidentio_configuration"`
	Severity                string                  `json:"severity"`
	AttachedFileName        string                  `json:"attached_file_name"`
	IsConfirmed             bool                    `json:"is_confirmed"`
}

type SlackConfiguration struct {
	AccessToken    string `json:"access_token"`
	SlackChannelId string `json:"slack_channel_id"`
}

type JiraConfiguration struct {
	Url         string `json:"url"`
	DataKey     []byte `json:"data_key"`
	AccessToken []byte `json:"access_token"`
	ProjectKey  string `json:"project_key"`
	IssueType   string `json:"issue_type"`
}

type IncidentIoConfiguration struct {
	ApiKey     []byte `json:"api_key"`
	DataKey    []byte `json:"data_key"`
	Mode       string `json:"mode"`
	Visibility string `json:"visibility"`
}
