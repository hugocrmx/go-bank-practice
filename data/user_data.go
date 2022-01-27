package data

import "fmt"

type User struct {
	Name  string
	Email string
	Alias string
}

func FindByAlias(alias string, users []User) (user User, err error) {
	for _, u := range users {
		if alias == u.Alias {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with alias %s not fount", alias)
}
