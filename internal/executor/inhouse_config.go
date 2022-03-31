package executor

type InhouseConfigurator func(*inhouseConfig)

type inhouseConfig struct {
	tempDir string
}

func InhouseWithTempDir(tempDir string) InhouseConfigurator {
	return func(c *inhouseConfig) {
		c.tempDir = tempDir
	}
}
