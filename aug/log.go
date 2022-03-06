package aug

import "log"

func logInfo(format string, args ...any) {
	log.Printf("[INFO]: "+format, args...)
}

func logErr(err error) {
	log.Printf("[ERROR]: %v\n", err)
}
