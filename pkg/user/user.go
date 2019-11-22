package user

import (
	"database/sql"
	"errors"
)

// User 用户结构
type User struct {
	ID       string `db:"id"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
	State    bool   `db:"state"`
	Created  int64  `db:"created"`
	Updated  int64  `db:"updated"`
}

// IsValid 用户是否有效
func (u *User) IsValid() bool {
	return u.State
}

// IsEqual 用户是否相等
func (u *User) IsEqual(username, password string) bool {
	return u.UserName == username && u.PassWord == password
}

// 用户错误
var (
	ErrNotFoundUser  = errors.New("unregistered user")
	ErrForbiddenUser = errors.New("forbidden user")
	ErrIncorrectUser = errors.New("incorrect username or password")
)

// Store 用户存储
type Store interface {
	SelectByUserNameAndPassWord(username, password string) (*User, error)
}

type store struct {
	db *sql.DB
}

func (s *store) SelectByUserNameAndPassWord(username, password string) (*User, error) {
	u := new(User)

	if err := s.db.QueryRow("select * from users where username=?", username).Scan(
		&u.ID,
		&u.UserName,
		&u.PassWord,
		&u.State,
		&u.Created,
		&u.Updated,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFoundUser
		}

		panic(err)
	}

	if !u.IsEqual(username, password) {
		return nil, ErrIncorrectUser
	}

	if !u.IsValid() {
		return nil, ErrForbiddenUser
	}

	return u, nil
}
