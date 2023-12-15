package types

type CheckSpamResponse struct {
	Number string `json:"number"`
	Spam bool `json:"spam"`
}