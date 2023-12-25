package implementation

import (
	"context"
	"fmt"

	"github.com/carepollo/multimodal-dating-matchmaker/protos"
)

func (s *AuthService) Login(ctx context.Context, req *protos.LoginRequest) (*protos.LoginResponse, error) {
	res := &protos.LoginResponse{
		Userid: "",
		Token:  "",
	}

	fmt.Printf("login")
	return res, nil
}
