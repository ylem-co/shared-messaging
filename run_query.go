package messaging

import "github.com/google/uuid"

const (
	TaskRunQueryMessageName = "tasks.run_query"

	SourceStatusNew = "new"
	SourceStatusOnline = "online"
	SourceStatusOffline = "offline"

	SourceTypeMySQL = "mysql"
	SourceTypeSnowflake = "snowflake"
	SourceTypePostgresql = "postgresql"
	SourceTypeAWSRDS = "aws-rds"
	SourceTypeGoogleCloudSQL = "google-cloud-sql"
	SourceTypeMicrosoftAzureSQL = "microsoft-azure-sql"

	SourceConnectionTypeDirect = "direct"
	SourceConnectionTypeSsh = "ssh"

	ErrorRunQueryTaskFailure = 10000
)

type RunQueryTask struct {
	Task
	Source Source `json:"source"`
	Query  string `json:"query"`
}

type Source struct {
	Uuid             uuid.UUID `json:"uuid"`
	CreatorUuid      uuid.UUID `json:"creator_uuid"`
	OrganizationUuid uuid.UUID `json:"organization_uuid"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	Name             string    `json:"name"`
	Host             string    `json:"host"`
	Port             int       `json:"port,omitempty"`
	User             string    `json:"user,omitempty"`
	Password         string    `json:"password"`
	Database         string    `json:"database,omitempty"`
	ConnectionType   string    `json:"connection_type"`
	SshHost          string    `json:"ssh_host,omitempty"`
	SshPort          int       `json:"ssh_port,omitempty"`
	SshUser          string    `json:"ssh_user,omitempty"`
}
