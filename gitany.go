package gitany

import (
	"context"
	"errors"

	"github.com/mpppk/gitany/etc"
)

var clientGenerators []ClientGenerator

func RegisterClientGenerator(clientGenerator ClientGenerator) {
	clientGenerators = append(clientGenerators, clientGenerator)
}

func GetClient(ctx context.Context, serviceConfig *etc.ServiceConfig) (Client, error) {
	for _, clientGenerator := range clientGenerators {
		if clientGenerator.GetType() == serviceConfig.Type {
			return clientGenerator.New(ctx, serviceConfig)
		}
	}
	return nil, errors.New("unknown serviceConfig type: " + serviceConfig.Type)
}
