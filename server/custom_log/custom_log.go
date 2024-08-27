package custom_log

import "log"

func ErrorWithPanic(message string, err interface{}) {
	log.Println("[ERROR]:: ", message)
	panic(err)
}

func ErrorWithFatal(message string, err interface{}) {
	log.Println("[ERROR]:: ", message)
	log.Fatal(err)
}

func Info(message string, v interface{}) {
	log.Println("[INFO]:: ", message)
	if v != nil {
		log.Println(v)
	}
}

func Warn(message string, v interface{}) {
	log.Println("[WARN]:: ", message)
	log.Println(v)
}
