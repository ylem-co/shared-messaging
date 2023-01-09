package messaging

import "github.com/google/uuid"

const (
	TaskRunQueryMessageName = "tasks.run_query"

	SourceStatusNew     = "new"
	SourceStatusOnline  = "online"
	SourceStatusOffline = "offline"

	SourceTypeMySQL             = "mysql"
	SourceTypeSnowflake         = "snowflake"
	SourceTypePostgresql        = "postgresql"
	SourceTypeAWSRDS            = "aws-rds"
	SourceTypeGoogleCloudSQL    = "google-cloud-sql"
	SourceTypeGoogleBigQuery    = "google-bigquery"
	SourceTypeMicrosoftAzureSQL = "microsoft-azure-sql"
	SourceTypePlanetScale       = "planet-scale"
	SourceTypeImmuta            = "immuta"
	SourceTypeElasticsearch     = "elasticsearch"
	SourceTypeRedshift          = "redshift"
	SourceTypeClickhouse        = "clickhouse"

	SourceConnectionTypeDirect = "direct"
	SourceConnectionTypeSsh    = "ssh"

	ErrorRunQueryTaskFailure        = 10000
	ErrorRunQueryTaskOpenConnection = 10001
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
