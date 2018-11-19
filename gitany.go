package gitany

import (
	"context"
	"errors"
)

var clientGenerators []ClientGenerator

func RegisterClientGenerator(clientGenerator ClientGenerator) {
	clientGenerators = append(clientGenerators, clientGenerator)
}

func GetClient(ctx context.Context, serviceConfig *ServiceConfig) (Client, error) {
	for _, clientGenerator := range clientGenerators {
		if clientGenerator.GetType() == serviceConfig.Type {
			return clientGenerator.New(ctx, serviceConfig)
		}
	}
	return nil, errors.New("unknown serviceConfig type: " + serviceConfig.Type)
}

func CreateToken(ctx context.Context, serviceConfig *ServiceConfig, username, pass string) (string, error) {
	for _, clientGenerator := range clientGenerators {
		if clientGenerator.GetType() == serviceConfig.Type {
			client, err := clientGenerator.NewViaBasicAuth(ctx, serviceConfig, username, pass)
			if err != nil {
				return "", err
			}
			return client.GetAuthorizations().CreateToken(ctx)
		}
	}
	return "", errors.New("token creating failed because unknown serviceConfig type is provided: " + serviceConfig.Type)
}
