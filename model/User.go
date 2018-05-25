package model

type User struct {
	Id    int
	Email string `dynamo:"Email"`
}
