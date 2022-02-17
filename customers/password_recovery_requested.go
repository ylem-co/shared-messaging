package customers

const CustomerPasswordRecoveryRequestedMessageName = "customer.password_recovery_requested"

type CustomerPasswordRecoveryRequested struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func (m *CustomerPasswordRecoveryRequested) GetMacawMessageKey() string {
	return CustomerPasswordRecoveryRequestedMessageName
}

