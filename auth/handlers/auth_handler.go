package handlers

import (
	"context"
	"fmt"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/protos"
)

type AuthService struct {
	protos.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *protos.LoginRequest) (*protos.LoginResponse, error) {
	res := &protos.LoginResponse{
		Userid: "",
		Token:  "",
	}

	fmt.Printf("the payload is: %T -  %v \n", req, req)
	return res, nil
}
