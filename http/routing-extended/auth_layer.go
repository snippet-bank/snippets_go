package main

import "github.com/google/uuid"

type Login string
type UserData struct {
	Password string
}
type User struct {
	Login    Login
	UserData UserData
}
type Token string

type AuthLayerInterface interface {
	SignIn(login Login, password string) *Token
	IsValidToken(token Token) bool
}

type AuthLayerImpl struct {
	tokens map[Token]User
	users map[Login]UserData
}
func NewAuthLayerImpl() AuthLayerImpl {
	impl := AuthLayerImpl{ make(map[Token]User), make(map[Login]UserData) }
	impl.users["a"] = UserData{"b"} // test values
	return impl
}
func (a AuthLayerImpl) SignIn(login Login, password string) *Token {
	userData, ok := a.users[login]
	if !ok {
		return nil
	}
	if userData.Password != password {
		return nil
	}

	newTokenUUID, _ := uuid.NewUUID()
	newToken := Token(newTokenUUID.String())
	a.tokens[newToken] = User{Login: login, UserData: userData}
	return &newToken
}
func (a AuthLayerImpl) IsValidToken(token Token) bool {
	_, ok :=  a.tokens[token]
	return ok
}
