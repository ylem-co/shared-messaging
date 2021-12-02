package messaging

import "github.com/google/uuid"

const TASK_SEND_NOTIFICATION = "tasks.send_notification"

/**
 * Send SMS
 */
type SendNotificationTask struct {
	Task
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Body        string      `json:"body"`
	Destination Destination `json:"destination"`
}

type Destination struct {
	Uuid               uuid.UUID          `json:"uuid"`
	Type               string             `json:"type"`
	AuthType           string             `json:"auth_type"`
	AuthToken          string             `json:"auth_token"`
	Code               string             `json:"code"`
	SlackAuthorization SlackAuthorization `json:"slack_authorization"`
	SlackChannelId     string             `json:"slack_channel_id"`
}

type SlackAuthorization struct {
	AccessToken string `json:"access_token"`
	Scopes      string `json:"scopes"`
	BotUserId   string `json:"bot_user_id"`
}
