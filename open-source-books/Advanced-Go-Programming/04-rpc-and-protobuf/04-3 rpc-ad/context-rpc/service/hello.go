package service

import (
	"errors"
	"fmt"
	"net"
)

var (
	ErrAuth  = errors.New("not login")
	ErrLogin = errors.New("username or password incorrect")
)

type HelloService struct {
	Conn    net.Conn
	isLogin bool
}

func (h *HelloService) Login(request string, reply *string) error {
	if request != "user:pass" {
		return ErrLogin
	}
	h.isLogin = true
	return nil
}

func (h *HelloService) Hello(request string, reply *string) error {
	if !h.isLogin {
		return ErrAuth
	}
	*reply = fmt.Sprintf("Hello %s, from %s", request, h.Conn.RemoteAddr())
	return nil
}
