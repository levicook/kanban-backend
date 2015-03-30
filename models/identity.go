package models

type Identity interface {
	CurrentUserId() UserId
}
