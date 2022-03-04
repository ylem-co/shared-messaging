package sources

const SourceStatusToggledMessageName = "source.status_toggled"

type SourceStatusToggled struct {
	Uuid string `json:"uuid"`
}

func (m *SourceStatusToggled) GetMacawMessageKey() string {
	return SourceStatusToggledMessageName
}
