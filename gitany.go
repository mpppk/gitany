package gitany

import (
	"context"
	"errors"
)

var clientGenerators []ClientGenerator
var defaultServiceConfigs []*ServiceConfig

func RegisterClientGenerator(clientGenerator ClientGenerator) {
	clientGenerators = append(clientGenerators, clientGenerator)
}

func RegisterDefaultServiceConfig(serviceConfig *ServiceConfig) {
	defaultServiceConfigs = append(defaultServiceConfigs, serviceConfig)
}

func ClearRegisteredClientGenerator() {
	clientGenerators = []ClientGenerator{}
}

func NewClient(ctx context.Context, serviceConfig *ServiceConfig) (Client, error) {
	for _, clientGenerator := range clientGenerators {
		if clientGenerator.GetType() == serviceConfig.Type {
			return clientGenerator.New(ctx, serviceConfig)
		}
	}
	return nil, errors.New("unknown serviceConfig type: " + serviceConfig.Type)
}

func NewDefaultServiceConfig(serviceType string) (*ServiceConfig, bool) {
	for _, config := range defaultServiceConfigs {
		if config.Type == serviceType {
			return config, true
		}
	}
	return nil, false
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
