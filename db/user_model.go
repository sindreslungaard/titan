package db

type User struct {
	ID               int
	Username         string
	Email            string
	Password         string
	SSO              string
	Figure           string
	Gender           string
	Motto            string
	Credits          int
	Respect          int
	RespectToGive    int
	PetRespectToGive int
}
