package models

import "net/url"

// swagger:model
type User struct {
	// the id for this user
	//
	// min: 1
	ID uint `gorm:"primary_key" json:"id"`
	// the email for this user
	Email string `gorm:"column:email" json:"email"`
	// the name for this user
	Name string `gorm:"column:name" json:"name"`
	// the age for this user
	Age int `gorm:"column:age" json:"age"`
}

// Validate for user
func (u *User) Validate() url.Values {
	errs := url.Values{}

	if u.Age == 0 {
		errs.Add("age", "Age is required")
	}
	if u.Email == "" {
		errs.Add("email", "Email is required")
	}
	if u.Name == "" {
		errs.Add("name", "Name is required")
	}
	return errs
}
