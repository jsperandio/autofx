package log

func Debug(args ...interface{}) {
	instance.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	instance.Debugf(template, args...)
}

func Debugln(args ...interface{}) {
	instance.Debugln(args...)
}

func Error(args ...interface{}) {
	instance.Error(args...)
}

func Fatal(args ...interface{}) {
	instance.Fatal(args...)
}
