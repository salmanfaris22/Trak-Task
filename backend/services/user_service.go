package services

import (
	users "todo/domain/user"
	"todo/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreatUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	psSlice, er := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if er != nl {
		return nil, errors.NewBadReqstError("fild to bcrypt Passord")
	}

	user.Password = string(psSlice[:])
	user.Save()
}
