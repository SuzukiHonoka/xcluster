package utils

import (
	"io"
	"log"
)

func CloseAndPrintError(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Printf("iocloser: close error, err=%s", err)
	}
}
