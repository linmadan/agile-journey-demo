package constant

import "os"

const SERVICE_NAME = "agile-journey-demo"

var LOG_LEVEL = "debug"

func init() {
	if os.Getenv("LOG_LEVEL") != "" {
		LOG_LEVEL = os.Getenv("LOG_LEVEL")
	}
}
