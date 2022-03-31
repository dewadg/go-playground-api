package inhouse

type Configurator func(*config)

type config struct {
	tempDir string
}

func WithTempDir(tempDir string) Configurator {
	return func(c *config) {
		c.tempDir = tempDir
	}
}
