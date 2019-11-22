package user

import (
	"github.com/seefan/gossdb"
)

// SessionStore 用户会话存储
type SessionStore interface {
	New(id, secret string)
	Del(id string)
}

type sessionstore struct {
}

func (ss *sessionstore) New(id, secret string) {
	if err := gossdb.Client().Set(id, secret); err != nil {
		panic(err)
	}
}

func (ss *sessionstore) Del(id string) {
	if err := gossdb.Client().Del(id); err != nil {
		panic(err)
	}
}
