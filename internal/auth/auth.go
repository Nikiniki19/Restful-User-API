package auth

const SignInKey = "nikiisthesigninkey"

type Authentication struct {
	SignInKey string
}

type AuthInterface interface {
	GenerateJWT(email, role string) (string, error)
	ValidateToken(token string) (Claims, error)
}

func NewAuth(signinkey string) (*Authentication, error) {
	return &Authentication{
		SignInKey: signinkey,
	}, nil
}
