package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/helpers"
	"github.com/carepollo/multimodal-dating-matchmaker/auth/protos"
)

func (s *AuthService) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	log.Println("regsiter request logged")

	// validate incoming data
	if !helpers.ValidateEmail(req.Email) {
		return nil, errors.New("email is not valid")
	}

	if !helpers.ValidatePassword(req.Password) {
		return nil, errors.New("password doesn't comply with minimum security requirements")
	}

	// create user in db

	// create associated stuff in other db

	// send confirmation email
	return &protos.RegisterResponse{}, nil
}
