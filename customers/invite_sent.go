package customers

const CustomerInviteSentMessageName = "customer.invite_sent"

type CustomerInviteSent struct {
	Code             string `json:"code"`
	Email            string `json:"email"`
	Firstname        string `json:"first_name"`
	Lastname         string `json:"last_name"`
	OrganizationName string `json:"organization_name"`
}

func (m *CustomerInviteSent) GetMacawMessageKey() string {
	return CustomerInviteSentMessageName
}
