package implementation

import (
	"context"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/helpers"
	"github.com/carepollo/multimodal-dating-matchmaker/protos"
)

func (s *AuthService) Login(ctx context.Context, req *protos.LoginRequest) (*protos.LoginResponse, error) {
	res := &protos.LoginResponse{
		Userid: req.Email,
		Token:  req.Password,
	}

	user, err := helpers.GetUserByEmailAndPassword(ctx, *s.GraphDB, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(user.Id, s.Env.JwtSecret)
	if err != nil {
		return nil, err
	}

	res.Token = token
	res.Userid = user.Id
	return res, nil
}
