package api

import (
	"log"
)

const logger = Logger("utils")

type Logger string

func (l Logger) LogIfError(err error) {
	if err == nil {
		return
	}
	l.LogError(err)
}

func (l Logger) LogError(err error) {
	log.Printf("api: [%s] ERROR=%s", l, err)
}

func (l Logger) Log(s string) {
	log.Printf("api: [%s] %s", l, s)
}
