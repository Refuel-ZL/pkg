package log

import (
	"unicode"
)

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		messages = []interface{}{}
		var level = values[0]
		if level == "sql" {
			messages = append(messages, values[2:]...)
		} else {
			messages = append(messages, values)
		}
	}

	return
}

type Logger struct{}

func (Logger) Print(v ...interface{}) {
	LogZap.Debug(LogFormatter(v...))
}
func (Logger) Println(v ...interface{}) {
	LogZap.Debug(v)
}
