package gitany

func GetClientGenerators() []ClientGenerator {
	return clientGenerators
}

func ClearRegisteredClientGenerator() {
	clientGenerators = nil
}

func GetDefaultServiceConfigs() []*ServiceConfig {
	return defaultServiceConfigs
}

func ClearRegisteredDefaultServiceConfig() {
	defaultServiceConfigs = nil
}
