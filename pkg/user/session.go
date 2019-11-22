package user

import (
	"github.com/seefan/gossdb"
)

// SessionStore 用户会话存储
type SessionStore interface {
	New(id, secret string)
}

type sessionstore struct {
}

func (ss *sessionstore) New(id, secret string) {
	if err := gossdb.Client().Set(id, secret); err != nil {
		panic(err)
	}
}
