package customers

const CustomerSendInviteMessageName = "customer.send_invite"

type CustomerSendInvite struct {
	Code             string `json:"code"`
	Email            string `json:"email"`
	Firstname        string `json:"first_name"`
	Lastname         string `json:"last_name"`
	OrganizationName string `json:"organization_name"`
}

func (m *CustomerSendInvite) GetMacawMessageKey() string {
	return CustomerSendInviteMessageName
}
