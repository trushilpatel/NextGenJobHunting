package user

import (
	"encoding/json"
	"next-gen-job-hunting/common/db"
)

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
