package logoutsvc

import (
	"context"

	"api.heey.io/pkg/user"
)

// Service 登出服务
type Service interface {
	Logout(ctx context.Context, id string)
}

type service struct {
	ss user.SessionStore
}

// NewService 创建服务实例
func NewService(ss user.SessionStore) Service {
	return &service{
		ss: ss,
	}
}

func (s *service) Logout(ctx context.Context, id string) {
	// 删除会话
	s.ss.Del(id)
}
