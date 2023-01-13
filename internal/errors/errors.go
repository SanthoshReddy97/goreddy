package errors

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CustomError(args ...interface{}) {
	log.Error(args...)
	os.Exit(1)
}
