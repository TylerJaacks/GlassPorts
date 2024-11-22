package main

import (
	"crypto/sha256"
	"encoding/json"
)

func (a *AuthToken) NewSHA256() string {
	json, _ := json.Marshal(a)
	hash := sha256.Sum256(json)
	return string(hash[:])
}
func NewSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return string(hash[:])
}

func UpdateAuthToken(old *AuthToken, new *AuthToken) {
	old.Retrieve()
	old.Current = "false"
	if new.Identifier == old.Identifier {
		new.PreviousHash = old.NewSHA256()
		new.Current = "true"
		new.Insert()
		old.Insert()
	}
}

func TryPassword(a *AuthToken, password string, ipAddr string) (bool, *User) {
	a.Current = "true"
	a.Retrieve()
	b := a
	if NewSHA256(password) == b.PasswordHash {
		b.LastIP = ipAddr
		UpdateAuthToken(a, b)
		u := User{
			Identifier: b.Identifier,
			Username:   b.Username,
			Email:      b.Email,
		}
		err := u.Retrieve()
		if err != nil {
			return false, nil
		}
		return true, &u
	}
	return false, nil

}

func ChangePassword(a *AuthToken, password string, ipAddr string) (bool, *AuthToken) {
	a.Current = "true"
	a.Retrieve()
	if a.Identifier == "" {
		return false, nil
	}
	b := a
	b.PasswordHash = NewSHA256(password)
	b.LastIP = ipAddr
	UpdateAuthToken(a, b)
	return true, b
}

func TryAuthToken(a *AuthToken) (bool, *User) {
	a.Current = "true"
	a.Retrieve()
	if a.Identifier != "" {
		u := User{
			Identifier: a.Identifier,
			Username:   a.Username,
			Email:      a.Email,
		}
		err := u.Retrieve()
		if err != nil {
			return false, nil
		}
		return true, &u
	}
	return false, nil
}
