package etc

import (
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

type HostType int

type ServiceConfig struct {
	Host     string
	Type     string
	Token    string `mapstructure:"oauth_token" yaml:"oauth_token"`
	Protocol string
}

type Config struct {
	Services []*ServiceConfig
}

func (c *Config) FindServiceConfig(host string) (*ServiceConfig, bool) {
	for _, h := range c.Services {
		if strings.Contains(host, h.Host) {
			return h, true
		}
	}
	return nil, false
}

func (c *Config) FindServiceConfigs(host string) *Config {
	var serviceConfigs []*ServiceConfig
	for _, s := range c.Services {
		if strings.Contains(s.Host, host) {
			serviceConfigs = append(serviceConfigs, s)
		}
	}
	return &Config{Services: serviceConfigs}
}

func (c *Config) ListServiceConfigHost() (hosts []string) {
	for _, s := range c.Services {
		hosts = append(hosts, s.Host)
	}
	return hosts
}

func GetConfigDirName() string {
	return path.Join(".config", "gitany")
}

func GetConfigFileName() string {
	return ".gitany.yaml"
}

func GetConfigDirPath() (string, error) {
	dir, err := homedir.Dir()
	return path.Join(dir, GetConfigDirName()), errors.Wrap(err, "Error occurred in etc.GetConfigDirPath")
}

func GetConfigFilePath() (string, error) {
	configDirPath, err := GetConfigDirPath()
	return path.Join(configDirPath, GetConfigFileName()), errors.Wrap(err, "Error occurred in etc.GetConfigFilePath")
}
