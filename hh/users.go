package hh

import "titan/db"

var users = store[int, User]()

type User struct {
	data db.User
}

func newuser(data db.User) *User {
	_, ok := users.find(data.ID)

	if ok {
		// block login or log the user out
	}

	u := &User{
		data,
	}

	users.add(data.ID, u)

	return u
}

func handshake(sso string) error {
	var data db.User
	if err := db.Conn.Take(&data, "sso = ?", sso).Error; err != nil {
		return err
	}

	newuser(data)

	return nil
}
