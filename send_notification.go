package messaging

const TaskSendNotificationMessageName = "tasks.send_notification"

type SendNotificationTask struct {
	Task
	Type                string             `json:"type"`
	Body                string             `json:"body"`
	Destination         Destination        `json:"destination"`
	SlackConfiguration  SlackConfiguration `json:"slack_configuration"`
	Severity            string             `json:"severity"`
	AttachedFileName    string             `json:"attached_file_name"`
	AttachedFilePayload []byte             `json:"attached_file_payload"`
}

type SlackConfiguration struct {
	AccessToken    string `json:"access_token"`
	SlackChannelId string `json:"slack_channel_id"`
}
