package users

import (
	"strings"
	"todo/utils/errors"
)

type User struct {
	ID        int64  `json"id"`
	FirstName string `json"firstName"`
	LastName  string `json"lastName"`
	Email     string `json"email"`
	Password  string `json"password"`
}

func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return errors.NewBadReqstError("Invalid Email adress")
	}
	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadReqstError("Invalid  Password")
	}

	return nil
}
