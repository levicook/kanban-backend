package models

type (
	UserId struct{ identifier }

	Users []User

	User struct {
		atStamps

		Id UserId
	}
)

func NewUserId() UserId { return UserId{newIdentifier()} }
