package messaging

import "github.com/google/uuid"

type Integration struct {
	Uuid             uuid.UUID `json:"uuid"`
	CreatorUuid      uuid.UUID `json:"-"`
	OrganizationUuid uuid.UUID `json:"-"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	Name             string    `json:"name"`
	Value            string    `json:"value"`
	UserUpdatedAt    string    `json:"user_updated_at"`
}
