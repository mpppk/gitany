package gitany_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/mpppk/gitany"

	"github.com/mpppk/gitany/mock"
)

func TestRegisterClientGenerator(t *testing.T) {
	type args struct {
		clientGenerator gitany.ClientGenerator
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should add client generator to clientGenerators slice",
			args: args{
				clientGenerator: mock.NewClientGenerator("serviceA"),
			},
		},
	}
	for _, tt := range tests {
		gitany.ClearRegisteredClientGenerator()
		t.Run(tt.name, func(t *testing.T) {
			gitany.RegisterClientGenerator(tt.args.clientGenerator)
			clientGeneratorsLen := len(gitany.GetClientGenerators())
			if clientGeneratorsLen != 1 {
				t.Errorf("unexpected clientGenerators length: actual:%d, want:%d", clientGeneratorsLen, 1)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		ctx              context.Context
		clientGenerators []gitany.ClientGenerator
		serviceConfig    *gitany.ServiceConfig
	}

	serviceAClient := mock.NewClient()
	serviceAClient.Repositories.URL = "http://service-a.com"

	tests := []struct {
		name    string
		args    args
		want    gitany.Client
		wantErr bool
	}{
		{
			name: "should return specified service client",
			args: args{
				ctx: context.Background(),
				clientGenerators: []gitany.ClientGenerator{
					&mock.ClientGenerator{
						Client: serviceAClient,
						Type:   "serviceA",
					},
					&mock.ClientGenerator{
						Client: &mock.Client{},
						Type:   "serviceB",
					},
				},
				serviceConfig: &gitany.ServiceConfig{
					Type: "serviceA",
				},
			},
			want:    serviceAClient,
			wantErr: false,
		},
		{
			name: "should be error if specified service does not registered",
			args: args{
				ctx: context.Background(),
				clientGenerators: []gitany.ClientGenerator{
					&mock.ClientGenerator{
						Client: &mock.Client{},
						Type:   "serviceB",
					},
				},
				serviceConfig: &gitany.ServiceConfig{
					Type: "serviceA",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gitany.ClearRegisteredClientGenerator()
			for _, clientGenerator := range tt.args.clientGenerators {
				gitany.RegisterClientGenerator(clientGenerator)
			}
			client, err := gitany.NewClient(tt.args.ctx, tt.args.serviceConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(client, tt.want) {
				t.Errorf("NewClient() = %v, want %v", client, tt.want)
			}
		})
	}
}
