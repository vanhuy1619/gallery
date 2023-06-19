package model

type ResponeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
