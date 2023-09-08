package form

import (
	"user-service/user"

	"github.com/go-playground/validator/v10"
)

type UserForm struct {
	UserName     string `validate:"required"`
	UserNameKana string ``
	DisplayName  string ``
	Email        string `validate:"required,email"`
	TwitterId    string ``
	LoginId      string ``
	Pass         string `validate:"required"`
}

func UserInputForm(req *user.CreateUserRequest) error {
	validate := validator.New()

	user := UserForm{
		UserName:     req.UserName,
		UserNameKana: getOrDefault(req.UserNameKana, ""),
		DisplayName:  getOrDefault(req.DisplayName, ""),
		Email:        req.Email,
		TwitterId:    getOrDefault(req.TwitterId, ""),
		LoginId:      getOrDefault(req.LoginId, ""),
		Pass:         req.Pass,
	}

	if err := validate.Struct(user); err != nil {
		return err
	}

	return nil
}

func getOrDefault(ptr *string, defaultValue string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}
