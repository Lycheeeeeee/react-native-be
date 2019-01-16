package errorer

import "net/http"

var (
	ErrInvalidReceiveTime = errInvalidReceiveTime{}
	ErrInvalidUserName    = errInvalidUserName{}
	ErrInvalidAccount     = errInvalidAccount{}
)

type errInvalidReceiveTime struct{}

func (errInvalidReceiveTime) Error() string {
	return "receive time must be after order time"
}

func (errInvalidReceiveTime) StatusCode() int {
	return http.StatusBadRequest
}

type errInvalidUserName struct{}

func (errInvalidUserName) Error() string {
	return "user name is already exist, please enter unique user name"
}

func (errInvalidUserName) StatusCode() int {
	return http.StatusBadRequest
}

type errInvalidAccount struct{}

func (errInvalidAccount) Error() string {
	return "login failed, caused wrong user name or password"
}

func (errInvalidAccount) StatusCode() int {
	return http.StatusBadRequest
}
