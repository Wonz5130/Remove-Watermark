package global

import (
	"os"
)

var (
	GoEnv   = getEnv()
	TIME_TZ = `2006-01-02T15:04:05.999Z`
)

func getEnv() string {
	env := os.Getenv(`GO_ENV`)
	if env == `` {
		env = `local`
	}
	return env
}

var Services map[string]string
