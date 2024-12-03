package app_errors

import "errors"

var (
	ErrPassAndConfirmDoseNotMatch = errors.New("password and confirm dose not match")
	ErrWrongCredentials           = errors.New("wrong credentials")
	ErrInvalidHash                = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion        = errors.New("incompatible version of argon2")
)
