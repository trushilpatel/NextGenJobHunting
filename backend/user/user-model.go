package user

import "next-gen-job-hunting/common/db"

type User struct {
	db.IdCreatedUpdated
	FirstName string `gorm:"size:25" json:"first_name"`
	LastName  string `gorm:"size:25" json:"last_name"`
	Username  string `gorm:"size:25;unique;not null" json:"username"`
	Password  string `gorm:"size:25;not null" json:"password"`
}
