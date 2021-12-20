package messaging

const TaskSendNotificationMessageName = "tasks.send_notification"

type SendNotificationTask struct {
	Task
	Name               string             `json:"name"`
	Type               string             `json:"type"`
	Body               string             `json:"body"`
	Destination        Destination        `json:"destination"`
	SlackAuthorization SlackAuthorization `json:"slack_authorization"`
}

type SlackAuthorization struct {
	AccessToken string `json:"access_token"`
	Scopes      string `json:"scopes"`
	BotUserId   string `json:"bot_user_id"`
}
