package auth

type Method uint8

const (
	MethodSkip Method = iota + 1
	MethodAccessToken
)
