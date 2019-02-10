package domain

type User struct {
	InternalId int    `json:"internalId"`
	UserId     string `json:"userId"`
	Name       string `json:"name"`
}
