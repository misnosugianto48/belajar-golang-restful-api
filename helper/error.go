package helper

import log "github.com/sirupsen/logrus"

func PanicIfError(msg string, err error) {
	if err != nil {
		log.Panic(msg, err)
	}
}
