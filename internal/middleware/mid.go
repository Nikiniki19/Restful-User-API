package middleware

import "myproject/internal/auth"

type Middleware struct {
	auth auth.AuthInterface
}

func NewMiddleWare(auth auth.AuthInterface) (*Middleware, error) {
	return &Middleware{
		auth: auth,
	}, nil
}
