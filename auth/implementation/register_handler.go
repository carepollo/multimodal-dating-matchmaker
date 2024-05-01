package implementation

import (
	"context"
	"errors"
	"log"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/helpers"
	"github.com/carepollo/multimodal-dating-matchmaker/protos"
	"github.com/google/uuid"
)

func (s *AuthService) Register(ctx context.Context, req *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	log.Println("register request incoming: ", req.String())

	// validate incoming data
	if !helpers.ValidateEmail(req.Email) {
		return nil, errors.New("email is not valid")
	}

	if !helpers.ValidatePassword(req.Password) {
		return nil, errors.New("password doesn't comply with minimum security requirements")
	}

	hashedPassword, err := helpers.HashAndSalt(req.Password)
	if err != nil {
		return nil, errors.New("could not hash password: " + err.Error())
	}

	data := map[string]any{
		"_id":      uuid.New().String(),
		"email":    req.Email,
		"password": hashedPassword,
		"age":      18,
		"status":   "pending",
		"name":     req.Name,
	}

	// create user in db
	if err := helpers.CreateUser(s.Ctx, s.GetGraphDB(), data); err != nil {
		return &protos.RegisterResponse{}, errors.New("could not create user on db: " + err.Error())
	}

	log.Println("register response outbound", data)
	return &protos.RegisterResponse{}, nil
}
