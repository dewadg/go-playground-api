package executor

type InhouseConfigurator func(*inhouseConfig)

type inhouseConfig struct {
	tempDir      string
	numOfWorkers int
}

func InhouseWithTempDir(tempDir string) InhouseConfigurator {
	return func(c *inhouseConfig) {
		c.tempDir = tempDir
	}
}

func InhouseWithNumOfWorkers(num int) InhouseConfigurator {
	return func(c *inhouseConfig) {
		c.numOfWorkers = num
	}
}
