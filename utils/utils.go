package utils

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func ThrowError(err error, msg string) error {
	log.Error(err, msg)
	return fmt.Errorf("%s", msg)
}
