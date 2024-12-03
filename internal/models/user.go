package models

import "github.com/lallison21/auth_service/pkg/grpc_stubs/auth_service"

type UserDao struct {
	Id       int
	Username string
	Password string
	Email    string
}

type CreateUserDto struct {
	Username        string `json:"username" example:"username"`
	Password        string `json:"password,omitempty" example:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" example:"password_confirm"`
	Email           string `json:"email,omitempty" example:"email"`
}

func EmptyCreateUserDto() *CreateUserDto {
	return &CreateUserDto{}
}

func (c *CreateUserDto) FromGRPC(in *auth_service.RegisterRequest) *CreateUserDto {
	return &CreateUserDto{
		Username:        in.Username,
		Password:        in.Password,
		PasswordConfirm: in.PasswordConfirmation,
		Email:           in.Email,
	}
}

type LoginUserDto struct {
	Email    string `json:"email" example:"email"`
	Password string `json:"password,omitempty" example:"password"`
}

func EmptyLoginUserDto() *LoginUserDto {
	return &LoginUserDto{}
}

func (c *LoginUserDto) FromGRPC(in *auth_service.LoginRequest) *LoginUserDto {
	return &LoginUserDto{
		Email:    in.Email,
		Password: in.Password,
	}
}

type Tokens struct {
	AccessToken  string `json:"access_token,omitempty" example:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty" example:"refresh_token"`
}

func (t *Tokens) ToGRPC() *auth_service.LoginResponse {
	return &auth_service.LoginResponse{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
	}
}
