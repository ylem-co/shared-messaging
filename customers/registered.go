package customers

const CustomerRegisteredMessageName = "customer.registered"

type CustomerRegistered struct {
	Id               string `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            *string `json:"phone"`
	OrganizationName string `json:"organization_name"`
	OrganizationUuid string `json:"organization_uuid"`
}

func (m *CustomerRegistered) GetMacawMessageKey() string {
	return CustomerRegisteredMessageName
}
