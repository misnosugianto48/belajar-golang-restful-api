package helper

import log "github.com/sirupsen/logrus"

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
