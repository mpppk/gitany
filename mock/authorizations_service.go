package mock

import "context"

type authorizationsService struct {
	Token string
}

func NewAuthorizationService() *authorizationsService {
	return &authorizationsService{}
}

func (as *authorizationsService) CreateToken(ctx context.Context) (string, error) {
	return as.Token, nil
}
