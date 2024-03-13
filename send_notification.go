package messaging

const (
	TaskSendNotificationMessageName = "tasks.send_notification"

	TaskSeverityLowest   = "lowest"
	TaskSeverityLow      = "low"
	TaskSeverityMedium   = "medium"
	TaskSeverityHigh     = "high"
	TaskSeverityCritical = "critical"

	NotificationTypeSms          = "sms"
	NotificationTypeEmail        = "email"
	NotificationTypeSlack        = "slack"
	NotificationTypeJira         = "jira"
	NotificationTypeIncidentIo   = "incidentio"
	NotificationTypeOpsgenie     = "opsgenie"
	NotificationTypeTableau      = "tableau"
	NotificationTypeHubspot      = "hubspot"
	NotificationTypeSalesforce   = "salesforce"
	NotificationTypeGoogleSheets = "google-sheets"
	NotificationTypeJenkins      = "jenkins"

	ErrorSendNotificationTaskFailure            = 10100
	ErrorSendNotificationTaskIntegrationOffline = 10101
	ErrorSendNotificationTaskUnconfirmedEmail   = 10102
	ErrorSendNotificationTaskUnconfirmedSms     = 10103

	GoogleSheetsModeOverwrite = "overwrite"
	GoogleSheetsModeAppend    = "append"
)

type SendNotificationTask struct {
	Task
	Type                      string                    `json:"type"`
	Body                      string                    `json:"body"`
	Integration               Integration               `json:"integration"`
	SlackConfiguration        SlackConfiguration        `json:"slack_configuration"`
	JiraConfiguration         JiraConfiguration         `json:"jira_configuration"`
	IncidentIoConfiguration   IncidentIoConfiguration   `json:"incidentio_configuration"`
	OpsgenieConfiguration     OpsgenieConfiguration     `json:"opsgenie_configuration"`
	TableauConfiguration      TableauConfiguration      `json:"tableau_configuration"`
	HubspotConfiguration      HubspotConfiguration      `json:"hubspot_configuration"`
	SalesforceConfiguration   SalesforceConfiguration   `json:"salesforce_configuration"`
	GoogleSheetsConfiguration GoogleSheetsConfiguration `json:"google_sheets_configuration"`
	JenkinsConfiguration      JenkinsConfiguration      `json:"jenkins_configuration"`
	Severity                  string                    `json:"severity"`
	AttachedFileName          string                    `json:"attached_file_name"`
	IsConfirmed               bool                      `json:"is_confirmed"`
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

type OpsgenieConfiguration struct {
	ApiKey  []byte `json:"api_key"`
	DataKey []byte `json:"data_key"`
}

type TableauConfiguration struct {
	Server         string `json:"server"`
	DataKey        []byte `json:"data_key"`
	Username       []byte `json:"username"`
	Password       []byte `json:"password"`
	Sitename       string `json:"site_name"`
	ProjectName    string `json:"project_name"`
	DatasourceName string `json:"datasource_name"`
	Mode           string `json:"mode"`
}

type HubspotConfiguration struct {
	DataKey           []byte `json:"data_key"`
	AccessToken       []byte `json:"access_token"`
	PipelineCode      string `json:"pipeline_code"`
	PipelineStageCode string `json:"pipeline_stage_code"`
	OwnerCode         string `json:"owner_code"`
}

type SalesforceConfiguration struct {
	DataKey     []byte `json:"data_key"`
	AccessToken []byte `json:"access_token"`
	Domain      string `json:"domain"`
}

type GoogleSheetsConfiguration struct {
	DataKey       []byte `json:"data_key"`
	Credentials   []byte `json:"credentials"`
	Mode          string `json:"mode"`
	SpreadsheetId string `json:"spreadsheet_id"`
	SheetId       int64  `json:"sheet_id"`
	WriteHeader   bool   `json:"write_header"`
}

type JenkinsConfiguration struct {
	BaseUrl string `json:"base_url"`
	Token   []byte `json:"token"`
	DataKey []byte `json:"data_key"`
}
