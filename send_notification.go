package messaging

const (
	TaskSendNotificationMessageName = "tasks.send_notification"

	TaskSeverityLowest = "lowest"
	TaskSeverityLow = "low"
	TaskSeverityMedium = "medium"
	TaskSeverityHigh = "high"
	TaskSeverityCritical = "critical"

	NotificationTypeSms = "sms"
	NotificationTypeEmail = "email"
	NotificationTypeSlack = "slack"

	ErrorSendNotificationTaskFailure = 10100
)

type SendNotificationTask struct {
	Task
	Type                string             `json:"type"`
	Body                string             `json:"body"`
	Destination         Destination        `json:"destination"`
	SlackConfiguration  SlackConfiguration `json:"slack_configuration"`
	Severity            string             `json:"severity"`
	AttachedFileName    string             `json:"attached_file_name"`
	AttachedFileContent []byte             `json:"attached_file_content"`
}

type SlackConfiguration struct {
	AccessToken    string `json:"access_token"`
	SlackChannelId string `json:"slack_channel_id"`
}
