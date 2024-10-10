package user

import (
	"encoding/json"
	"next-gen-job-hunting/common/db"
)

// User represents a user in the system.
//
// swagger:model User
//
// Fields:
//   - FirstName: The first name of the user. It is a string with a maximum size of 25 characters.
//     example: John
//   - LastName: The last name of the user. It is a string with a maximum size of 25 characters.
//     example: Doe
//   - Email: The email address of the user. It is a unique string with a maximum size of 50 characters and cannot be null.
//     example: john.doe@example.com
//   - Username: The username of the user. It is a unique string with a maximum size of 25 characters and cannot be null.
//     example: johndoe
//   - Password: The password of the user. It is a string with a maximum size of 25 characters and cannot be null.
//     example: secretpassword
type User struct {
	db.IdCreatedUpdated
	FirstName string `gorm:"size:25" json:"firstname"`
	LastName  string `gorm:"size:25" json:"lastname"`
	Email     string `gorm:"size:50;unique;not null" json:"email"`
	Username  string `gorm:"size:25;unique;not null" json:"username"`
	Password  string `gorm:"size:25;not null" json:"password"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
		Password string `json:"-"`
	}{
		Alias:    (Alias)(u),
		Password: "",
	})
}
