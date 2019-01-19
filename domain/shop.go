package domain

type Shop struct {
	Model
	Name      string `json:"name"`
	Address   string `json:"address"`
	AccountID UUID   `json:"account_id"`
}
