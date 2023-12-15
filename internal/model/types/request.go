package types

type CheckSpamRequest struct {
	Number string `json:"number" binding:"required"`
}