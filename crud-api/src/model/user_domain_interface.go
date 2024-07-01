package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetId() string
	SetId(string)
	EncryptPassword()
}
