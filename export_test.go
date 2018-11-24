package gitany

func GetClientGenerators() []ClientGenerator {
	return clientGenerators
}

func ClearRegisteredClientGenerator() {
	clientGenerators = nil
}
