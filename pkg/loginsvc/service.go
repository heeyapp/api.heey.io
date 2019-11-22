package loginsvc

import (
	"context"
	"encoding/base64"

	"api.heey.io/pkg/crypto"
	"api.heey.io/pkg/errors"
	"api.heey.io/pkg/user"
	uuid "github.com/satori/go.uuid"
)

// Service 登入服务
type Service interface {
	Login(ctx context.Context, username, password string) (Session, error)
}

type service struct {
	us user.Store
	ss user.SessionStore
	c  crypto.Crypto
}

// NewService 创建服务实例
func NewService(us user.Store, ss user.SessionStore, c crypto.Crypto) Service {
	return &service{
		us: us,
		ss: ss,
		c:  c,
	}
}

func (s *service) Login(ctx context.Context, username, password string) (Session, error) {
	// 验证
	if !user.ValidateEmail(username) || !user.ValidatePassWord(password) { // 无效账户密码
		return loginErrHandler(errors.ErrInvalidArgs)
	}

	// 查询用户
	u, err := s.us.SelectByUserNameAndPassWord(username, password)
	if err != nil {
		return loginErrHandler(err)
	}

	// 保存会话
	sessionID := s.encrypt(u.ID)
	sessionSecret := uuid.Must(uuid.NewV4()).String()
	s.ss.New(sessionID, sessionSecret)

	return Session{
		ID:     sessionID,
		Secret: sessionSecret,
	}, nil
}

// 加密
func (s *service) encrypt(text string) string {
	c, err := s.c.Encrypt([]byte(text))
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(c)
}

// 登入错误处理器
func loginErrHandler(err error) (Session, error) {
	return Session{}, err
}

// Session 会话
type Session struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}
