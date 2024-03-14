package messaging

import "github.com/google/uuid"

const (
	TaskRunQueryMessageName = "tasks.run_query"

	SQLIntegrationStatusNew     = "new"
	SQLIntegrationStatusOnline  = "online"
	SQLIntegrationStatusOffline = "offline"

	SQLIntegrationTypeMySQL             = "mysql"
	SQLIntegrationTypeSnowflake         = "snowflake"
	SQLIntegrationTypePostgresql        = "postgresql"
	SQLIntegrationTypeAWSRDS            = "aws-rds"
	SQLIntegrationTypeGoogleCloudSQL    = "google-cloud-sql"
	SQLIntegrationTypeGoogleBigQuery    = "google-bigquery"
	SQLIntegrationTypeMicrosoftAzureSQL = "microsoft-azure-sql"
	SQLIntegrationTypePlanetScale       = "planet-scale"
	SQLIntegrationTypeImmuta            = "immuta"
	SQLIntegrationTypeElasticsearch     = "elasticsearch"
	SQLIntegrationTypeRedshift          = "redshift"
	SQLIntegrationTypeClickhouse        = "clickhouse"

	SQLIntegrationConnectionTypeDirect = "direct"
	SQLIntegrationConnectionTypeSsh    = "ssh"

	ErrorRunQueryTaskFailure        = 10000
	ErrorRunQueryTaskOpenConnection = 10001
)

type RunQueryTask struct {
	Task
	Source SQLIntegration `json:"source"`
	Query  string `json:"query"`
}

type SQLIntegration struct {
	Uuid             uuid.UUID `json:"uuid"`
	CreatorUuid      uuid.UUID `json:"creator_uuid"`
	OrganizationUuid uuid.UUID `json:"organization_uuid"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	Name             string    `json:"name"`
	DataKey          []byte    `json:"data_key"`
	Host             []byte    `json:"host"`
	Port             int       `json:"port,omitempty"`
	User             string    `json:"user,omitempty"`
	Password         []byte    `json:"password"`
	Database         string    `json:"database,omitempty"`
	ConnectionType   string    `json:"connection_type"`
	SslEnabled       bool      `json:"ssl_enabled,omitempty"`
	SshHost          []byte    `json:"ssh_host,omitempty"`
	SshPort          int       `json:"ssh_port,omitempty"`
	SshUser          string    `json:"ssh_user,omitempty"`
	ProjectId        string    `json:"project_id,omitempty"`
	Credentials      []byte    `json:"credentials,omitempty"`
	EsVersion        uint8     `json:"es_version,omitempty"`
}
